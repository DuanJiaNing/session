package app

import (
	"errors"
	"session/log"
)

func CheckDbExecuteResult(affected int64, err error, exceptAffected int64) error {
	if err != nil {
		return WithInternalError(err)
	}
	if affected != exceptAffected {
		return DbExecuteAffectedRowsIncorrect()
	}

	return nil
}

func DbExecuteAffectedRowsIncorrect() error {
	log.Warningf("execute sql affected rows incorrect")
	return internalError("db error")
}

func WithInternalError(err error) error {
	return withInternalError(err, "Internal error")
}

// withInternalError report internal error with error stack
func withInternalError(err error, msg string) error {
	log.Warningf("got internal error: %v, msg: %s", err, msg)
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
