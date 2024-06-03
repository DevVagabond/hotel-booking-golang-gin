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
}
