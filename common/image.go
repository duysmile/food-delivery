package common

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	Width     int    `json:"width" gorm:"column:width;"`
	Height    int    `json:"height" gorm:"column:height"`
	CloudName string `json:"cloud_name,omitempty" gorm:"-"`
	Extension string `json:"extension,omitempty" gorm:"-"`
}

func (Image) TableName() string { return "images" }

func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

type Images []Image

func (i *Images) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value: ", value))
	}

	var imgs []Image
	if err := json.Unmarshal(bytes, &imgs); err != nil {
		return err
	}

	*i = imgs
	return nil
}

func (i *Images) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}

	return json.Marshal(i)
}

type ImageStore interface {
	FindImageByCondition(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*Image, error)
}

type ImagesStore interface {
	ListImages(ctx context.Context, ids []int, moreKeys ...string) ([]Image, error)
}

func (i *Image) Validate(ctx context.Context, imageStore ImageStore) error {
	if i != nil {
		image, _ := imageStore.FindImageByCondition(ctx, map[string]interface{}{"id": i.Id})
		return i.ValidateData(image)
	}

	return nil
}

func (i *Image) ValidateData(j *Image) error {
	if j.Url != i.Url {
		return ErrInvalidRequest(errors.New("image is not existed"))
	}

	return nil
}

func (i *Images) Validate(ctx context.Context, imagesStore ImagesStore) error {
	images := []Image(*i)
	ids := make([]int, len(images))

	for index := range images {
		ids[index] = images[index].Id
	}

	if i != nil {
		listImage, _ := imagesStore.ListImages(ctx, ids)
		if len(listImage) != len(ids) {
			return ErrInvalidRequest(errors.New("image list is not enough"))
		}

		for index := range listImage {
			if err := listImage[index].ValidateData(&images[index]); err != nil {
				return err
			}
		}
	}

	return nil
}
