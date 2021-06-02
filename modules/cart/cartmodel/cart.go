package cartmodel

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"time"
)

const EntityName = "Cart"

type Cart struct {
	UserId    int            `json:"-" gorm:"column:user_id;"`
	FoodId    int            `json:"-" gorm:"column:food_id;"`
	Food      foodmodel.Food `json:"food" gorm:"preload:false;"`
	Quantity  int            `json:"quantity" gorm:"column:quantity;"`
	Status    int            `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:updated_at;"`
}

type CartCreate struct {
	UserId     int         `json:"-" gorm:"column:user_id;"`
	FoodId     int         `json:"-" gorm:"column:food_id;"`
	FakeFoodId *common.UID `json:"food_id" gorm:"-"`
	Quantity   int         `json:"quantity" gorm:"column:quantity;"`
}

type CartUpdate struct {
	FoodId     int         `json:"-" gorm:"column:food_id;"`
	FakeFoodId *common.UID `json:"food_id" gorm:"-"`
	Quantity   int         `json:"quantity" gorm:"column:quantity;"`
}

func (Cart) TableName() string {
	return "carts"
}

func (CartCreate) TableName() string {
	return "carts"
}

func (CartUpdate) TableName() string {
	return "carts"
}
