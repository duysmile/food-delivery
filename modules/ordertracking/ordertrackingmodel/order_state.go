package ordertrackingmodel

import (
	"database/sql/driver"
)

type OrderState string

const (
	WaitingForShipper OrderState = "waiting_for_shipper"
	Preparing         OrderState = "preparing"
	OnTheWay          OrderState = "on_the_way"
	Deliveried        OrderState = "delivered"
	Cancel            OrderState = "cancel"
)

func (o *OrderState) Scan(value interface{}) error {
	*o = OrderState(value.([]byte))
	return nil
}

func (o *OrderState) Value() (driver.Value, error) {
	return string(*o), nil
}
