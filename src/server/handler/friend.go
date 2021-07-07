package handler

import (
	"encoding/json"
	"net/http"
	"way/pkg/logger"
	"way/src/core/user"
	"way/src/server/response"
)

func GetFriends(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	logger.Echo(userId)

	friends, code, msg, err := user.GetFriends(userId)
	if err != nil {
		logger.Log(err.Error())
		w.WriteHeader(code)
		_ = json.NewEncoder(w).Encode(
			response.Error{
				Status: "failed",
				Data: response.ErrorData{
					Code:    code,
					Message: msg,
				},
			},
		)
		return
	}

	successResponse := &user.GetFriendsResponse{
		UserId:  userId,
		Friends: friends,
		Status:  "success",
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}
