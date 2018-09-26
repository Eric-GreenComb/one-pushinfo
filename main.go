package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/sync/errgroup"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/Eric-GreenComb/one-pushinfo/config"
	"github.com/Eric-GreenComb/one-pushinfo/ethereum"
	"github.com/Eric-GreenComb/one-pushinfo/handler"
	// "github.com/Eric-GreenComb/one-pushinfo/nsq"
	"github.com/Eric-GreenComb/one-pushinfo/persist"
)

var (
	g errgroup.Group
)

func main() {
	if config.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	persist.InitDatabase()

	ethereum.Init()

	// nsq.Init()

	_nonce, err := ethereum.PendingNonce(config.Ethereum.Address)
	if err != nil {
		log.Fatal(err)
	}
	config.PendingNonce = _nonce
	fmt.Println(config.Ethereum.Address, " PendingNonce ", config.PendingNonce)

	router := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.MaxMultipartMemory = 64 << 20 // 64 MiB

	router.Use(Cors())

	/* api base */
	r0 := router.Group("/")
	{
		r0.GET("", handler.Index)
		r0.GET("health", handler.Health)
	}

	// api
	r1 := router.Group("/block")
	{
		r1.POST("/write", handler.WriteBlock)
		r1.GET("/read/:orderid", handler.ReadBlock)
		r1.POST("/winer", handler.PutWinerTxID)
		r1.GET("/orders/:catid/:patchid", handler.GetAllOrders)
	}

	r2 := router.Group("/ethereum")
	{
		r2.GET("/nonce", handler.PendingNonce)
		r2.POST("/send", handler.SendEthCoin)
		r2.GET("/balance/:addr", handler.GetBalance)
	}

	r3 := router.Group("/nsq")
	{
		r3.POST("/write", handler.WriteNsq)
	}

	r100 := router.Group("/badger")
	{
		r100.POST("/set", handler.SetBadgerKey)
		r100.GET("/get/:key", handler.GetBadgerKey)
	}

	for _, _port := range config.Server.Port {
		server := &http.Server{
			Addr:         ":" + _port,
			Handler:      router,
			ReadTimeout:  300 * time.Second,
			WriteTimeout: 300 * time.Second,
		}

		g.Go(func() error {
			return server.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
