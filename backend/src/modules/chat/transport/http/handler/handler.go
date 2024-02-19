package handler

import (
	"net/http"

	gorrillawebsocket "github.com/gorilla/websocket"
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/chat/endpoint"
	"github.com/tukangremot/gochat"
)

type Handler struct {
	app              *app.App
	endpoint         endpoint.EndpointInterface
	wsUpgrader       gorrillawebsocket.Upgrader
	websocketsServer *gochat.Server
}

func InitHandler(app *app.App, endpoint endpoint.EndpointInterface) *Handler {
	server := gochat.NewServer(&gochat.Server{
		PubSub:  gochat.NewPubSub(gochat.PubSubDriverRedis, app.GetRedis()),
		Session: gochat.NewSession(gochat.SessionDriverRedis, app.GetRedis()),
	})
	go server.Run()

	handler := &Handler{
		app:      app,
		endpoint: endpoint,
		wsUpgrader: gorrillawebsocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		websocketsServer: server,
	}

	return handler
}
