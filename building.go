package main

import (
	"errors"
	"fmt"
	"sync"
)

type Building struct {
	observerList []Observer
	name         string
	status       string
}

func newBuilding(name string) *Building {
	return &Building{
		name: name,
	}
}
func (b *Building) updateAvailability(status string) error {
	if status != "ready" && status != "inProcess" && status != "projected" {
		return errors.New("non-existing status, choose one of three [ready, inProcess, projected]")
	}
	fmt.Printf("\nBuilding %s updated status to %s\n", b.name, status)
	b.status = status
	b.notifyAll()
	return nil
}
func (b *Building) sign(o Observer) {
	b.observerList = append(b.observerList, o)
}

func (b *Building) unsign(o Observer) {
	for ind, sub := range b.observerList {
		if o.getID() == sub.getID() {
			b.observerList = append(b.observerList[:ind], b.observerList[ind+1:]...)
		}
	}
}

func (b *Building) notifyAll() {
	wg := sync.WaitGroup{}
	for _, observer := range b.observerList {
		wg.Add(1)
		go observer.update(b.name, b.status, &wg)
	}
	wg.Wait()
}
