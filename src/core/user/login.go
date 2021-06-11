package user

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"way/pkg/db"
	"way/pkg/logger"
)

const (
	// Cost is the integer value used by bcrypt in password hashing
	Cost int = 15
)

// LoginRequestBody represents a model of what login credentials look like.
type LoginRequestBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResponseBody is the success response sent when a user is verified and login credentials are saved in the database
type LoginResponseBody struct {
	LoginId int    `json:"login_id"`
	UserId  int    `json:"user_id"`
	Token   string `json:"token"`
}

// LoginData is the data model of a login in the database
type LoginData struct {
	LoginId   int    `json:"login_id"`
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

// ToJson represents the LoginRequestBody struct as a json object.
// It returns the json and an error which is nil on success.
func (login LoginRequestBody) ToJson() ([]byte, error) {
	return json.Marshal(login)
}

// FromJson converts json data from an http.Request Body and decodes it as LoginRequestBody.
// It returns an error which is nil on success.
func (login *LoginRequestBody) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(login)
	if err != nil {
		return err
	}
	return nil
}

// HashPassword creates a salted hash of a string.
// It returns the hash of the password and an error with is nil if the password is successfully hashed.
func (login *LoginRequestBody) HashPassword() (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(login.Password), Cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Create token creates a JSON web token from a user's username and password
func (login *LoginRequestBody) CreateToken() (string, error) {

	var (
		// store database query results
		userData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Get items from db query
		getUserData = `SELECT login_id, user_id, username, password
		FROM way_api.login
		WHERE username = $1`
	)

	row := db.DBConnection.QueryRow(getUserData, login.Username)
	err := row.Scan(
		&userData.Username,
		&userData.Password,
	)
	if err != nil {
		return "", err
	}

	err = ComparePasswords(login.Password, userData.Password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": login.Username,
		"password": login.Password,
	})

	tokenString, err := token.SignedString([]byte("way"))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

func ComparePasswords(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}


func LogIn(login LoginRequestBody) (LoginResponseBody, int, string, error) {

	successResponse := LoginResponseBody{}

	var (
		userData LoginData

		getUserData = `SELECT login_id, user_id, username, password
		FROM way_api.login
		WHERE username = $1`
	)

	// check username
	row, err := db.DBConnection.Query(getUserData, login.Username)
	if err != nil {
		logger.Log(err)
		return successResponse, 400, "invalid username or password", err
	}

	err = row.Scan(
		&userData.LoginId,
		&userData.UserId,
		&userData.Username,
		&userData.Password,
	)
	if err != nil {
		logger.Log(err)
		return successResponse, 500, "internal server error", err
	}

	// compare passwords
	err = ComparePasswords(login.Password, userData.Password)
	if err != nil {
		logger.Log(err)
		return successResponse, 400, "invalid username or password", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login_id": userData.LoginId,
		"username": login.Username,
		"password": login.Password,
	})

	tokenString, err := token.SignedString([]byte("way"))

	if err != nil {
		fmt.Println(err)
		return successResponse, 500, "internal server error", err
	}

	successResponse.LoginId = userData.LoginId
	successResponse.UserId = userData.UserId
	successResponse.Token = tokenString

	return successResponse, 200, "logged in", nil

}

