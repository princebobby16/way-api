package user

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"way/pkg/db"
)

func CreateUser(newUser CreateUserRequestBody) (AddUserResponseBody, int, string, error) {

	response := AddUserResponseBody{}

	// compare password
	if newUser.Password != newUser.ConfirmPassword {
		err := errors.New("passwords do not match")
		log.Println("passwords do not match")
		return response, 400, "passwords do not match", err
	}

	// check if phone number is already in use
	var phoneNumberExists bool
	checkPhoneNumberQuery := `SELECT EXISTS(SELECT 1 FROM way_api.user where phone_number=$1)`
	err := db.DBConnection.QueryRow(checkPhoneNumberQuery, newUser.PhoneNumber).Scan(&phoneNumberExists)
	if err != nil {
		log.Println(err)
		return response, 500, "internal server error", err
	}

	//insert user into user table
	insertUserQuery := `INSERT INTO way_api.user (first_name, last_name, phone_number)
		VALUES ($1, $2, $3)
		RETURNING user_id
		`
	var userId string

	err = db.DBConnection.QueryRow(insertUserQuery, newUser.FirstName, newUser.LastName, newUser.PhoneNumber).Scan(&userId)
	if err != nil {

		log.Println(err)
		return response, 500, "internal server error", err
	}

	//insert user into user table
	insertLoginQuery := `INSERT INTO way_api.login (user_id, username, password)
		VALUES ($1, $2, $3)
`
	// hash password
	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		log.Println(err)
		log.Println("Password hash error")
		return response, 500, "internal server error", err
	}

	_, err = db.DBConnection.Exec(insertLoginQuery, userId, newUser.UserName, hashedPassword)
	if err != nil {
		log.Println(err)
		return response, 500, "internal server error", err
	}

	// send temporary pin

	// todo: send confirmation pin
	// Todo: abstract pin generation and saving to new endpoint /users/newpin {phone_number} or the function

	response.UserId = userId

	return response, 200, "user created", nil

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

	successResponse.Token = tokenString

	return successResponse, 200, "logged in", nil

}
