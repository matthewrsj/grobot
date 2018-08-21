package main

import (
	"sync"
)

var scores sync.Map

func updateScore(nick string, change int) {
	var current int
	if val, ok := scores.Load(nick); !ok {
		current = 0
	} else {
		current = val.(int)
	}

	scores.Store(nick, current+change)
}

func getScore(nick string) int {
	val, ok := scores.Load(nick)
	if !ok {
		return 0
	}
	return val.(int)
}
