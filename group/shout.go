package group

import (
	"bytes"
	"encoding/json"
	"fmt"
	requests "github.com/Clan-Labs/RoGo/helpers"
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
	cookieJar, err := auth.NewJar(acc.SecurityCookie, endpoint) // Create JAR
	if err != nil {
		return err
	}
	type reqBody struct {
		Message string `json:"message"`
	}
	reqB := reqBody{Message: shout}
	bodyJson, err := json.Marshal(reqB) // Create & marshal body
	if err != nil {
		return err
	}

	//Create req
	client := &http.Client{Timeout: 10 * time.Second, Jar: cookieJar}
	// Create an authorized request
	req, err := requests.NewAuthorizedRequest(acc, endpoint, "PATCH", bytes.NewReader(bodyJson))
	if err != nil { return err }

	//Send Request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = res.Body.Close() }()

	// Check errors
	if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden  { // roblox returns 400 for unauthorized apparently?
		return errs.ErrUnauthorized
	}
	if res.StatusCode == http.StatusBadRequest {
		return errs.ErrBadRequest
	}

	return nil
}
