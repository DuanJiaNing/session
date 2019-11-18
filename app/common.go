package app

import (
	"errors"

	"session/log"
)

func RpcError(err error, errMsg string) error {
	log.Warningf("invoke rpc method error: %v", err)
	return errors.New("invoke failed: " + errMsg)
}
