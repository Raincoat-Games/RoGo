package group

import "github.com/Clan-Labs/RoGo/account"

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

	BotAccount *account.Account `json:"-"` // Don't encode account
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
	Content 		string `json:"body"`
	Created 		string `json:"created"`
	Updated 		string `json:"updated"`
	Poster  		*User
}

type Roles struct {
	GroupId 		int `json:"groupId"`
	Roles   		[]Role `json:"roles"`
}

type Role struct {
	Id          	int    `json:"id"`
	Name        	string `json:"name"`
	Rank        	int    `json:"rank"`
	MemberCount 	int    `json:"memberCount"`
}

type RobloxGroupData struct {
	RobloxGroup 	Group `json:"group"`
	RobloxRole  	Role  `json:"role"`
}

type UserGroupData struct {
	Data 			[]RobloxGroupData `json:"data"`
}
