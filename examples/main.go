package main

import (
	"fmt"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"github.com/jpfuentes2/go-env"
	"log"
	"os"
)

func main() {
	err := env.ReadEnv("./env")
	if err != nil { log.Fatal(err.Error()) }
	SimpleGroupInfo()
	AuthenticatedGroupFuncs()
}


func SimpleGroupInfo() {
	acc := account.New(nil)
	g, err := group.Get(5, acc)
	if err != nil { log.Fatal(err.Error()) }
	fmt.Printf("GROUP NAME: '%s'\nGROUP DESC: '%s'\n", g.Name, g.Description)
}

func AuthenticatedGroupFuncs() {
	acc := account.New(os.Getenv("COOKIE"))
	g, err := group.Get(4953490, acc)
	if err != nil { log.Fatal(err.Error()) }
	old, curr, err := g.ChangeRank(1505886708, -1)
	if err != nil { log.Fatal(err.Error()) }
	fmt.Printf("The user has been demoted from `%s` to `%s`\n", old.Name, curr.Name)
}