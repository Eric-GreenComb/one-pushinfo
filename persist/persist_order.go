package persist

import (
	"github.com/Eric-GreenComb/one-pushinfo/bean"
)

// CreateOrder CreateOrder Persist
func (persist *Persist) CreateOrder(order bean.Order) error {
	err := persist.db.Create(&order).Error
	return err
}

// OrderInfo OrderInfo Persist
func (persist *Persist) OrderInfo(orderID string) (bean.Order, error) {

	var order bean.Order

	err := persist.db.Table("orders").Where("order_id = ?", orderID).First(&order).Error

	return order, err
}

// GetAllOrders GetAllOrders Persist
func (persist *Persist) GetAllOrders(catid, patchid string) ([]bean.Order, error) {

	var orders []bean.Order

	err := persist.db.Table("orders").Where("cat_id = ? AND patch_id = ?", catid, patchid).Find(&orders).Error

	return orders, err
}
