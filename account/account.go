package account

import (
	"fmt"
	"net/http"
)

//Account represents a Roblox account with auth information.
type Account struct {
	SecurityCookie *http.Cookie
}

func (a Account) IsAuthenticated() bool {
	return fmt.Sprint(a.SecurityCookie) != ".ROBLOSECURITY="
}

//New creates a new Account
func New(securityCookie interface{}) *Account {
	if securityCookie == nil { securityCookie = "" }
	cookie := http.Cookie{Name: ".ROBLOSECURITY", Value: securityCookie.(string)}
	acc := Account{SecurityCookie: &cookie}
	return &acc
}
