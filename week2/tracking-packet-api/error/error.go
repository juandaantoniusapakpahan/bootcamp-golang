package errorhttp

import "fmt"

type Http struct {
	StatusCode  int    `json:"statusCode"`
	Description string `json:"description,omitempty"`
}

func (e Http) Error() string {
	return fmt.Sprintf("description: %s", e.Description)
}

func NewHttpError(description string, statusCode int) Http {
	return Http{
		Description: description,
		StatusCode:  statusCode,
	}
}
