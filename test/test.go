package test

import (
	"fmt"
	"time"
)

func MockTopic() string {
	return fmt.Sprintf("test_%s", time.Now())
}
