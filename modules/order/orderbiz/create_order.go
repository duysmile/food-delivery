package orderbiz

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/modules/food/foodmodel"
	"200lab/food-delivery/modules/order/ordermodel"
	"200lab/food-delivery/pubsub"
	"context"
)

type CreateOrderStore interface {
	CreateOrder(ctx context.Context, order *ordermodel.Order, orderDetails []ordermodel.OrderDetail) error
}

type GetFoodsStore interface {
	ListFoodByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		filter *foodmodel.Filter,
		paging *common.Paging,
		moreInfo ...string,
	) ([]foodmodel.Food, error)
}

type createOrderBiz struct {
	store     CreateOrderStore
	foodStore GetFoodsStore
	pubsub    pubsub.Pubsub
}

func NewCreateOrderBiz(
	store CreateOrderStore,
	foodStore GetFoodsStore,
	pb pubsub.Pubsub,
) *createOrderBiz {
	return &createOrderBiz{
		store:     store,
		foodStore: foodStore,
		pubsub:    pb,
	}
}

func (biz *createOrderBiz) CreateOrder(ctx context.Context, userId int, data *ordermodel.OrderCreate) error {
	ids := make([]int, len(data.Carts))
	quantities := map[int]int{}

	for i := range data.Carts {
		ids[i] = data.Carts[i].FoodId
		quantities[ids[i]] += data.Carts[i].Quantity
	}

	foods, err := biz.foodStore.ListFoodByCondition(ctx, map[string]interface{}{
		"id": ids,
	}, nil, nil, "Restaurant")

	if err != nil {
		return common.ErrInvalidRequest(err)
	}

	totalPrice := calculateTotalPrice(foods, quantities, data.Discount)

	order := ordermodel.Order{
		UserId:     userId,
		TotalPrice: totalPrice,
	}

	orderDetails := make([]ordermodel.OrderDetail, len(foods))
	for i := range foods {
		food := foods[i]
		orderDetails[i] = ordermodel.OrderDetail{
			// FoodOrigin: (*foodmodel.FoodOrigin)(unsafe.Pointer(&foods[i])),
			FoodOrigin: &foodmodel.FoodOrigin{
				Id:             food.Id,
				Name:           food.Name,
				RestaurantName: food.Restaurant.Name,
				Description:    food.Description,
				Images:         food.Images,
			},
			Price:    food.Price,
			Quantity: quantities[food.Id],
			Discount: data.Discount,
		}
	}

	if err = biz.store.CreateOrder(ctx, &order, orderDetails); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}

	// TODO: handle payment before side effect

	// side effect
	foodIds := make([]int, len(foods))
	for i := range foods {
		foodIds[i] = foods[i].Id
	}

	biz.pubsub.Publish(ctx, common.TopicCreateOrder, pubsub.NewMessage(ordermodel.DataPublish{
		UserId:  userId,
		FoodIds: foodIds,
	}))

	return nil
}

func calculateTotalPrice(foods []foodmodel.Food, quantities map[int]int, discount float32) float32 {
	var total float32
	for i := range foods {
		total += foods[i].Price * float32(quantities[foods[i].Id])
	}

	return total * (1 - discount/100)
}
