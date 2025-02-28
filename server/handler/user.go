package handler

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"log"
	"net/http"
	"way/core/user"
	"way/server/response"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: creating user")

	var newUser user.CreateUserRequestBody

	// Get new user object
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

	// decode body
	err = json.Unmarshal(requestBody, &newUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code: 0,

					Message: "JSON request object not properly formed",
				},
			},
		)
		return
	}

	// check required fields
	err = validator.New().Struct(newUser)
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
	successResponse, code, message, err := user.CreateUser(newUser)
	if err != nil {
		log.Println(err)
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: string(code),
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

func RequestPIN(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: requesting PIN")

	var pinRequest user.RequestPINBody

	// Get new user object
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
	err = json.Unmarshal(requestBody, &pinRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code: 0,

					Message: "JSON request object not properly formed",
				},
			},
		)
		return
	}

	// check required fields
	err = validator.New().Struct(pinRequest)
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
	code, message, err := user.SendPIN(pinRequest)
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
	return
}

func Verify(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: verifying user")

	var verificationBody user.VerificationRequestBody

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(requestBody)
	// decode body
	err = json.Unmarshal(requestBody, &verificationBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code: 0,
					Message: "JSON request object not properly formed",
				},
			},
		)
		return
	}

	// Call service
	code, message, err := user.VerifyUser(verificationBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: string(code),
				Data: response.ErrorData{
					Code:    code,
					Message: message,
				},
			},
		)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("handler: user login")

	var userLogin user.LoginRequestBody

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

	// decode body
	err = json.Unmarshal(requestBody, &userLogin)
	if err != nil {
	log.Println(err)
	w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code: 0,

					Message: "JSON request object not properly formed",
				},
			},
		)
		return
	}

	// check required fields
	err = validator.New().Struct(userLogin)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "",
				Data: response.ErrorData{
					Code: 0,

					Message: "bad data",
				},
			},
		)
		return
	}

	// Call service
	successResponse, code, message, err := user.LogIn(userLogin)
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
