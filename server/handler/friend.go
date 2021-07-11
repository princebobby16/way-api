package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"way/core/user"
	"way/server/response"
)

func GetFriends(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")
	log.Println(userId)

	friends, code, msg, err := user.GetFriends(userId)
	if err != nil {
		log.Println(err.Error())
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

	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(successResponse)
	return
}
