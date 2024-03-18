package utils

import (
	"context"

	"github.com/sirupsen/logrus"
)

type ContextKey string

const (
	REQUEST_ID_KEY = ContextKey("request_id")
	LOGGER_KEY     = ContextKey("logger")
)

func GetContextLogger(ctx context.Context) *logrus.Entry {
	logger, ok := ctx.Value(LOGGER_KEY).(*logrus.Entry)
	if !ok {
		defaultLogger := logrus.New()
		defaultLogger.SetLevel(logrus.InfoLevel)
		return defaultLogger.WithField("default", true)
	}
	return logger
}

func GetRequestIDFromCtx(ctx context.Context) string {
	return ctx.Value(REQUEST_ID_KEY).(string)
}

func UpdateCtxLoggerWithMethod(ctx context.Context, methodName string) context.Context {
	contextLogger := GetContextLogger(ctx)
	newContextLogger := contextLogger.WithFields(logrus.Fields{
		"method": methodName,
	})
	return context.WithValue(ctx, LOGGER_KEY, newContextLogger)
}
