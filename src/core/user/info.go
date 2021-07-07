package user

import (
	"encoding/json"
	"net/http"
	"way/pkg/logger"
)

// Info is a model of the basic information that makes up a users
type Info struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	ContactId   int    `json:"contact_id"`
}

type AddUserRequestBody struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}

// AddUserResponseBody is the
type AddUserResponseBody struct {
	UserId int `json:"user_id"`
}

// ToJson converts Info to json.
// Returns a []byte and and error which is not nil if it fails to convert
func (i Info) ToJson() ([]byte, error) {
	return json.Marshal(i)
}

// FromJson converts a json object to Info
// Returns an error if there is a failure in conversion
func (i *AddUserRequestBody) FromJson(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		logger.Echo(err.Error())
		return err
	}
	return nil
}
