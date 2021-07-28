package user

import (
	random "crypto/rand"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	// Cost is the integer value used by bcrypt in password hashing
	Cost int = 15
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func encodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(random.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

// ToJson converts User to json.
// Returns a []byte and and error which is not nil if it fails to convert
func (i User) ToJson() ([]byte, error) {
	return json.Marshal(i)
}

// FromJson converts a json object to User
// Returns an error if there is a failure in conversion
func (i *CreateUserRequestBody) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
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
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), Cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// Create token creates a JSON web token from a user's username and password

func ComparePasswords(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SendSMS(message string, numberTo string) (int, error) {

	urlStr := "https://sms.arkesel.com/sms/api?action=send-sms&api_key=OnlvZm9ycmVhbC5jb20=&%20to=" + numberTo + "&from=Way&sms=" + message

	client := &http.Client{}
	req, _ := http.NewRequest("GET", urlStr, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return 500, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Println(resp.StatusCode)
		log.Println(resp.Body)
		log.Println(resp.Request.RequestURI)
		return 201, err

	} else {
		log.Println(resp.Status)
		return resp.StatusCode, nil
	}
}

func randomNumber(min int, max int) string {
	return strconv.Itoa(rand.Intn(max-min) + min)
}
