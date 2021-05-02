package restaurantlikemodel

import "time"

const EntityName = "RestaurantLike"

type Like struct {
	RestaurantId int        `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int        `json:"user_id" gorm:"column:user_id;"`
	CreatedAt    *time.Time `json:"created_at" gorm:"column:created_at;"`
}

type LikeCreate struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id;"`
	UserId       int `json:"user_id" gorm:"column:user_id;"`
}

type LikeDelete struct {
	RestaurantId int `json:"restaurant_id" gorm:"column:restaurant_id;primaryKey;"`
	UserId       int `json:"user_id" gorm:"column:user_id;primaryKey;"`
}

func (Like) TableName() string {
	return "restaurant_likes"
}

func (LikeCreate) TableName() string {
	return "restaurant_likes"
}

func (LikeDelete) TableName() string {
	return "restaurant_likes"
}
