package tests

import (
	"fmt"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"log"
	"os"
	"testing"
)

func BenchmarkMain(m *testing.B) {
	acc := account.New("")
	_, err := group.Get(4953490, acc)
	if err != nil { log.Fatal(err.Error()) }
}

func TestMain(m *testing.M) {
	acc := account.New(os.Getenv("COOKIE"))
	g, err := group.Get(4953490, acc)
	if err != nil { log.Fatal(err.Error()) }
	//err = g.Exile(1505886708)
	//if err != nil { log.Fatal(err.Error()) }
	r, err := g.GetJoinRequests(1)
	if err != nil { log.Fatal(err.Error()) }
	for _, i := range r {
		fmt.Printf("Username: %s\nUserID: %d\nCreated: %v\n\n", i.Requester.Username, i.Requester.UserID, i.Created)
	}

}
