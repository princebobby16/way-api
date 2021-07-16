package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"way/pkg/db"
)

func CreateUser(newUser CreateUserRequestBody) (AddUserResponseBody, int, string, error) {

	successResponse := AddUserResponseBody{}

	var (
		insertUserQuery = `INSERT INTO way_api.user (first_name, last_name, phone_number)
		VALUES ($1, $2, $3)
		RETURNING user_id
`
		lastInsertedId string
	)

	// save new user
	err := db.DBConnection.QueryRow(insertUserQuery, newUser.FirstName, newUser.LastName, newUser.PhoneNumber).Scan(&lastInsertedId)
	if err != nil {

		log.Println(err)
		return successResponse, 400, "phone number already exists", err
	}

	// send temporary pin

	// todo: send confirmation pin
	// Todo: abstract pin generation and saving to new endpoint /users/newpin {phone_number} or the function

	successResponse.UserId = lastInsertedId

	return successResponse, 200, "user created", nil

}

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

	row := db.DBConnection.QueryRow(getUserData, login.PhoneNumber)
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
		"username": login.PhoneNumber,
		"password": login.Password,
	})

	tokenString, err := token.SignedString([]byte("way"))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
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
	row, err := db.DBConnection.Query(getUserData, login.PhoneNumber)
	if err != nil {
		log.Println(err)
		return successResponse, 400, "invalid username or password", err
	}

	err = row.Scan(
		&userData.LoginId,
		&userData.UserId,
		&userData.Username,
		&userData.Password,
	)
	if err != nil {
		log.Println(err)
		return successResponse, 500, "internal server error", err
	}

	// compare passwords
	err = ComparePasswords(login.Password, userData.Password)
	if err != nil {
		log.Println(err)
		return successResponse, 400, "invalid username or password", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login_id": userData.LoginId,
		"username": login.PhoneNumber,
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
