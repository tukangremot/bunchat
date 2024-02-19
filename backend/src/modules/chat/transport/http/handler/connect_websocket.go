package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/tukangremot/bunchat/backend/src/modules/chat/endpoint"
	"github.com/tukangremot/gochat"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) ConnectWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user := gochat.NewUser(conn, h.websocketsServer)

	go user.WritePump()
	go user.ReadPump()

	ctx := r.Context()
	result, err := h.endpoint.ClientConnected(ctx, endpoint.ClientConnectedRequest{})
	if err != nil {
		message, err := json.Marshal(result)
		if err != nil {

			user.Send([]byte(err.Error()))

			time.Sleep(time.Second * 1)

			user.GetConn().Close()
		}

		user.Send(message)

		time.Sleep(time.Second * 1)

		user.GetConn().Close()
	}

	for activity := range user.GetActivity() {
		switch activity.Type {
		case gochat.TypeUserActivityChannelConnect:
			// do somthing when user connect to channel
		case gochat.TypeUserActivityGroupJoin:
			// do somthing when user join to group
		case gochat.TypeUserActivityGroupLeave:
			// do somthing when user leave from group
		case gochat.TypeUserActivityMessageSend:
			// do somthing when user send message
		case gochat.TypeUserActivityDisconnect:
			// do somthing when user disconnet
		}
	}
}
