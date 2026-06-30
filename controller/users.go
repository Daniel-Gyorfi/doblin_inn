package controller

import "time"

type CreateUserInput struct {
	Username   string    `"json:"username" binding:"required"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}
