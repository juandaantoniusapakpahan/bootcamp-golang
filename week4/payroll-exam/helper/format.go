package helper

func NewResponse(code int, status string, data map[string]interface{}) *JsonResponse {
	return &JsonResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
