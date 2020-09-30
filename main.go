package main

import (
	"fmt"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"log"
)

func main() {
	acc := account.New("_|WARNING:-DO-NOT-SHARE-THIS.--Sharing-this-will-allow-someone-to-log-in-as-you-and-to-steal-your-ROBUX-and-items.|_3859AA717BF4E96D00FE062B68A2E72F19A2C81FD70B0FB51A64C43CCFF4A3DDDEC8EE57CDBF534B905C4293E2809E95F4C00F0C3C71DAA762DEB75057A08B32C219F7B0C37C9C4A7724C1C16339DA519529AB148F477504F997B4E6B8264C95F7B18D75D89E7329E749C5621E82F8A57B8E80A803D8A602DCCDB5411585A2FCD25F72EF079D9471FBD14B46CBEF6DCBB907484E08C2D7F3B80795C16403739B4393D23DDAFC5473940E2E41A4B774E4FFB64D4DCE0DE4EC335AF8917F1B728D02BC370363AD1537622365FB2DC81DDFB60FD7DA87AAD43CA507DA5F149D81E160D985DCB0EE0A22DA4CBFA2BE352E5135E8AFC3AF92C943B39FC107D5BB7DFAC1EA154C8EFA3240064F404E1E9C6E76E21D0BD2C54E6F8EC10F5B6DEAD62455183BB5203560671B6CDBAEEE44F64E1C973D137D")
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

