package impl

import (
	"fmt"
	"log"
	"strings"

	"github.com/danilluk1/test-task-6/internal/services/logger"
)

type clogger struct{}

func NewLogger() logger.Logger {
	return &clogger{}
}

func (c clogger) Info(msg string, args ...string) {
	fmt.Printf("%s %s\n", msg, strings.Join(args, " "))
}

func (c clogger) Error(args ...any) {
	log.Println(args)
}
