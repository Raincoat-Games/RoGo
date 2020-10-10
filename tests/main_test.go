package tests

import (
	"fmt"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"github.com/Clan-Labs/RoGo/user"
	"log"
	"testing"
	"time"
)

func BenchmarkMain(m *testing.B) {
	acc := account.New("")
	_, err := group.Get(1, acc)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestMain(m *testing.M) {
	//usr := user.User{UserId: 1}
	//gdata, _ := usr.GetRankInGroups()
	//for _, data := range gdata {
	//	fmt.Println((*data.RobloxGroup.Owner).Username)
	//}
	//acc := account.New(os.Getenv("COOKIE"))
	u, err := user.NewUserFromID(171029413)
	if err != nil { log.Fatal(err.Error()) }
	fmt.Println(u.Created.Format(time.RFC822Z))
	//err = g.Exile(1505886708)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//r, errch, err := g.GetGroupPosts(1)
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	//for {
	//	select {
	//	case err, open := <-errch:
	//		if !open {
	//			fmt.Println("Error Channel has been closed")
	//			return
	//		}
	//		fmt.Println(err.Error())
	//	case res, open := <-r:
	//		if !open {
	//			fmt.Println("Channel has been closed")
	//			return
	//		}
	//		for _, v := range res {
	//			fmt.Println(v.Body)
	//		}
	//	}
	//}
}
