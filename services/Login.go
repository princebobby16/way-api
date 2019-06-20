package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"way/pkg/db"
	"way/pkg/logger"
	"way/pkg/user"
)

func LogIn(login user.LoginRequestBody) (user.LoginResponseBody, int, string, error) {

	successResponse := user.LoginResponseBody{}

	var (
		userData user.LoginData

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
	err = user.ComparePasswords(login.Password, userData.Password)
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
