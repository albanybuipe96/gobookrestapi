package utils

import (
	"errors"
	"time"

	"github.com/albanybuipe96/bookrestapi/internal/database"
	"github.com/google/uuid"
)

type CreateBookDto struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Pages       int32  `json:"pages"`
}

func (book *CreateBookDto) Validate() error {
	if book.Title == "" && book.Author == "" && book.Description == "" {
		return errors.New("title, author, description are required fields")
	}

	if book.Title == "" {
		return errors.New("title is a required field")
	}

	if book.Author == "" {
		return errors.New("author is a required field")
	}

	if book.Description == "" {
		return errors.New("description is a required field")
	}
	return nil
}

type BookResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Pages       int32     `json:"pages"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      uuid.UUID `json:"user_id"`
}

type GetBookDto struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description"`
	Pages       int32     `json:"pages"`
	UserID      uuid.UUID `json:"user_id"`
}

type BookDto database.Book

func (book *BookDto) ToCreateBookDto() BookResponseDto {
	return BookResponseDto{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		Description: book.Description,
		Pages:       book.Pages,
		UserID:      book.UserID,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}
}
