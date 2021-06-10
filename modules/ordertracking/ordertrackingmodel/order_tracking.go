package ordertrackingmodel

import "200lab/food-delivery/common"

const EntityName = "OrderTracking"

type OrderTracking struct {
	common.SQLModel `json:",inline"`
	OrderId         int        `json:"-" gorm:"column:order_id;"`
	State           OrderState `json:"state" gorm:"column:state;"`
}

type OrderTrackingCreate struct {
	FakeOrderId *common.UID `json:"order_id" gorm:"-"`
	OrderId     int         `json:"-" gorm:"column:order_id;"`
	State       OrderState  `json:"state" gorm:"column:state;"`
}

type OrderTrackingUpdate struct {
	State OrderState `json:"-" gorm:"column:state"`
}

func (OrderTracking) TableName() string {
	return "order_trackings"
}

func (OrderTrackingCreate) TableName() string {
	return "order_trackings"
}

func (OrderTrackingUpdate) TableName() string {
	return "order_trackings"
}

func (o OrderTrackingCreate) GetOrderId() int {
	return o.OrderId
}

func (o OrderTrackingCreate) GetState() OrderState {
	return o.State
}
