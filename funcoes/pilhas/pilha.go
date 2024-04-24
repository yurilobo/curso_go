package main

import "runtime/debug"

func F3() {
	debug.PrintStack()
}
func f2() {
	F3()
}
func f1() {
	f2()
}

func main() {
	f1()

}
