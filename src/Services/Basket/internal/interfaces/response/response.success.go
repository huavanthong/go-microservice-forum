package response

type SuccessResponse struct {
	Response `json:"response"`
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Response: Response{
			Success: true,
			Data:    data,
		},
	}
}
