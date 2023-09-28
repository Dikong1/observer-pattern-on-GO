package main

import "fmt"

func main() {

	construction := newBuilding("ЖК Юрта")
	err1 := construction.updateAvailability("projectedd") // handling exception if another status name is entered
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	buyer1 := &Client{id: "paluan777@gmail.com"}
	buyer2 := &Client{id: "baidyn_uly@gmail.com"}
	buyer3 := &Client{id: "zhurttyn_balasy@gmail.com"}

	construction.sign(buyer1)
	construction.sign(buyer2)
	construction.sign(buyer3)

	err2 := construction.updateAvailability("inProcess")
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	construction.unsign(buyer1) // he realized he has no money to buy apartments in a such cool house

	err3 := construction.updateAvailability("ready")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

}
