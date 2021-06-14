package foodbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"200lab/food-delivery/modules/food/foodmodel"
	"context"
)

type ListFoodByCategoryStore interface {
	ListFoodByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *foodmodel.Filter,
		paging *common.Paging,
		moreInfo ...string,
	) ([]foodmodel.Food, error)
}

type GetCategoryStore interface {
	GetCategoryByCondition(ctx context.Context, condition map[string]interface{}) (*categorymodel.Category, error)
}

type listFoodByCategoryBiz struct {
	store         ListFoodByCategoryStore
	categoryStore GetCategoryStore
}

func NewListFoodByCategoryBiz(store ListFoodByCategoryStore, categoryStore GetCategoryStore) *listFoodByCategoryBiz {
	return &listFoodByCategoryBiz{store: store, categoryStore: categoryStore}
}

func (biz *listFoodByCategoryBiz) ListFoodByCategory(
	ctx context.Context,
	categoryId int,
	paging *common.Paging,
) ([]foodmodel.Food, error) {
	if _, err := biz.categoryStore.GetCategoryByCondition(ctx, map[string]interface{}{"id": categoryId}); err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(categorymodel.EntityName, err)
		}

		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	foods, err := biz.store.ListFoodByCondition(
		ctx,
		map[string]interface{}{"category_id": categoryId},
		nil,
		paging,
		"Restaurant",
	)

	if err != nil {
		return nil, common.ErrCannotListEntity(foodmodel.EntityName, err)
	}

	return foods, nil
}
