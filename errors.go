package gopp

import "errors"

func Error(message string) error {
	return errors.New(message)
}
