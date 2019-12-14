package middleware

import (
	"context"
	"map-friend/src/internal/logger"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func ReqIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := ""
		if id = r.Header.Get("UUID"); id == "" {
			id = uuid.New().String()
		}
		ctx := r.Context()
		r = r.WithContext(context.WithValue(ctx, os.Getenv("APP_REQUEST_ID"), id))

		log := logger.GetLoggerWithCtx(r.Context())
		log.Debugf("Incomming request %s %s %s", r.Method, r.RequestURI, r.RemoteAddr)

		next.ServeHTTP(w, r)

		log.Debugf("Finished handling request")
	})
}
