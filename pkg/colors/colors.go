package colors

import "fmt"

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Redf(format string, a ...interface{}) string {
	return Red + fmt.Sprintf(format, a...) + Reset
}

func Greenf(format string, a ...interface{}) string {
	return Green + fmt.Sprintf(format, a...) + Reset
}
