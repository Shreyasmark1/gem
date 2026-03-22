package utils

import (
	"fmt"
	"runtime/debug"
)

func DebugError(format string, args ...any) error {

	debug.PrintStack()

	msg := fmt.Sprintf(format, args...)

	return fmt.Errorf("%s", msg)
}
