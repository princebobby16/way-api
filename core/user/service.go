package user

import (
	"errors"
	"log"
	"math/rand"
	"time"
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

	response.UserId = userId

	return response, 200, "user created", nil

}

func SendPIN(newUser RequestPINBody) (int, string, error) {

	// todo: check if phone number is in the system and unverified

	// generate otp
	rand.Seed(time.Now().UnixNano())
	pin := randomNumber(1000, 9999)

	// set expiry date
	t := time.Now()
	expiration := t.Add(time.Minute * 10)

	// save to db
	//insert pin into user table
	insertPINQuery := `UPDATE way_api.user SET temporary_pin = $1, temporary_pin_expiry = $2
		WHERE phone_number = $3
`

	result, err := db.DBConnection.Exec(insertPINQuery, pin, expiration, newUser.PhoneNumber)
	if err != nil {
		log.Println(err)
		return 500, "internal server error", err
	}
	rows, _ := result.RowsAffected()
	if rows != 1 {
		return 400, "invalid phone number", errors.New("invalid phone number")
	}

	message := pin
	code, err := SendSMS(message, newUser.PhoneNumber)
	if err != nil {
		log.Println(err)
		return code, "fail", err
	}
	return 200, "success", nil
}

func VerifyUser(verifyBody VerificationRequestBody) (int, string, error) {
	// an object of the database data we will work with
	type dbData struct {
		PhoneNumber string
		PIN         string
		Expiry      string
	}

	var d dbData

	// check phone and pin
	checkPhoneQuery := `SELECT phone_number,  temporary_pin, temporary_pin_expiry, verified FROM way_api.user WHERE phone_number=$1`

	var isVerified bool

	err := db.DBConnection.QueryRow(checkPhoneQuery, verifyBody.PhoneNumber).Scan(&d.PhoneNumber, &d.PIN, &d.Expiry, &isVerified)
	if err != nil {
		log.Println(err)
		return 400, "invalid phone number", err
	}

	if isVerified {
		return 400, "user already verified", errors.New("user already verified")
	}

	// compare pins
	if d.PIN != verifyBody.Pin {
		return 400, "invalid pin", errors.New("invalid pin")
	}

	// check pin expiry
	expiry, err := time.Parse(time.RFC3339, d.Expiry)
	if err != nil {
		log.Println(err)
		return 500, "internal server error", err
	}

	if time.Now().After(expiry) {
		return 400, "your pin has expired", errors.New("your pin has expired")
	}

	// set user as verified
	verifyQuery := `UPDATE way_api.user SET verified = true
		WHERE phone_number = $1
`

	_, err = db.DBConnection.Exec(verifyQuery, verifyBody.PhoneNumber)
	if err != nil {
		log.Println(err)
		return 500, "internal server error", err
	}

	return 201, "user verified successfully", nil

}

func LogIn(login LoginRequestBody) (LoginResponseBody, int, string, error) {
	var response LoginResponseBody

	// an object of the database data we will work with
	type dbData struct {
		UserId   string
		LoginId  string
		Username string
		Password string
	}

	var d dbData

	// get details from database by username
	checkUsernameQuery := `SELECT user_id, login_id, username,  password FROM way_api.login WHERE username=$1`

	err := db.DBConnection.QueryRow(checkUsernameQuery, login.Username).Scan(&d.UserId, &d.LoginId, &d.Username, &d.Password)
	if err != nil {
		log.Println(err)
		return response, 400, "username or password invalid", err
	}

	// check if user is verified
	checkVerifyQuery := `SELECT verified FROM way_api.user WHERE user_id=$1`

	var verified bool

	err = db.DBConnection.QueryRow(checkVerifyQuery, d.UserId).Scan(&verified)
	if err != nil {
		log.Println(err)
		return response, 400, "user not verified", err
	}

	// compare passwords
	err = ComparePasswords(login.Password, d.Password)
	if err != nil {
		log.Println(err)
		return response, 400, "username or password invalid", err
	}

	// at this point username and password are valid and user is a verified user

	// create token
	tokenBody, err := CreateToken(d.UserId, d.LoginId)
	if err != nil {
		log.Println(err)
		return response, 500, "internal server error", err
	}

	response.Token = tokenBody

	return response, 200, "login successful", nil
}
