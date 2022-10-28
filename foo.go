package test

import "log"

func LibFunc(foo string) string {
	log.Print("asd")
	return foo
}

func FooBar(bar int) int {
	return bar
}
