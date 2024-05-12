package response

type ApiErrorResponse struct {
	Error string
}

func NewApiErrorResponse(errorMsg string) *ApiErrorResponse {
	return &ApiErrorResponse{
		Error: errorMsg,
	}
}
