package citymodel

import "200lab/food-delivery/common"

type City struct {
	common.SQLModel `json:"-"`
	Title           string `json:"title" gorm:"column:title;"`
}
