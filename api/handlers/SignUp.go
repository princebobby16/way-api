package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"way/pkg/logger"
	"way/pkg/response"
	"way/pkg/user"
)

func SignUp(w http.ResponseWriter, r *http.Request){
	logger.Log("handler: creating user")

	var newUser user.SignUp

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log(err)
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
				Message: "JSON request object not properly formed",
			},
		)
		return
	}

	// Call service

	successResponse := user.Created{
		UserId: 7,
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}