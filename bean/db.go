package bean

import (
	"github.com/jinzhu/gorm"
)

// Order 下单
type Order struct {
	gorm.Model
	OrderID string `form:"orderid" json:"orderid"` // 订单编码
	Amount  string `form:"amount" json:"amount"`   // 订单金额
	CatID   string `form:"catid" json:"catid"`     // 货物id
	PatchID string `form:"patchid" json:"patchid"` // 期数ID
	BuyTime string `form:"buytime" json:"buytime"` // 购买时间
	Account string `form:"account" json:"account"` // 账户名称
	Mobile  string `form:"mobile" json:"mobile"`   // 账户手机/邮箱
	Type    int8   `form:"type" json:"type"`       // 类型  0为下单，1为抽奖
	Desc    string `form:"desc" json:"desc"`       // 备注
	TxID    string `form:"txid" json:"txid"`       // 入账txid
}
