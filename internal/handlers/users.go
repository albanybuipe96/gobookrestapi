package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/albanybuipe96/bookrestapi/internal/database"
	"github.com/albanybuipe96/bookrestapi/internal/utils"
	"github.com/google/uuid"
)

func (dbConfig *DbConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := utils.CreateUserDto{}
	if err := decoder.Decode(&params); err != nil {
		log.Println(err.Error())
		utils.SendErrorResponse(
			w, http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest),
		)
		return
	}

	if err := params.Validate(); err != nil {
		log.Println("Some required fields are empty")
		utils.SendErrorResponse(w,
			http.StatusInternalServerError,
			err.Error(),
		)
		return
	}

	_, err := dbConfig.DB.GetUserByEmail(r.Context(), params.Email)
	if err == nil {
		log.Println("Email already in use")
		utils.SendErrorResponse(
			w,
			http.StatusBadRequest,
			"Email/username already in use",
		)
		return
	}

	dbUser, err := dbConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Username:  params.Username,
		Email:     params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		log.Println(err.Error())
		utils.SendErrorResponse(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("%v", http.StatusText(http.StatusInternalServerError)),
		)
		return
	}

	user := utils.UserDto(dbUser)

	utils.SendJSONResponse(w, http.StatusCreated, user.ToCreateUserDto())
}

func (dbConfig *DbConfig) GetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	usr := utils.UserDto(user)
	utils.SendJSONResponse(w, http.StatusOK, usr.ToGetUserDto())
}
