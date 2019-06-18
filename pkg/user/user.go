package user

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	ContactId   int    `json:"contact_id"`
}

type SignUp struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
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

type LoggedIn struct {
	LoginId int `json:"login_id"`
	UserId  int `json:"user_id"`
}

type ContactRequestSent struct {
	ContactId int    `json:"contact_id"`
	Status    string `json:"status"`
}

type ContactResponse struct {
	Action string `json:"action"`
}
