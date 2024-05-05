package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/albanybuipe96/bookrestapi/internal/database"
	"github.com/albanybuipe96/bookrestapi/utils"
	"github.com/google/uuid"
)

func (dbConfig *DbConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	if err := decoder.Decode(&params); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := dbConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  params.Username,
		Email:     params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	utils.SendJSONResponse(w, http.StatusCreated, utils.UserDto(user))
}
