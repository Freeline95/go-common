package error

import "fmt"

func Annotate (err error, message string) error {
	return fmt.Errorf(message + ": %w", err)
}