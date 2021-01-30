package database

import (
	"context"
	"log"
	"restapi/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(u model.User) (err error) {

	c := Db.Collection("users")

	user, err := c.InsertOne(context.TODO(), u)

	if err != nil {
		log.Printf("[ERROR] Error in insert user: %v, %v", err, u)
	}

	log.Printf("[INFO] User inserted %v", user.InsertedID)

	return

}

func SearchUser(email string) (u model.User, err error) {

	c := Db.Collection("users")

	fnd := c.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: email}})
	err = fnd.Decode(&u)

	if err != nil {
		log.Printf("[ERROR] problem searching user: %v %v", err, u)
		return
	}

	log.Printf("[INFO] user found: %v", u.Email)

	return

}
