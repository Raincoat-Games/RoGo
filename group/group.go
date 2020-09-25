//Package group provides functions to interact with Roblox groups.
package group

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
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
	Shout              string `json:"shout"`
	Owner              *Owner `json:"owner"`
}

//The Owner struct provides information about a group owner.
type Owner struct {
	BuildersClubMembershipType string `json:"buildersClubMembershipType"`
	DisplayName                string `json:"displayName"`
	UserID                     int    `json:"userId"`
	Username                   string `json:"username"`
}

//The Get function retrieves info about a Roblox group.
func Get(groupId int) (*Group, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	groupIdString := strconv.Itoa(groupId)

	//Create req
	req, err := http.NewRequest("GET", endpoint+groupIdString, nil)
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
	}

	//Parse Response Body
	var group *Group
	err = json.NewDecoder(res.Body).Decode(&group)
	if err != nil {
		return nil, err
	}

	return group, nil
}
