package account

import (
	"net/http"
)

//Account represents a Roblox account with auth information.
type Account struct {
	SecurityCookie *http.Cookie
}

//Default is a blank Account
var Default = New("")

//New creates a new Account
func New(securityCookie string) *Account {
	cookie := http.Cookie{Name: ".ROBLOSECURITY", Value: securityCookie}
	acc := Account{SecurityCookie: &cookie}
	return &acc
}
