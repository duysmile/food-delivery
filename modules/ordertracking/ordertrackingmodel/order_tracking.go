package ordertrackingmodel

const EntityName = "OrderTracking"

type OrderTrackingCreate struct {
	OrderId int        `json:"order_id" gorm:"column:order_id;"`
	State   OrderState `json:"-" gorm:"column:state;"`
}

func (OrderTrackingCreate) TableName() string {
	return "order_trackings"
}

func (o OrderTrackingCreate) GetOrderId() int {
	return o.OrderId
}

func (o OrderTrackingCreate) GetState() OrderState {
	return o.State
}
