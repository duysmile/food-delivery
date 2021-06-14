// Business Model
package restaurantmodel

import (
	"200lab/food-delivery/common"
	"strings"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel  `json:",inline"`
	OwnerId          int                `json:"owner_id" gorm:"column:owner_id;"`
	Owner            *common.SimpleUser `json:"owner" gorm:"preload:false;foreignKey:OwnerId;"`
	Name             string             `json:"name" gorm:"column:name;"`
	Addr             string             `json:"address" gorm:"column:addr;"`
	CityId           int                `json:"city_id" gorm:"column:city_id;"`
	Lat              float32            `json:"lat" gorm:"column:lat;"`
	Long             float32            `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64            `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
	Cover            *common.Images     `json:"cover" gorm:"column:cover;"`
	Logo             *common.Image      `json:"logo" gorm:"column:logo;"`
	LikeCount        int                `json:"liked_count" gorm:"column:liked_count;"` // computed field
}

type SimpleRestaurant struct {
	Id               int     `json:"-" gorm:"column:id;"`
	OwnerId          int     `json:"-" gorm:"column:owner_id;"`
	Name             string  `json:"name" gorm:"column:name;"`
	Addr             string  `json:"address" gorm:"column:addr;"`
	CityId           int     `json:"city_id" gorm:"column:city_id;"`
	Lat              float32 `json:"lat" gorm:"column:lat;"`
	Long             float32 `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64 `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (SimpleRestaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel  `json:",inline"`
	OwnerId          int            `json:"owner_id" gorm:"column:owner_id;"`
	Name             string         `json:"name" gorm:"column:name;"`
	Addr             string         `json:"address" gorm:"column:addr;"`
	CityId           int            `json:"city_id" gorm:"column:city_id;"`
	Lat              float32        `json:"lat" gorm:"column:lat;"`
	Long             float32        `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64        `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
	Cover            *common.Images `json:"cover" gorm:"column:cover;"`
	Logo             *common.Image  `json:"logo" gorm:"column:logo;"`
}

type RestaurantUpdate struct {
	Name             string         `json:"name" gorm:"column:name;"`
	Addr             string         `json:"address" gorm:"column:addr;"`
	CityId           int            `json:"city_id" gorm:"column:city_id;"`
	Lat              float32        `json:"lat" gorm:"column:lat;"`
	Long             float32        `json:"lng" gorm:"column:lng;"`
	ShippingFeePerKm float64        `json:"shipping_fee_per_km" gorm:"column:shipping_fee_per_km;"`
	Cover            *common.Images `json:"cover" gorm:"column:cover;"`
	Logo             *common.Image  `json:"logo" gorm:"column:logo;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func (res *RestaurantCreate) Validate() error {
	res.Name = strings.TrimSpace(res.Name)
	if len(res.Name) == 0 {
		return ErrNameCannotBeEmpty
	}

	return nil
}

var (
	ErrNameCannotBeEmpty = common.NewCustomError(nil, "restaurant can't be blank", "ErrNameCannotBeBlank")
)

func (data *Restaurant) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypeRestaurant)

	if data.Owner != nil {
		data.Owner.Mask(isAdminOrOwner)
	}
}
