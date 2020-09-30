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
	old, new, err := g.ChangeRank(1505886708, 1)
	if err != nil { log.Fatal(err.Error()) }
	fmt.Printf("Promoted %d from rank `%s` to rank `%s`\n", 1505886708, old.Name, new.Name)
	//for _, r := range roles {
	//	fmt.Println(r.Rank, r.Name)
	//}
	//g, err := group.Get(4953490, acc)
	//fmt.Println(acc.SecurityCookie)
	//if err != nil { log.Fatal(err.Error()) }
	////err = g.PostShout("[ENTRY] Join our communication server to get accepted into SCAR - /yxHDRfS")
	////if err != nil { log.Fatal(err.Error()) }
	//shout, err := g.GetShout()
	//if err != nil { log.Fatal(err.Error()) }
	//fmt.Println(shout.Poster.UserID)
}

