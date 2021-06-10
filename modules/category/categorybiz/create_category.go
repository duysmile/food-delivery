package categorybiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
	"errors"
)

type CreateCategoryStore interface {
	CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error
	GetCategoryByCondition(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
}

type createCategoryBiz struct {
	store CreateCategoryStore
}

func NewCreateCategoryBiz(store CreateCategoryStore) *createCategoryBiz {
	return &createCategoryBiz{store: store}
}

func (biz *createCategoryBiz) CreateCategoryBiz(ctx context.Context, data *categorymodel.CategoryCreate) error {
	_, err := biz.store.GetCategoryByCondition(ctx, map[string]interface{}{
		"name": data.Name,
	})

	if err != nil && err != common.RecordNotFound {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	} else if err == nil {
		return common.ErrEntityExisted(categorymodel.EntityName, errors.New("category name is existed"))
	}

	if err := biz.store.CreateCategory(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}

	return nil
}
