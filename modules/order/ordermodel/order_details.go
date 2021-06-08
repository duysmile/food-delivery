package ordermodel

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
)

type OrderDetail struct {
	common.SQLModel `json:",inline"`
	OrderId         int                   `json:"-" gorm:"column:order_id;"`
	FoodOrigin      *foodmodel.FoodOrigin `json:"food" gorm:"column:food_origin;"`
	Price           float32               `json:"price" gorm:"column:price;"`
	Quantity        int                   `json:"quantity" gorm:"column:quantity;"`
	Discount        float32               `json:"discount" gorm:"column:discount;"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}
