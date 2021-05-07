package foodmodel

import (
	"200lab/food-delivery/common"
)

const EntityName = "Food"

type Food struct {
	common.SQLModel  `json:",inline"`
	RestaurantId     int            `json:"-" gorm:"column:restaurant_id;"`
	FakeRestaurantId *common.UID    `json:"restaurant_id" gorm:"-"`
	CategoryId       int            `json:"-" gorm:"column:category_id;"`
	FakeCategoryId   *common.UID    `json:"category_id,omitempty" gorm:"-"`
	Name             string         `json:"name" gorm:"column:name;"`
	Description      string         `json:"description,omitempty" gorm:"column:description;"`
	Price            float32        `json:"price" gorm:"column:price;"`
	Images           *common.Images `json:"images" column:"images"`
}

type FoodCreate struct {
	RestaurantId     int            `json:"-" gorm:"column:restaurant_id;"`
	FakeRestaurantId string         `json:"restaurant_id" gorm:"-"`
	CategoryId       *int           `json:"-" gorm:"column:category_id;"`
	FakeCategoryId   string         `json:"category_id,omitempty" gorm:"-"`
	Name             string         `json:"name" gorm:"column:name;"`
	Description      string         `json:"description,omitempty" gorm:"column:description;"`
	Price            float32        `json:"price" gorm:"column:price;"`
	Images           *common.Images `json:"images" column:"images"`
}

func (Food) TableName() string {
	return "foods"
}

func (FoodCreate) TableName() string {
	return "foods"
}

func (f *Food) Mask(isAdminOrOwner bool) {
	f.GenUID(common.DbTypeFood)
	fakeRestaurantId := common.NewUID(uint32(f.RestaurantId), common.DbTypeFood, 1)
	f.FakeRestaurantId = &fakeRestaurantId
	if f.CategoryId != 0 {
		fakeCategoryUID := common.NewUID(uint32(f.CategoryId), common.DbTypeFood, 1)
		f.FakeCategoryId = &fakeCategoryUID
	}
}

func (f *FoodCreate) Unmask() error {
	uid, error := common.FromBase58(f.FakeRestaurantId)
	if error != nil {
		return common.ErrWrongUID
	}

	f.RestaurantId = int(uid.GetLocalID())

	if f.FakeCategoryId != "" {
		uid, error := common.FromBase58(f.FakeCategoryId)
		if error != nil {
			return common.ErrWrongUID
		}

		categoryId := int(uid.GetLocalID())
		f.CategoryId = &categoryId
	}

	return nil
}
