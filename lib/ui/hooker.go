package ui

type Hooker struct {
	Arg interface{}
	Do  func(arg interface{})
}
