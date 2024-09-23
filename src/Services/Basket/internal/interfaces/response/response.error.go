package response

type ErrorResponse struct {
	Response
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{
		Response: Response{
			Success: false,
			Error: &APIError{
				Code:    code,
				Message: message,
			},
		},
	}
}
