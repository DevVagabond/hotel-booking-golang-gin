package response_handler

import error_handler "hotel-booking-golang-gin/util/error"

type Response struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	MessageCode int         `json:"message_code"`
	Data        interface{} `json:"data"`
}

func NewResponse(success bool, code int, message string, data interface{}) *Response {
	if code == 0 {
		code = 200
	}

	return &Response{
		Success:     success,
		Message:     message,
		MessageCode: code,
		Data:        data,
	}
}

func OK(data interface{}) *Response {
	if data == nil {
		data = make(map[string]interface{})
	}
	return NewResponse(true, 200, "OK", data)
}

func Error(code int, err *error_handler.ErrArg) *Response {
	return NewResponse(false, code, err.Error(), nil)
}
