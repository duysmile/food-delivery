package useraddressmodel

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/city/citymodel"
)

const EntityName = "User address"

type UserAddress struct {
	common.SQLModel `json:",inline"`
	UserId          int            `json:"-" gorm:"column:user_id;"`
	FakeUserId      *common.UID    `json:"user_id" gorm:"-"`
	CityId          int            `json:"-" gorm:"column:city_id;"`
	City            citymodel.City `json:"city" gorm:"preload:false;"`
	Title           string         `json:"title,omitempty" gorm:"column:city;"`
	Icon            *common.Image  `json:"icon,omitempty" gorm:"column:icon;"`
	Address         string         `json:"address" gorm:"column:addr;"`
	Lat             *float32       `json:"lat,omitempty" gorm:"column:lat;"`
	Lng             *float32       `json:"lng,omitempty" gorm:"column:lng;"`
}

type UserAddressCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int           `json:"-" gorm:"column:user_id;"`
	CityId          int           `json:"city_id" gorm:"column:city_id;"`
	Title           *string       `json:"title,omitempty" gorm:"column:title;"`
	Icon            *common.Image `json:"icon,omitempty" gorm:"column:icon;"`
	Address         string        `json:"address" gorm:"column:addr;"`
	Lat             *float32      `json:"lat,omitempty" gorm:"column:lat;"`
	Lng             *float32      `json:"lng,omitempty" gorm:"column:lng;"`
}

func (UserAddress) TableName() string {
	return "user_addresses"
}

func (UserAddressCreate) TableName() string {
	return "user_addresses"
}

func (u *UserAddress) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUserAddress)
	fakeUserId := common.NewUID(uint32(u.UserId), common.DbTypeUser, 1)
	u.FakeUserId = &fakeUserId
}
