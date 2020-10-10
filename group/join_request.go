package group

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Clan-Labs/RoGo/auth"
	"github.com/Clan-Labs/RoGo/errs"
	requests "github.com/Clan-Labs/RoGo/helpers"
	"net/http"
	"time"
)

func interiorAccept(r JoinRequest, group *Group) error {
	URI := fmt.Sprintf("https://groups.roblox.com/v1/groups/%d/join-requests/users/%d", group.Id, r.Requester.UserID)
	if !group.BotAccount.IsAuthenticated() { return errs.ErrRequiresCookie }
	cookieJar, err := auth.NewJar(group.BotAccount.SecurityCookie, URI) // Create JAR
	if err != nil { return err }
	client := &http.Client{Timeout: 10 * time.Second, Jar: cookieJar}

	type EmptyBody struct {} // We don't want to send a body here
	jsonBody, err := json.Marshal(EmptyBody{})
	if err != nil { return err }

	req, err := requests.NewAuthorizedRequest(group.BotAccount,
		URI, "POST", bytes.NewReader(jsonBody)) // Helper func for dealing with XCSRF

	if err != nil { return err }

	res, err := client.Do(req)
	if err != nil { return err }
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("unexpected status code '%d'", res.StatusCode))
	}
	return nil
}

func interiorDecline(r JoinRequest, group *Group) error {
	URI := fmt.Sprintf("https://groups.roblox.com/v1/groups/%d/join-requests/users/%d", group.Id, r.Requester.UserID)
	if !group.BotAccount.IsAuthenticated() { return errs.ErrRequiresCookie }
	cookieJar, err := auth.NewJar(group.BotAccount.SecurityCookie, URI) // Create JAR
	if err != nil { return err }
	client := &http.Client{Timeout: 10 * time.Second, Jar: cookieJar}

	type EmptyBody struct {} // We don't want to send a body here
	jsonBody, err := json.Marshal(EmptyBody{})
	if err != nil { return err }

	req, err := requests.NewAuthorizedRequest(group.BotAccount,
		URI,"DELETE", bytes.NewReader(jsonBody)) // Helper func for dealing with XCSRF

	if err != nil { return err }

	res, err := client.Do(req)
	if err != nil { return err }

	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("unexpected status code '%d'", res.StatusCode))
	}
	return nil
}

func (r JoinRequest) Accept() error {
	return interiorAccept(r, r.Group)
}

func (r JoinRequest) Decline() error {
	return interiorDecline(r, r.Group)
}