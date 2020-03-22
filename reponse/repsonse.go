package reponse

type Response struct {
	Data    interface{} `json:"data,omitempy"`
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempy"`
}

func ReturnData(data interface{}) Response {
	return Response{
		Data:    data,
		Success: true,
		Error:   "",
	}
}

func ReturnError(err error) Response {
	return Response{
		Data:    nil,
		Success: false,
		Error:   err.Error(),
	}
}
