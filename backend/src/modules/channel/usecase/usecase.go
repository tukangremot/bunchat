package usecase

import (
	"context"
	"time"

	gologconstant "github.com/jabardigitalservice/golog/constant"
	"github.com/jabardigitalservice/golog/logger"
	"github.com/tukangremot/bunchat/backend/src/app"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/constant"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/entity"
	loginterface "github.com/tukangremot/bunchat/backend/src/modules/channel/interface/log"
	"github.com/tukangremot/bunchat/backend/src/modules/channel/repository"
)

type (
	UseCaseInterface interface {
		loginterface.UseCaseLogInterface

		CreateChannel(ctx context.Context, data *entity.Channel) (*entity.Channel, error)
	}

	UseCase struct {
		app    *app.App
		repo   repository.RepositoryInterface
		logger *logger.Logger
	}
)

var (
	MethodCreateChannel = "channel:create-channel"
)

func Init(app *app.App, repo repository.RepositoryInterface) UseCaseInterface {
	return &UseCase{
		app:    app,
		repo:   repo,
		logger: app.GetLogger(),
	}
}

func (uc *UseCase) LogSuccess(ctx context.Context, method string, message string, duration time.Duration, additionalInfo map[string]interface{}) {
	data := buildLogData(ctx, method, uc.app.GetVersion(), duration, additionalInfo)

	uc.logger.Info(data, message)

}

func (uc *UseCase) LogError(ctx context.Context, method string, e error, duration time.Duration, additionalInfo map[string]interface{}) {
	data := buildLogData(ctx, method, uc.app.GetVersion(), duration, additionalInfo)

	uc.logger.Error(data, e)
}

func buildLogData(ctx context.Context, method string, version string, duration time.Duration, additionalInfo map[string]interface{}) *logger.LoggerData {
	data := &logger.LoggerData{
		Category:       logger.LoggerUsecase,
		Service:        constant.ServiceName,
		Module:         constant.ModuleName,
		Method:         method,
		Version:        version,
		AdditionalInfo: additionalInfo,
		Duration:       int64(duration),
	}

	if ctx.Value(gologconstant.CtxRequestIDKey) != nil {
		data.RequestID = ctx.Value(gologconstant.CtxRequestIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxRequestNameKey) != nil {
		data.RequestName = ctx.Value(gologconstant.CtxRequestNameKey).(string)
	}

	if ctx.Value(gologconstant.CtxUserIDKey) != nil {
		data.UserID = ctx.Value(gologconstant.CtxUserIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxSessionIDKey) != nil {
		data.SessionID = ctx.Value(gologconstant.CtxSessionIDKey).(string)
	}

	if ctx.Value(gologconstant.CtxClientIDKey) != nil {
		data.ClientID = ctx.Value(gologconstant.CtxClientIDKey).(string)
	}

	return data
}
