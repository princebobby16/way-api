package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"
	"way/pkg/logger"
	"way/pkg/response"
	"way/pkg/stringConversion"
	"way/pkg/user"
	"way/services"
)

func AddContact(w http.ResponseWriter, r *http.Request) {
	logger.Log("handler: requesting contact")

	// get user_id from uri
	vars := mux.Vars(r)

	userId, err := stringConversion.ConvertStringToInt(vars["user_id"])
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code:    0,
					Message: "invalid user id",
				},
			})
		return
	}

	var newContact user.AddContactRequestBody

	newContact.UserId = userId

	// Get new contact object
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code:    0,
					Message: "could not read request body",
				},
			},
		)
		return
	}

	logger.Log(string(requestBody))

	// decode body
	err = json.Unmarshal(requestBody, &newContact)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code:    0,
					Message: "JSON request object not properly formed",
				},
			},
		)
		return
	}

	logger.Log(userId)

	// check required fields
	err = validator.New().Struct(newContact)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code:    0,
					Message: "bad data",
				},
			},
		)
		return
	}

	// Call service
	successResponse, code, message, err := services.AddContact(newContact)
	if err != nil {
		logger.Log(err)
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code:    code,
					Message: message,
				},
			},
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	logger.Log("handler: get contacts")

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

	logger.Log(userId)

	// Call service

	var successResponse []user.ContactRequestSent

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}

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
