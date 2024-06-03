package error_handler

import "fmt"

type ErrArg struct {
	Title       string `json:"title"`
	Code        int16  `json:"code"`
	Description string `json:"description"`
}

func (e *ErrArg) Error() string {
	return fmt.Sprintf("%d - %s -- %s", e.Code, e.Title, e.Description)
}
