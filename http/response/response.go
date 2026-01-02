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

// SuccessNoData creates an successful response with empty data
func SuccessNoData(message string) *Response[struct{}] {
	return &Response[struct{}]{
		Status:  true,
		Message: message,
		Data:    struct{}{},
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

// ErrorNoData creates an error response with empty data
func ErrorNoData(message string) *Response[struct{}] {
	return &Response[struct{}]{
		Status:  false,
		Message: message,
		Data:    struct{}{},
	}
}
