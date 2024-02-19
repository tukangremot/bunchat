package log

import (
	"context"
	"time"

	"github.com/jabardigitalservice/golog/logger"
)

type UseCaseLogInterface interface {
	LogSuccess(ctx context.Context, method string, message string, duration time.Duration, additionalInfo map[string]interface{})
	LogError(ctx context.Context, method string, err error, duration time.Duration, additionalInfo map[string]interface{})
}

type ExternalLogInterface interface {
	LogSuccess(ctx context.Context, method string, duration time.Duration,
		externalInfo *logger.ExternalLoggerData, message string, additionalInfo map[string]interface{})
	LogError(ctx context.Context, method string, duration time.Duration,
		externalInfo *logger.ExternalLoggerData, err error, additionalInfo map[string]interface{})
}
