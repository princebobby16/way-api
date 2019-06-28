package services

import (
	"way/pkg/db"
	"way/pkg/logger"
	"way/pkg/user"
)

func AddContact(newContact user.AddContactRequestBody) (user.AddContactResponseBody, int, string, error) {

	successResponse := user.AddContactResponseBody{}

	var (
		insertUserQuery = `INSERT INTO way_api.relationship (user_1, user_2, status, last_actor)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id
`
		lastInsertedId int
	)


	// save new user
	err := db.DBConnection.QueryRow(insertUserQuery, newContact.UserId, newContact.ContactId, user.Pending,newContact.UserId).Scan(&lastInsertedId)
	if err != nil {
		logger.Log(err)
		return successResponse, 400, "invalid id", err
	}

	// send request to contact

	// todo: notify contact

	successResponse.ContactId = newContact.ContactId
	successResponse.Status=user.Pending

	return successResponse, 200, "contact added", nil

}
