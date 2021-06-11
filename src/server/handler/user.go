package handler

import (
	"encoding/json"
	"gopkg.in/go-playground/validator.v9"
	"io/ioutil"
	"net/http"

	"way/pkg/logger"
	"way/src/core/user"
	"way/src/server/response"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Log("handler: creating user")

	var newUser user.AddUserRequestBody

	// Get new user object
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
	err = json.Unmarshal(requestBody, &newUser)
	if err != nil {
		logger.Log(err)
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
	successResponse, code, message, err := user.CreateUser(newUser)
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

func Verify(w http.ResponseWriter, r *http.Request) {
	logger.Log("handler: verifying user")

	var unverifiedUser user.Verify

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(err)
		return
	}

	logger.Log(string(requestBody))

	// decode body
	err = json.Unmarshal(requestBody, &unverifiedUser)
	if err != nil {
		logger.Log(err)
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

	successResponse := user.Verified{
		LoginId: 7,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger.Log("handler: user login")

	var userLogin user.LoginRequestBody

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
	err = json.Unmarshal(requestBody, &userLogin)
	if err != nil {
		logger.Log(err)
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
		logger.Log(err)
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
