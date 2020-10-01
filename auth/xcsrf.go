package auth

import (
	"net/http"
	"time"

	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/errs"
)

const endpoint = "https://auth.roblox.com/v2/logout"

//GetCSRF gets the CSRF token.
func GetCSRF(acc *account.Account) (string, error) {
	//Create Jar
	cookieJar, err := NewJar(acc.SecurityCookie, endpoint)
	if err != nil { return "", err }

	//Create req
	client := &http.Client{Timeout: 10 * time.Second, Jar: cookieJar}
	req, err := http.NewRequest("POST", endpoint, nil)
	if err != nil { return "", err }

	//Set header
	req.Header.Set("Accept", "application/json")

	//Send Request
	res, err := client.Do(req)
	if err != nil { return "", err }
	defer func() { _ = res.Body.Close() }()

	//Get token
	csrf := res.Header.Get("x-csrf-token")
	if csrf != "" { return csrf, nil
	} else { return "", errs.DidNotReceiveCSRF }
}
