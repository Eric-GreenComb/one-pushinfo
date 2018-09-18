package handler

import (
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/one-pushinfo/badger"
	"github.com/Eric-GreenComb/one-pushinfo/bean"
	"github.com/Eric-GreenComb/one-pushinfo/common"
	"github.com/Eric-GreenComb/one-pushinfo/config"
	"github.com/Eric-GreenComb/one-pushinfo/ethereum"
	"github.com/Eric-GreenComb/one-pushinfo/persist"
)

// WriteBlock WriteBlock
func WriteBlock(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_txID, err := SendEthereumCoin(_formParams.Desc)

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

	err = persist.GetPersist().CreateOrder(_order)
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

	_txID, err := SendEthereumCoin(_formParams.Desc)

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

	err = persist.GetPersist().CreateOrder(_order)
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

// SendEthereumCoin SendEthereumCoin
func SendEthereumCoin(desc string) (string, error) {
	txID := &ethcommon.Hash{}

	_from := config.Ethereum.Address
	_to := _from
	_pwd := string(common.FromHex(config.Ethereum.Passphrase))

	_value, err := badger.NewRead().Get(_from)
	if err != nil {
		return txID.String(), err
	}

	var _keystore string
	_keystore = strings.Replace(string(_value), "\\\"", "\"", -1)
	_key, err := keystore.DecryptKey([]byte(_keystore), _pwd)
	if err != nil {
		return txID.String(), err
	}

	_amountBigInt := ethereum.StringToWei("0.01", 18)
	fmt.Println(_amountBigInt)
	_chainIDBigInt := big.NewInt(config.Ethereum.ChainID)

	_nonce, err := ethereum.PendingNonce(_from)
	if err != nil {
		return txID.String(), err
	}

	_inputData := []byte(desc)
	_txID, err := ethereum.SendEthCoins(_to, _nonce, _amountBigInt, _key.PrivateKey, _chainIDBigInt, _inputData)
	if err != nil {
		return txID.String(), err
	}

	return _txID, err
}
