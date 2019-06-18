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

func AddCOntact(w http.ResponseWriter, r *http.Request) {
	logger.Log("handler: requesting contact")

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

	logger.Log(contactId)
	logger.Log(userId)

	// Call service

	successResponse := user.ContactRequestSent{
		ContactId: 7,
		Status:    "PENDING",
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}
