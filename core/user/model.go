package user

// User is a model of the basic information that makes up a users
type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	ContactId   int    `json:"contact_id"`
}

// CreateUserRequestBody is the json body of a request to create a user
type CreateUserRequestBody struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name" validate:"required"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	UserName        string `json:"username" validate:"required" `
	Password        string `json:"password" validate:"required" `
	ConfirmPassword string `json:"confirm_password" validate:"required" `
}

// AddUserResponseBody is the response object of successful user creation
type AddUserResponseBody struct {
	UserId string `json:"user_id"`
}

// LoginRequestBody represents a model of what login credentials look like.
type LoginRequestBody struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

// LoginResponseBody is the success response sent when a user is verified and login credentials are saved in the database
type LoginResponseBody struct {
	Token string `json:"token"`
}

// LoginData is the data model of a login in the database
type LoginData struct {
	LoginId   int    `json:"login_id"`
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
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
