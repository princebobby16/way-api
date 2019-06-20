package user

import "encoding/json"

// Info is a model of the basic information that makes up a users
type Info struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	ContactId   int    `json:"contact_id"`
}

// ToJson converts Info to json.
// Returns a []byte and and error which is not nil if it fails to convert
func (i Info) ToJson() ([]byte, error) {
	return json.Marshal(i)
}


