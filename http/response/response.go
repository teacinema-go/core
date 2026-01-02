package response

// Response standard API response structure with generic data type
type Response[T any] struct {
	Status  bool   `json:"status" example:"true"`
	Message string `json:"message" example:"ok"`
	Data    T      `json:"data"`
}

// Success creates a successful response with generic data
func Success[T any](message string, data T) *Response[T] {
	return &Response[T]{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

// Error creates an error response with generic data
func Error[T any](message string, data T) *Response[T] {
	return &Response[T]{
		Status:  false,
		Message: message,
		Data:    data,
	}
}
