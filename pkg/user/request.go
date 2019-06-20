package user

import (
	"encoding/json"
	"errors"
	"net/http"
	"way/pkg/logger"
)

type CreateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

// FromJson converts a json object to Info
// Returns an error if there is a failure in conversion
func (i *CreateRequest) FromJson(r *http.Request) (error) {
	err := json.NewDecoder(r.Body).Decode(i)
	if err != nil {
		logger.Echo(err.Error())
		return err
	}
	return nil
}

/*
	ValidateRequestData check for empty fields in the request body
	Returns an error if one or more entries are empty
*/
func (i CreateRequest) ValidateRequestData() error {
	if len(i.LastName) == 0 || len(i.Email) == 0 || len(i.DateOfBirth) == 0 || len(i.Gender) == 0 || len(i.Category) == 0 {
		err := errors.New("empty request body")
		return err
	}
	return nil
}

