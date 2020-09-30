package main

import (
	"fmt"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"log"
)

func main() {
	acc := account.New("")
	g, err := group.Get(4953490, acc)
	if err != nil { log.Fatal(err.Error()) }
}

