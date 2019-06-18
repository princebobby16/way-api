package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"way/pkg/logger"
	"way/pkg/response"
	"way/pkg/stringConversion"
	"way/pkg/user"
)

func RespondToContactRequest(w http.ResponseWriter, r *http.Request) {
	logger.Log("handler: responding to contact request")

	vars := mux.Vars(r)

	userId, err := stringConversion.ConvertStringToInt(vars["user_id"])
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Message: "invalid user id",
			})
		return
	}

	contactId, err := stringConversion.ConvertStringToInt(vars["contact_id"])
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Message: "invalid contact id",
			})
		return
	}

	var responseToContactRequestBody user.ContactResponse

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(err)
		return
	}

	logger.Log(string(requestBody))

	// decode body
	err = json.Unmarshal(requestBody, &responseToContactRequestBody)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Message: "JSON request object not properly formed",
			},
		)
		return
	}

	logger.Log(contactId)
	logger.Log(userId)

	// Call service

	w.WriteHeader(http.StatusCreated)
	return
}
