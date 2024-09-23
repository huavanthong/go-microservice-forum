package response

type SuccessResponse struct {
	Response
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Response: Response{
			Success: true,
			Data:    data,
		},
	}
}
