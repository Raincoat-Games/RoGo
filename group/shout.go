package group

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/auth"
	"github.com/Clan-Labs/RoGo/errs"
)

//The GetShout function returns the supplied group's shout.
func GetShout(groupId int, acc interface{}) (*Shout, error) {
	group, err := Get(groupId, acc)
	if err != nil {
		return nil, err
	}

	//Check if unauthorized
	if group.Shout == nil {
		return nil, errs.ErrUnauthorized
	}

	return group.Shout, nil
}

var postShoutEndpoint = "https://groups.roblox.com/v1/groups/%v/status"

//The PostShout function posts a group shout.
func PostShout(shout string, groupId int, acc *account.Account) error {

	endpoint := fmt.Sprintf(postShoutEndpoint, groupId)

	//Create Jar
	cookieJar, err := auth.NewJar(acc.SecurityCookie, endpoint)
	if err != nil {
		return err
	}

	type reqBody struct {
		Message string `json:"message"`
	}

	//Marshal body
	reqB := reqBody{Message: shout}
	bodyJson, err := json.Marshal(reqB)
	if err != nil {
		return err
	}

	//Create req
	client := &http.Client{Timeout: 10 * time.Second, Jar: cookieJar}
	req, err := http.NewRequest("PATCH", endpoint, bytes.NewReader(bodyJson))
	if err != nil {
		return err
	}

	//Get XCSRF
	csrf, err := auth.GetCSRF(acc)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-TOKEN", csrf)

	//Send Request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()

	//Check for errors
	if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
		return errs.ErrUnauthorized
	}

	return nil
}
