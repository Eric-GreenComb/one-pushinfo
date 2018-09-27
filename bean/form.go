package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params string `form:"params" json:"params"` // params
	Key    string `form:"key" json:"key"`       // key
	Value  string `form:"value" json:"value"`   // value

	OrderCode string `form:"ordercode" json:"ordercode"` // 订单编码
	Amount    string `form:"amount" json:"amount"`       //
	GoodsID   string `form:"goodsid" json:"goodsid"`     // 货物id
	GoodName  string `form:"goodname" json:"goodname"`   // Iphone(第三期）
	BuyTime   string `form:"buytime" json:"buytime"`     // 购买时间
	UserName  string `form:"username" json:"username"`   // 购买用户名称
	Type      int8   `form:"type" json:"type"`           // 类型  0为下单，1为抽奖

	WinTime string `form:"wintime" json:"wintime"` // 开奖时间
}
