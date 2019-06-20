package services

import (
	"crypto/rand"
	"io"
	"way/pkg/db"
	"way/pkg/logger"
	"way/pkg/user"
)

func SignUp(newUser user.AddUserRequestBody) (user.AddUserResponseBody, int, string, error) {

	successResponse := user.AddUserResponseBody{}

	var (
		insertUserQuery = `INSERT INTO way_api.user (first_name, last_name, phone_number)
		VALUES ($1, $2, $3)
		RETURNING user_id
`
		lastInsertedId int
	)


	// Todo: abstract pin generation and saving to new endpoint /users/newpin
	// create confirmation pin
	temporaryPin := encodeToString(5)

	// set expiration date

logger.Log(temporaryPin)

	// check username
	err := db.DBConnection.QueryRow(insertUserQuery, newUser.FirstName, newUser.LastName, newUser.PhoneNumber).Scan(&lastInsertedId)
	if err != nil {
		logger.Log(err)
		return successResponse, 400, "invalid phone number", err
	}

	// todo: send confirmation pin

	successResponse.UserId = lastInsertedId

	return successResponse, 200, "logged in", nil

}

func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

