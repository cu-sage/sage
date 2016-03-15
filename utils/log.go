package utils

import (
	"errors"
	"fmt"
	"log"
)

func LogAndReturnError(s string, a ...interface{}) error {
	msg := fmt.Sprintf(s, a)
	log.Println(msg)

	return errors.New(msg)
}
