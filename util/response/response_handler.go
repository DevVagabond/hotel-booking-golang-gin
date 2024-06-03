package response_handler

import error_handler "hotel-booking-golang-gin/util/error"

type Response struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	MessageCode string      `json:"message_code"`
	Data        interface{} `json:"data"`
}

func NewResponse(success bool, code string, message string, data interface{}) *Response {

	code_obj := ResponseMap[code]
	if code_obj == nil {
		code_obj = ResponseMap["UNKNOWN"]
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	return &Response{
		Success:     success,
		Message:     message,
		MessageCode: code_obj["messageCode"].(string),
		Data:        data,
	}
}

func OK(data interface{}) *Response {
	if data == nil {
		data = make(map[string]interface{})
	}
	return NewResponse(true, "OK", "SUCCESS", data)
}

func Error(code string, err *error_handler.ErrArg) *Response {
	return NewResponse(false, code, err.Error(), nil)
}
