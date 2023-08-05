package common

func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": "request OK",
		"data":    data,
	}
}

func FailureResponse(errorMessage string) map[string]interface{} {
	return map[string]interface{}{
		"success": true,
		"message": errorMessage,
		"data":    nil,
	}
}
