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

func (dbConfig *DbConfig) AddBook(w http.ResponseWriter, r *http.Request, user database.User) {

	decoder := json.NewDecoder(r.Body)
	params := utils.CreateBookDto{}

	if err := decoder.Decode(&params); err != nil {
		log.Println(err.Error())
		utils.SendErrorResponse(w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
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

	bk, err := dbConfig.DB.CreateBook(r.Context(), database.CreateBookParams{
		ID:          uuid.New(),
		Title:       params.Title,
		Author:      params.Author,
		Description: params.Description,
		Pages:       params.Pages,
		UserID:      user.ID,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
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

	book := utils.BookDto(bk)
	utils.SendJSONResponse(w, http.StatusCreated, book.ToCreateBookDto())
}

func (dbConfig *DbConfig) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := dbConfig.DB.GetBooks(r.Context())
	if err != nil {
		log.Println(err.Error())
		utils.SendErrorResponse(w,
			http.StatusInternalServerError,
			http.StatusText(http.StatusInternalServerError),
		)
	}

	utils.SendJSONResponse(w, http.StatusOK, books)
}
