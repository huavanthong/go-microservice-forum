package response

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewResponse(data interface{}, err *APIError) *Response {
	if err != nil {
		return &Response{
			Success: false,
			Error:   err,
		}
	}

	return &Response{
		Success: true,
		Data:    data,
	}
}
