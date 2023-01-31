package common

import "errors"

func Error(errorInfo string) error {
	return errors.New(errorInfo)
}
