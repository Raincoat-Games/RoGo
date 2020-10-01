package tests

import (
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/group"
	"log"
	"testing"
)

func BenchmarkMain(m *testing.B) {
	acc := account.New("")
	_, err := group.Get(4953490, acc)
	if err != nil { log.Fatal(err.Error()) }
}
