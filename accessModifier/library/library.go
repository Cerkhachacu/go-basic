package library

import "log"

// private struct
type student struct {
	Name   string
	grande int
}

// Items public struct
type Items struct {
	Number string
	Value  int
}

// SayHello public func
func SayHello() {
	log.Println("say hello ")
}

// introduce private func
func introduce(name string) {
	log.Println("hello my name is :", name)
}
