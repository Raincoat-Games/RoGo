package main

import (
	"fmt"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"log"
)

func main() {
	acc := account.New(nil)
	g, err := group.Get(4953490, acc)
	fmt.Println(acc.SecurityCookie)
	if err != nil { log.Fatal(err.Error()) }
	//err = g.PostShout("[ENTRY] Join our communication server to get accepted into SCAR - /yxHDRfS")
	//if err != nil { log.Fatal(err.Error()) }
	shout, err := g.GetShout()
	if err != nil { log.Fatal(err.Error()) }
	fmt.Println(shout.Poster.UserID)
}
