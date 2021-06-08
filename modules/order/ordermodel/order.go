package ordermodel

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/cart/cartmodel"
	"200lab/food-delivery/modules/ordertracking/ordertrackingmodel"
	"errors"
)

const EntityName = "Order"

type Order struct {
	common.SQLModel `json:",inline"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	TotalPrice      float32            `json:"total_price" gorm:"column:total_price;"`
	ShipperId       int                `json:"-" gorm:"column:shipper_id;"`
	Shipper         *common.SimpleUser `json:"shipper" gorm:"preload:false;"`
}

type OrderCreate struct {
	Carts    []cartmodel.CartItem `json:"carts" gorm:"-"`
	Discount float32              `json:"discount" gorm:"-"`
}

type OrderUpdate struct {
	ShipperId *int `json:"-" gorm:"column:shipper_id"`
	Status    *int `json:"-" gorm:"column:status"`
}

type DataPublish struct {
	UserId  int
	FoodIds []int
	OrderId int
	State   ordertrackingmodel.OrderState
}

func (d DataPublish) GetOrderId() int {
	return d.OrderId
}

func (d DataPublish) GetFoodIds() []int {
	return d.FoodIds
}

func (d DataPublish) GetUserId() int {
	return d.UserId
}

func (d DataPublish) GetState() ordertrackingmodel.OrderState {
	return d.State
}

func (Order) TableName() string {
	return "orders"
}

func (OrderUpdate) TableName() string {
	return "orders"
}

func (o *OrderCreate) Unmask() {
	for i := range o.Carts {
		o.Carts[i].FoodId = int(o.Carts[i].FakeFoodId.GetLocalID())
	}
}

func (o *OrderCreate) Validate() error {
	if len(o.Carts) > 100 {
		return errors.New("too much items in order")
	}

	return nil
}
