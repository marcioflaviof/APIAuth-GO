package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//User just the user
type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name     string             `json:"name" validate:"required,min=6"`
	Email    string             `json:"email" validate:"required,email"`
	Password string             `json:"password" validate:"required,min=6"`
	Date     time.Time
}

type LoginUser struct {
	ID       primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email    string             `json:"email" validate:"required"`
	Password string             `json:"password" validate:"required"`
}
