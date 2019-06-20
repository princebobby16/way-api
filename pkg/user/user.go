package user

type Created struct {
	UserId int `json:"user_id"`
}

type Verify struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Pin      string `json:"pin"`
}

type Verified struct {
	LoginId int `json:"login_id"`
}


type ContactRequestSent struct {
	ContactId int    `json:"contact_id"`
	Status    string `json:"status"`
}

type ContactResponse struct {
	Action string `json:"action"`
}
