package user

const (
	Pending = "PENDING"
)

type AddContactRequestBody struct {
	UserId int `json:"user_id"`
	ContactId int `json:"contact_id" validate:"required"`
}

type AddContactResponseBody struct {
	ContactId int `json:"contact_id"`
	Status string `json:"status"`
}


