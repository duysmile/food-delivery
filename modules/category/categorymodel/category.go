package categorymodel

import "200lab/food-delivery/common"

const EntityName = "Category"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     *string       `json:"description,omitempty" gorm:"column:description;"`
	Icon            *common.Image `json:"icon,omitempty" gorm:"column:icon;"`
}

type CategoryCreate struct {
	Name        string        `json:"name" gorm:"column:name;"`
	Description *string       `json:"description,omitempty" gorm:"column:description;"`
	Icon        *common.Image `json:"icon,omitempty" gorm:"column:icon;"`
}

type CategoryUpdate struct {
	Name        string        `json:"name" gorm:"column:name;"`
	Description *string       `json:"description,omitempty" gorm:"column:description;"`
	Icon        *common.Image `json:"icon,omitempty" gorm:"column:icon;"`
}

func (Category) TableName() string {
	return "categories"
}

func (CategoryCreate) TableName() string {
	return "categories"
}

func (CategoryUpdate) TableName() string {
	return "categories"
}
