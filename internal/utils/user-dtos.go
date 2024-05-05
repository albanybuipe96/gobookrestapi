package utils

import (
	"errors"
	"time"

	"github.com/albanybuipe96/bookrestapi/internal/database"
	"github.com/google/uuid"
)

type CreateUserDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (user *CreateUserDto) Validate() error {

	if user.Email == "" && user.Username == "" {
		return errors.New("email, username are required fields")
	}

	if user.Email == "" {
		return errors.New("email is a required field")
	}

	if user.Username == "" {
		return errors.New("username is a required field")
	}
	return nil
}

type UserResponseDto struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

type GetUserDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

type UserDto database.User

func (user *UserDto) ToCreateUserDto() UserResponseDto {
	return UserResponseDto{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

func (user *UserDto) ToGetUserDto() GetUserDto {
	return GetUserDto{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
