package handlers

import (
	"log"
	"net/http"

	"github.com/albanybuipe96/bookrestapi/internal/auth"
	"github.com/albanybuipe96/bookrestapi/internal/database"
	"github.com/albanybuipe96/bookrestapi/internal/utils"
)

type AuthHandler func(http.ResponseWriter, *http.Request, database.User)

func (dbConfig *DbConfig) AuthMiddleware(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.ExtractAPIKey(r.Header)
		if err != nil {
			log.Println(err.Error())
			utils.SendErrorResponse(w, http.StatusForbidden, err.Error())
			return
		}

		user, err := dbConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			log.Println(err.Error())
			utils.SendErrorResponse(w, http.StatusNotFound, "No user found with given credentials")
			return
		}

		handler(w, r, user)
	}
}
