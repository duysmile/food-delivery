package foodmodel

import (
	"200lab/food-delivery/common"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type FoodOrigin struct {
	Id             int            `json:"-" gorm:"-"`
	RestaurantName string         `json:"restaurant" gorm:"-"`
	Name           string         `json:"name" gorm:"-"`
	Description    string         `json:"description,omitempty" gorm:"-"`
	Images         *common.Images `json:"images" gorm:"-"`
}

func (f *FoodOrigin) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var foodOrigin FoodOrigin
	if err := json.Unmarshal(bytes, &foodOrigin); err != nil {
		return err
	}

	*f = foodOrigin
	return nil
}

func (f *FoodOrigin) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}

	return json.Marshal(f)
}
