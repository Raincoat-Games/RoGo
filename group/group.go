//Package group provides functions to interact with Roblox groups.
package group

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"

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

//The Get function retrieves info about a Roblox group.
func Get(groupId int, securityToken string) (*Group, error) {

	//Create cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	//Create cookies
	cookie := http.Cookie{Name: ".ROBLOSECURITY", Value: securityToken}
	groupIdString := strconv.Itoa(groupId)
	URI, _ := url.Parse(endpoint + groupIdString)
	jar.SetCookies(URI, []*http.Cookie{&cookie})

	//Create req
	client := &http.Client{Timeout: 10 * time.Second, Jar: jar}
	req, err := http.NewRequest("GET", URI.String(), nil)
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
	if err != nil {
		return nil, err
	}

	return group, nil
}
