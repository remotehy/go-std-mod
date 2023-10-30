package log

import "fmt"

func StdLog(tpl string, args ...any) {
	fmt.Printf(tpl, args...)
	fmt.Println()
}
