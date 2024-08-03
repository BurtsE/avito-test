package service_errors

import "fmt"

type DatabaseError struct {
	err error
}

func (e DatabaseError) Error() string {
	return fmt.Sprintf("database error: %v", e.err)
}

type AuthError struct {
	err error
}

func (e AuthError) Error() string {
	return fmt.Sprintf("authentification error: %v", e.err)
}

type ValidationError struct {
	err error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("data validation error: %v", e.err)
}
