package categorybiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
)

type DeleteCategoryStore interface {
	DeleteCategory(ctx context.Context, id int) error
	GetCategoryByCondition(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
}

type deleteCategoryBiz struct {
	store DeleteCategoryStore
}

func NewDeleteCategoryBiz(store DeleteCategoryStore) *deleteCategoryBiz {
	return &deleteCategoryBiz{store: store}
}

func (biz *deleteCategoryBiz) DeleteCategory(ctx context.Context, id int) error {
	if _, err := biz.store.GetCategoryByCondition(ctx, map[string]interface{}{
		"id": id,
	}); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	if err := biz.store.DeleteCategory(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	return nil
}
