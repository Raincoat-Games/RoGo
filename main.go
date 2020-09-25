package main

import (
	"log"

	"github.com/Clan-Labs/RoGo/account"

	"github.com/Clan-Labs/RoGo/group"
)

func main() {

	acc := account.New("")

	//Check initial shout
	shout, err := group.GetShout(4099453, acc)
	if err != nil {
		log.Fatal(err)
	}
	println(shout.Content)

	//Change shout
	err = group.PostShout("RoGo goes brrrt", 4099453, acc)
	if err != nil {
		log.Fatal(err)
	}

	//Check shout
	shout, err = group.GetShout(4099453, acc)
	if err != nil {
		log.Fatal(err)
	}
	println(shout.Content)

}
