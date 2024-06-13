package error_handler

import "fmt"

type ErrArg struct {
	Title       string `json:"title"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (e *ErrArg) Error() string {
	return fmt.Sprintf("%s - %s -- %s", e.Code, e.Title, e.Description)
}
