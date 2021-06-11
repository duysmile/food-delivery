package categorybiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
)

type UpdateCategoryStore interface {
	UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error
	GetCategoryByCondition(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
}

type updateCategoryBiz struct {
	store UpdateCategoryStore
}

func NewUpdateCategoryBiz(store UpdateCategoryStore) *updateCategoryBiz {
	return &updateCategoryBiz{store: store}
}

func (biz *updateCategoryBiz) UpdateCategory(ctx context.Context, id int, data *categorymodel.CategoryUpdate) error {
	if _, err := biz.store.GetCategoryByCondition(ctx, map[string]interface{}{"id": id}); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	if err := biz.store.UpdateCategory(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}

	return nil
}
