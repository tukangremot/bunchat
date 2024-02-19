package response

import (
	"fmt"
	"net/http"

	goliatherrors "github.com/fazpass/goliath/v3/errors"
	helperresponse "github.com/fazpass/goliath/v3/helper/http/response"
	"github.com/go-chi/render"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/constant"
	_error "github.com/tukangremot/bunchat/backend/src/modules/channel/error"
)

type Response struct {
	Status  bool        `json:"status"`
	Errors  interface{} `json:"errors,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RenderErrResponse(w http.ResponseWriter, r *http.Request, err error, errorDetails interface{}) {
	var newErr, isCustom = err.(*goliatherrors.Error)
	if !isCustom {
		newErr = _error.ErrGeneral.(*goliatherrors.Error)
	}

	render.Status(r, newErr.GetHttpStatus())

	err = render.Render(w, r, helperresponse.Build(
		&helperresponse.Response{
			Code:    fmt.Sprint(newErr.HttpStatus, constant.ServiceCode, newErr.CaseCode),
			Data:    newErr.Data,
			Message: newErr.Message,
			Errors:  errorDetails,
			Status:  false,
		}))
	if err != nil {
		panic(err)
	}
}
