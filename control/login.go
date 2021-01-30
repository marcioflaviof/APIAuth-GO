package control

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"restapi/database"
	"restapi/model"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	validate = validator.New()

	var user model.LoginUser

	bytes, err := ioutil.ReadAll(r.Body)

	if err = json.Unmarshal(bytes, &user); err != nil {
		log.Printf("[ERROR] Cannot unmarshal the JSON: %v %v", err, user)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = validate.Struct(user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	dbUser, err := database.SearchUser(user.Email)

	if err != nil {
		http.Error(w, "[ERROR] Email not found", 400)
		log.Printf("[ERROR] Email not found")
		return
	}

	//Check password
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))

	if err != nil {
		w.Write([]byte("[ERROR] Password doesn't match"))
		log.Println("[ERROR] Password doesn't match")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"password": user.Password,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))

	if err != nil {
		http.Error(w, "[ERROR] Token doesn't match", 400)
		log.Println("[ERROR] Token doesn't match")
		return
	}

	log.Println()
	w.Header().Set("auth-token", tokenString)

	w.Write([]byte("Logged In!"))

}
