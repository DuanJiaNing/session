package log

import (
	"fmt"
	"log"
)

func Info(args ...interface{}) {
	log.Println("info: " + fmt.Sprint(args))
}

func Infof(format string, args ...interface{}) {
	log.Println("info: " + fmt.Sprintf(format, args))
}

func Warning(args ...interface{}) {
	log.Println("warn: " + fmt.Sprint(args))
}

func Warningf(format string, args ...interface{}) {
	log.Println("warn: " + fmt.Sprintf(format, args))
}

func Error(args ...interface{}) {
	log.Println("error: " + fmt.Sprint(args))
}

func Errorf(format string, args ...interface{}) {
	log.Println("error: " + fmt.Sprintf(format, args))
}
