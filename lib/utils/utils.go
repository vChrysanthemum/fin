package utils

import (
	"log"
	"runtime/debug"
)

func RecoverPanic() {
	if r := recover(); nil != r {
		log.Println(r)
		debug.PrintStack()
	}
}
