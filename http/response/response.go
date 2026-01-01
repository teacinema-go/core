package response

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data map[string]any

func NewSuccessResponse(message string, data Data) *Response {
	if data == nil {
		data = Data{}
	}

	return &Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string, data Data) *Response {
	if data == nil {
		data = Data{}
	}

	return &Response{
		Status:  false,
		Message: message,
		Data:    data,
	}
}
