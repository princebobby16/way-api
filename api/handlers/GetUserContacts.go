package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"way/pkg/logger"
	"way/pkg/response"
	"way/pkg/stringConversion"
	"way/pkg/user"
)

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
