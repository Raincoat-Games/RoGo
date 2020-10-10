package group

import (
	"github.com/Clan-Labs/RoGo/account"
	"time"
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

// Roles represents the response given by GetGroupRoles
type Roles struct {
	GroupId 		int `json:"groupId"`
	Roles   		[]Role `json:"roles"`
}


// Role represents an item in Roles
type Role struct {
	Id          	int    `json:"id"`
	Name        	string `json:"name"`
	Rank        	int    `json:"rank"`
	MemberCount 	int    `json:"memberCount"`
}

// Represents a user's Role in a Group
type RobloxGroupData struct {
	RobloxGroup 	Group `json:"group"`
	RobloxRole  	Role  `json:"role"`
}

// UserGroupData represents the Groups a user is in, see RobloxGroupData
type UserGroupData struct {
	Data 			[]RobloxGroupData `json:"data"`
}

// Requester represents a subsection of a JoinRequest
type Requester struct {
	Username string `json:"username"`
	DisplayName string `json:"displayName"`
	UserID int `json:"userId"`
}

// JoinRequest represents a user who is pending to join a group
type JoinRequest struct {
	Group *Group
	Created time.Time `json:"created"`
	Requester Requester `json:"requester"`
}

type InteriorPost struct {
	User User `json:"user"`
	Role Role `json:"role"`
}

type GroupPost struct {
	Id int `json:"id"`
	Poster InteriorPost `json:"poster"`
	Body string `json:"body"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Group *Group
}


