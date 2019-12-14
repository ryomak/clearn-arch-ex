package handler

import (
	"encoding/json"
	"strconv"

	"net/http"

	"map-friend/src/application/usecase"
	"map-friend/src/infrastructure/repository"
	"map-friend/src/internal/logger"

	"github.com/gorilla/mux"
)

func GetUserHandler() func(w http.ResponseWriter, r *http.Request) {
	uUsecase := usecase.NewIUserUseCase(
		repository.NewIUserRepository(),
	)
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.GetLoggerWithCtx(r.Context())

		vars := mux.Vars(r)
		userID, err := strconv.Atoi(vars["id"])
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "id not number"})
			return
		}
		logger.Info("Start usecase:", userID)
		user, err := uUsecase.GetUserByID(r.Context(), uint(userID))
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error": "notfound"})
			return
		}
		logger.Info("End usecase:", userID)
		json.NewEncoder(w).Encode(user)
	}
}
