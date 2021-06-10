package categorybiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/category/categorymodel"
	"context"
)

type ListCategoryStore interface {
	ListCategoryByCondition(
		ctx context.Context,
		condition map[string]interface{},
		paging *common.Paging,
		moreKeys ...string,
	) ([]categorymodel.Category, error)
}

type listCategoryBiz struct {
	store ListCategoryStore
}

func NewListCategoryBiz(store ListCategoryStore) *listCategoryBiz {
	return &listCategoryBiz{store: store}
}

func (biz *listCategoryBiz) ListCategory(ctx context.Context, paging *common.Paging) ([]categorymodel.Category, error) {
	categories, err := biz.store.ListCategoryByCondition(ctx, nil, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(categorymodel.EntityName, err)
	}

	return categories, nil
}
