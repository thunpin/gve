package network

import "log"

type Result interface {
	Log()
	Message() string
}

type Success struct {
	message string
}

type Warning struct {
	message string
}

func (result Success) Message() string {
	return "[SUCCESS] " + result.message
}
func (result Success) Log() {
	log.Println(result.Message())
}

func (result Warning) Message() string {
	return "[WARNING] " + result.message
}
func (result Warning) Log() {
	log.Println(result.Message())
}
