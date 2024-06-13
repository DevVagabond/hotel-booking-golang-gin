package response_handler

var ResponseMap = map[string]map[string]interface{}{
	"UNKNOWN": {
		"code":        500,
		"title":       "Unknown Error",
		"desc":        "An unknown error occurred",
		"messageCode": "UNKNOWN",
	},
	"VALIDATION_ERROR": {
		"code":        400,
		"title":       "Validation Error",
		"desc":        "Validation error occurred",
		"messageCode": "VALIDATION_ERROR",
	},
	"OK": {
		"code":        200,
		"title":       "OK",
		"desc":        "Request was successful",
		"messageCode": "OK",
	},
	"USER_ALREADY_EXISTS": {
		"code":        400,
		"title":       "User Already Exists",
		"desc":        "User already exists",
		"messageCode": "USER_ALREADY_EXISTS",
	},
	"INVALID_PASSWORD": {
		"code":        400,
		"title":       "Invalid Password",
		"desc":        "Invalid password.",
		"messageCode": "INVALID_PASSWORD",
	},
}
