package response

type ErrorData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Error struct {
	Status string    `json:"status"`
	Data   ErrorData `json:"data"`
}

func (response *Error) ErrorResponse(status string, code int, message string) Error {

	response.Data.Message = message
	response.Data.Code = code
	response.Status = status

	return *response
}
