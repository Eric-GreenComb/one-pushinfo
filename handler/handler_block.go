package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/one-pushinfo/bean"
	"github.com/Eric-GreenComb/one-pushinfo/persist"
)

// WriteBlock WriteBlock
func WriteBlock(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_txID := "0x33a4f4edcedc66eb05ce4bf54dee113694cffeaf34fc0e6c71ddb454c4c72a15"

	var _order bean.Order
	_order.OrderID = _formParams.OrderID
	_order.Amount = _formParams.Amount
	_order.CatID = _formParams.CatID
	_order.PatchID = _formParams.PatchID
	_order.BuyTime = _formParams.BuyTime
	_order.Account = _formParams.Account
	_order.Mobile = _formParams.Mobile
	_order.Type = 0
	_order.Desc = _formParams.Desc
	_order.TxID = _txID

	err := persist.GetPersist().CreateOrder(_order)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _txID})
}

// ReadBlock ReadBlock
func ReadBlock(c *gin.Context) {

	_orderid := c.Params.ByName("orderid")

	_order, err := persist.GetPersist().OrderInfo(_orderid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _order})
}

// PutWinerTxID PutWinerTxID
func PutWinerTxID(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_txID := "0xcbe3201c1699ad7b6be1855aed4fae8e6b16561bdaf54bf08f09964531067adb"

	var _order bean.Order
	_order.OrderID = _formParams.OrderID
	_order.Amount = _formParams.Amount
	_order.CatID = _formParams.CatID
	_order.PatchID = _formParams.PatchID
	_order.BuyTime = _formParams.BuyTime
	_order.Account = _formParams.Account
	_order.Mobile = _formParams.Mobile
	_order.Type = 1
	_order.Desc = _formParams.Desc
	_order.TxID = _txID

	err := persist.GetPersist().CreateOrder(_order)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _txID})
}

// GetAllOrders GetAllOrders
func GetAllOrders(c *gin.Context) {

	_catid := c.Params.ByName("catid")
	_patchid := c.Params.ByName("patchid")

	_orders, err := persist.GetPersist().GetAllOrders(_catid, _patchid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _orders})
}
