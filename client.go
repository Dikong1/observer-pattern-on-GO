package main

import (
	"fmt"
	"sync"
)

type Client struct {
	id string
}

func (c *Client) update(itemName string, status string, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()
	if status == "ready" {
		fmt.Printf("Sending message \"Its ready for check in to building %s\" to buyer`s email %s \n", itemName, c.id)
	} else if status == "inProcess" {
		fmt.Printf("Sending message \"Building %s still under construction\" to buyer`s email %s \n", itemName, c.id)
	}
}

func (c *Client) getID() string {
	return c.id
}
