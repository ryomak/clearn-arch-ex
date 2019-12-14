package logger

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

func GetLoggerWithCtx(ctx context.Context) *log.Entry {
	return log.WithField(os.Getenv("APP_REQUEST_ID"), ctx.Value(os.Getenv("APP_REQUEST_ID")).(string))
}
