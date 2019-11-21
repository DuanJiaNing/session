package app

import (
	"errors"
	"session/log"
)

func DbExecuteEffectRowsIncorrect() error {
	log.Warningf("execute sql effect rows incorrect")
	return internalError("db error")
}

func WithInternalError(err error) error {
	return withInternalError(err, "Internal error")
}

// withInternalError report internal error with error stack
func withInternalError(err error, msg string) error {
	log.Warningf("got internal error: %v, msg: %v", err, msg)
	return errors.New(msg)
}

// internalError report error these defined by developer
func internalError(msg string) error {
	log.Warningf("got internal error, error msg: %v", msg)
	return errors.New(msg)
}

// Error report checked error
func Error(str string) error {
	log.Warningf("got checked error: %v", str)
	return errors.New(str)
}
