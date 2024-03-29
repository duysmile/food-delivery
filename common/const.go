package common

const (
	DbTypeRestaurant    = 1
	DbTypeUser          = 2
	DbTypeFood          = 3
	DbTypeUserAddress   = 4
	DbTypeOrder         = 5
	DbTypeOrderTracking = 6
	DbTypeCategory      = 6
)

const CurrentUser = "user"

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
