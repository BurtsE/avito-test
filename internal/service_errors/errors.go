package service_errors

import "fmt"

type ServerError struct {
	Err error
}

func (e ServerError) Error() string {
	return fmt.Sprintf("server error: %v", e.Err)
}

type AuthError struct {
	Err error
}

func (e AuthError) Error() string {
	return fmt.Sprintf("authentification error: %v", e.Err)
}

type ValidationError struct {
	Err error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("data validation error: %v", e.Err)
}
