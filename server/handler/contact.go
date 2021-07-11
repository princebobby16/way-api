package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"log"
	"net/http"
	"way/core/user"
	"way/pkg/stringConversion"
	"way/server/response"
)

func AddContact(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: requesting contact")

	// get user_id from uri
	vars := mux.Vars(r)

	userId, err := stringConversion.ConvertStringToInt(vars["user_id"])
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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

	log.Println(string(requestBody))

	// decode body
	err = json.Unmarshal(requestBody, &newContact)
	if err != nil {
		log.Println(err)
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

	log.Println(userId)

	// check required fields
	err = validator.New().Struct(newContact)
	if err != nil {
		log.Println(err)
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
	successResponse, code, message, err := user.AddContact(newContact)
	if err != nil {
		log.Println(err)
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
	log.Println("handler: get contacts")

	vars := mux.Vars(r)

	userId, err := stringConversion.ConvertStringToInt(vars["user_id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "ERROR",
				Data: response.ErrorData{
					Code:    0,
					Message: "invalid",
				},
			})
		return
	}

	log.Println(userId)

	// Call service

	var successResponse []user.ContactRequestSent

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}

func RespondToContactRequest(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: responding to contact request")

	vars := mux.Vars(r)

	userId, err := stringConversion.ConvertStringToInt(vars["user_id"])
	if err != nil {
		log.Println(err)
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

	contactId, err := stringConversion.ConvertStringToInt(vars["contact_id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "ERROR",
				Data: response.ErrorData{
					Code:    400,
					Message: "invalid contact id",
				},
			})
		return
	}

	var responseToContactRequestBody user.ContactResponse

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(string(requestBody))

	// decode body
	err = json.Unmarshal(requestBody, &responseToContactRequestBody)
	if err != nil {
		log.Println(err)
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

	log.Println(contactId)
	log.Println(userId)

	// Call service

	w.WriteHeader(http.StatusCreated)
	return
}
