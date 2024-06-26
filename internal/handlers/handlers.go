package handlers

import (
	"net/http"

	"github.com/albanybuipe96/bookrestapi/internal/configs"
	"github.com/albanybuipe96/bookrestapi/internal/utils"
)

type DbConfig configs.DbConfig

func Index(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONResponse(w, http.StatusOK, struct{}{})
}

func Error(w http.ResponseWriter, r *http.Request) {
	utils.SendErrorResponse(w, http.StatusBadRequest, "Something went wrong")
}
