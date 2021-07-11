package user

import (
	"log"
	"way/pkg/db"
)

const (
	Pending = "PENDING"
)

type AddContactRequestBody struct {
	UserId    int `json:"user_id"`
	ContactId int `json:"contact_id" validate:"required"`
}

type AddContactResponseBody struct {
	ContactId int    `json:"contact_id"`
	Status    string `json:"status"`
}

func AddContact(newContact AddContactRequestBody) (AddContactResponseBody, int, string, error) {

	successResponse := AddContactResponseBody{}

	var (
		insertUserQuery = `INSERT INTO way_api.relationship (user_1, user_2, status, last_actor)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id
`
		lastInsertedId int
	)

	// save new user
	err := db.DBConnection.QueryRow(insertUserQuery, newContact.UserId, newContact.ContactId, Pending, newContact.UserId).Scan(&lastInsertedId)
	if err != nil {
		log.Println(err)
		return successResponse, 400, "invalid id", err
	}

	// send request to contact

	// todo: notify contact

	successResponse.ContactId = newContact.ContactId
	successResponse.Status = Pending

	return successResponse, 200, "contact added", nil

}
