package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Eric-GreenComb/contrib/rand"
	"github.com/gin-gonic/gin"
	nsq "github.com/nsqio/go-nsq"

	"github.com/Eric-GreenComb/one-pushinfo/bean"
	"github.com/Eric-GreenComb/one-pushinfo/config"
)

// WriteBlock WriteBlock
func WriteBlock(c *gin.Context) {

	var _oneOrder bean.OneOrder
	c.BindJSON(&_oneOrder)

	_config := nsq.NewConfig()
	_Producer, err := nsq.NewProducer(config.Nsq.Host, _config)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_oneOrder.Type = 0

	_topicNum := rand.GetRandomItNum(config.Nsq.TopicNum)
	_topic := fmt.Sprintf("%s%d", config.Nsq.Topic, _topicNum)

	_json, err := json.Marshal(_oneOrder)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	err = _Producer.Publish(_topic, _json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	defer _Producer.Stop()

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "OK"})
}

// PutWinerTxID PutWinerTxID
func PutWinerTxID(c *gin.Context) {

	var _oneOrder bean.OneOrder
	c.BindJSON(&_oneOrder)

	_config := nsq.NewConfig()
	_Producer, err := nsq.NewProducer(config.Nsq.Host, _config)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_oneOrder.Type = 1

	_topicNum := rand.GetRandomItNum(config.Nsq.TopicNum)
	_topic := fmt.Sprintf("%s%d", config.Nsq.Topic, _topicNum)

	_json, err := json.Marshal(_oneOrder)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	err = _Producer.Publish(_topic, _json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	defer _Producer.Stop()

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "OK"})
}
