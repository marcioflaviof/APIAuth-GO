package control

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"restapi/database"
	"restapi/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

func Auth(w http.ResponseWriter, r *http.Request) {
	validate = validator.New()

	var user model.User

	bytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("[ERROR] Can't read response body", err)
	}

	if err = json.Unmarshal(bytes, &user); err != nil {
		fmt.Printf("[ERROR] Cannot unmarshal the JSON: %v %v", err, user)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	_, err = database.SearchUser(user.Email)

	if err == nil {
		http.Error(w, "[ERROR] User already exists", 400)
		fmt.Printf("[ERROR] User already exists")
		return
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	user.Password = string(hashPass)
	user.Date = time.Now()

	if err != nil {
		fmt.Printf("[ERROR] in hashing %s", err)
	}

	err = database.AddUser(user)

	if err != nil {
		fmt.Println("[ERROR] Couldn't add user into database")
	}

	marshUser, _ := json.Marshal(user)

	w.Write(marshUser)
}

func VerifyToken(tokenString string) bool {

	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err == nil && result.Valid {
		fmt.Println("Valid")
		return true
	}

	fmt.Println("Invalid")
	return false
}
