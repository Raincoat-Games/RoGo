package main

import (
	"fmt"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"log"
	"os"
)

func main() {
	acc := account.New(os.Getenv("COOKIE"))
	g, _ := group.Get(4953490, acc)
	old, new, err := g.Demote(1505886708)
	if err != nil { log.Fatal(err.Error()) }
	fmt.Printf("Demoted from %s to %s", old.Name, new.Name)
}
