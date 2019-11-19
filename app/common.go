package app

import (
	"errors"

	"session/log"
)

func DbError() error {
	log.Warningf("execute sql failed")
	return errors.New("invoke failed: Db error")
}

func RpcError(err error, errMsg string) error {
	log.Warningf("invoke rpc method error: %v", err)
	return errors.New("invoke failed: " + errMsg)
}

func NewError(str string) error {
	return errors.New(str)
}
