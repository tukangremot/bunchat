package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	helperresponse "github.com/fazpass/goliath/v3/helper/http/response"
	"github.com/go-chi/render"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/constant"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/endpoint"
	_error "github.com/tukangremot/bunchat/backend/src/modules/channel/error"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/transport/http/response"
)

func (h *Handler) CreateChannel(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		result  interface{}
		reqData *endpoint.CreateChannelRequest
	)

	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		response.RenderErrResponse(w, r, _error.ErrInvalidRequest, nil)
		return
	}

	result, err = h.endpoint.CreateChannel(ctx, reqData)
	if err != nil {
		response.RenderErrResponse(w, r, err, result)
		return
	}

	err = render.Render(w, r, helperresponse.Build(&helperresponse.Response{
		Code:    fmt.Sprint(http.StatusOK, constant.ServiceCode, "00"),
		Data:    result,
		Message: "success",
		Errors:  nil,
		Status:  true,
	}))
	if err != nil {
		panic(err)
	}
}
