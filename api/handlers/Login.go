package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"way/pkg/logger"
	"way/pkg/response"
	"way/pkg/user"
	"way/services"

	"gopkg.in/go-playground/validator.v9"
)

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
	successResponse, code, message, err := services.LogIn(userLogin)
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
