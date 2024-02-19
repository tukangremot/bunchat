package main

import (
	"fmt"
	"log"
	"net/http"

	helperresponse "github.com/fazpass/goliath/v3/helper/http/response"
	middlewaregoliath "github.com/fazpass/goliath/v3/middleware"
	"github.com/go-chi/chi/middleware"
	gologconstant "github.com/jabardigitalservice/golog/constant"
	middlewareheadertoctx "github.com/jabardigitalservice/utilities-go/middleware/http/headers-to-ctx"
	"github.com/spf13/viper"
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/constant"
	"github.com/tukangremot/bunchat/backend/src/modules"
)

func main() {
	var (
		app    = app.Init()
		router = app.GetHttpRouter()
		module = modules.Init(app)
	)

	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/"))
	router.Use(middleware.CleanPath)
	router.Use(middlewareheadertoctx.Mapping(map[string]interface{}{
		gologconstant.CtxSessionIDKey: constant.HeaderSessionIDKey,
		gologconstant.CtxClientIDKey:  constant.HeaderClientIDKey,
		gologconstant.CtxUserIDKey:    constant.HeaderUserIDKey,
	}))

	router.Use(middlewaregoliath.Recoverer(middlewaregoliath.RecovererOptions{
		Debug: viper.GetBool("APP_DEBUG"),
		Response: middlewaregoliath.RecovererResponse{
			Status: http.StatusInternalServerError,
			Body: helperresponse.Response{
				Status:  false,
				Message: "Error",
				Code:    fmt.Sprint(http.StatusInternalServerError, constant.ServiceCode, "00"),
			},
		},
	}))

	router.Get("/v1/chat/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ping")
	})

	router.Mount("/v1/chat", module.Chat.GetHttpRouter())
	router.Mount("/v1/chat/channels", module.Channel.GetHttpRouter())

	var err = app.RunHttp()
	if err != nil {
		log.Panic(err)
	}
}
