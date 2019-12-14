package handler

import (
	"net/http"
	"os"
)

func HelloHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello" + os.Getenv("APP_NAME")))
	}
}
