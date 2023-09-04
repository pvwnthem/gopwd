package colors

import (
	"fmt"

	"github.com/torbenconto/zeus"
)

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Redf(format string, a ...interface{}) string {
	return zeus.Concat(Red, fmt.Sprintf(format, a...), Reset)
}

func Greenf(format string, a ...interface{}) string {
	return zeus.Concat(Green, fmt.Sprintf(format, a...), Reset)
}
