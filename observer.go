package main

import "sync"

type Observer interface {
	update(string, string, *sync.WaitGroup)
	getID() string
}
