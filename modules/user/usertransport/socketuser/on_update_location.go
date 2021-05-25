package socketuser

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

type LocationData struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func OnUserUpdateLocation(appCtx component.AppContext, requester common.Requester) func(s socketio.Conn, location LocationData) {
	return func(s socketio.Conn, location LocationData) {
		log.Println("User update location: user id is", requester.GetUserId(), " at location: ", location)
	}
}
