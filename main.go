package main

import (
	"log"

	"github.com/Clan-Labs/RoGo/group"
)

func main() {

	shout, err := group.GetShout(4099453, "")
	if err != nil {
		log.Fatal(err)
	}

	if shout != nil {
		println(shout.Content)
	}
}
