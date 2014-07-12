package view

import (
	"fmt"
)

func Print(a ...interface{}) String {
	return String(fmt.Sprint(a...))
}

func Printf(format string, a ...interface{}) String {
	return String(fmt.Sprintf(format, a...))
}

func Println(a ...interface{}) String {
	return String(fmt.Sprintln(a...))
}
