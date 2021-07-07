package user

import (
	"github.com/lib/pq"
	"way/pkg/db"
	"way/pkg/logger"
)

type Friend struct {
	UserId      int      `json:"user_id"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	PhoneNumber string   `json:"phone_number"`
	Location    []string `json:"location"`
}

type GetFriendsResponse struct {
	UserId  string   `json:"user_id"`
	Friends []Friend `json:"friends"`
	Status  string   `json:"status"`
}

func GetFriends(userId string) ([]Friend, int, string, error) {
	var (
		friends []Friend

		getFriendsQuery = `select 
       u.user_id, u.first_name, u.last_name, u.phone_number, u.last_location 
from way_api.user u, way_api.relationship rel 
where rel.user_1 = $1 and rel.status = 'active' and u.user_id = rel.user_2;
`
	)
	logger.Log(getFriendsQuery)

	rows, err := db.DBConnection.Query(getFriendsQuery, &userId)
	if err != nil {
		return nil, 400, "", err
	}

	for rows.Next() {
		var friend Friend
		err := rows.Scan(
			&friend.UserId,
			&friend.FirstName,
			&friend.LastName,
			&friend.PhoneNumber,
			pq.Array(&friend.Location),
		)
		if err != nil {
			return nil, 500, err.Error(), err
		}

		friends = append(friends, friend)
	}

	return friends, 200, "", nil
}
