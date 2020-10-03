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
	r, errch, err := g.GetJoinRequests()
	if err != nil { log.Fatal(err.Error()) }
	for {
		select {
		case err, open := <-errch:
			if !open {
				fmt.Println("Error Channel has been closed")
				return
			}
			fmt.Println(err.Error())
		case res, open := <- r:
			if !open {
				fmt.Println("Channel has been closed")
				return
			}
			for _, v := range res {
				fmt.Println(v.Requester.Username)
			}
		}
	}
}
