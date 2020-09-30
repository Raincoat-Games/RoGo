//Package group provides functions to interact with Roblox groups.
package group

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	requests "github.com/Clan-Labs/RoGo/helpers"
	"net/http"
	"strconv"
	"time"

	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/auth"

	"github.com/Clan-Labs/RoGo/errs"
)

const endpoint = "https://groups.roblox.com/v1/groups/"

var (
	ErrGroupDoesntExist = errors.New("group doesn't exist")
)

//The Group struct provides information about a Roblox group.
type Group struct {
	Description        string `json:"description"`
	Id                 int    `json:"id"`
	IsBuildersClubOnly bool   `json:"isBuildersClubOnly"`
	MemberCount        int    `json:"memberCount"`
	Name               string `json:"name"`
	PublicEntryAllowed bool   `json:"publicEntryAllowed"`
	Shout              *Shout `json:"shout"`
	Owner              *User  `json:"owner"`

	BotAccount          *account.Account `json:"-"` // Don't encode account
}

//The User struct provides information about a Roblox user.
type User struct {
	BuildersClubMembershipType string `json:"buildersClubMembershipType"`
	DisplayName                string `json:"displayName"`
	UserID                     int    `json:"userId"`
	Username                   string `json:"username"`
}

//The Shout struct provides information about a group shout.
type Shout struct {
	Content string `json:"body"`
	Created string `json:"created"`
	Updated string `json:"updated"`
	Poster  *User
}

func (c Group) PostShout(shout string) error {
	var postShoutEndpoint = "https://groups.roblox.com/v1/groups/%v/status"
	endpoint := fmt.Sprintf(postShoutEndpoint, c.Id)
	// Check if account was provided
	if !c.BotAccount.IsAuthenticated() { return errors.New("this endpoint requires a valid cookie") }
	cookieJar, err := auth.NewJar(c.BotAccount.SecurityCookie, endpoint) // Create JAR
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
	req, err := requests.NewAuthorizedRequest(c.BotAccount, endpoint, "PATCH", bytes.NewReader(bodyJson))
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
	} else if res.StatusCode == http.StatusBadRequest {
		return errs.ErrBadRequest
	}

	return nil
}

func (c Group) GetShout() (*Shout, error) {
	//Check if unauthorized
	if c.Shout == nil {
		return nil, errs.ErrUnauthorized
	}
	return c.Shout, nil
}

//The Get function retrieves info about a Roblox group.
func Get(groupId int, acc *account.Account) (*Group, error) {

	//Make endpoint
	groupIdString := strconv.Itoa(groupId)
	URI := endpoint + groupIdString

	//Get account
	if acc == nil { acc = account.New("") }
	//Create Jar
	cookieJar, err := auth.NewJar(acc.SecurityCookie, endpoint)
	if err != nil {
		return nil, err
	}

	//Create req
	client := &http.Client{Timeout: 10 * time.Second, Jar: cookieJar}
	req, err := http.NewRequest("GET", URI, nil)
	if err != nil {
		return nil, err
	}

	//Send Request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = res.Body.Close() }()

	//Check if exists
	if res.StatusCode == http.StatusBadRequest {
		return nil, ErrGroupDoesntExist
	} else if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
		return nil, errs.ErrUnauthorized
	}

	//Parse Response Body
	var group *Group
	err = json.NewDecoder(res.Body).Decode(&group)
	group.BotAccount = acc
	if err != nil {
		return nil, err
	}

	return group, nil
}
