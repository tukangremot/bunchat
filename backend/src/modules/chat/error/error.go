package error

import (
	"net/http"

	goliatherrors "github.com/fazpass/goliath/v3/errors"
)

var (
	ErrPayloadValidation = goliatherrors.New(http.StatusBadRequest, "00", "Invalid payload")
	ErrGeneral           = goliatherrors.New(http.StatusInternalServerError, "00", "General error")
	ErrFeatureInactive   = goliatherrors.New(http.StatusForbidden, "00", "Cannot access this feature")
)
