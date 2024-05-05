package utils

import (
	"time"

	"github.com/albanybuipe96/bookrestapi/internal/database"
	"github.com/google/uuid"
)

type UserDto struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User database.User

func (user *User) ToUserDto() UserDto {
	return UserDto{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
