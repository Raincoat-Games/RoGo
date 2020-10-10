package group

import (
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/user"
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
	Owner              *user.User  `json:"owner"`

	BotAccount *account.Account `json:"-"` // Don't encode account
}

//The Shout struct provides information about a group shout.
type Shout struct {
	Content 		string `json:"body"`
	Created 		string `json:"created"`
	Updated 		string `json:"updated"`
	Poster  		*user.User
}


// Role represents a Group Role
type Role struct {
	Id          	int    `json:"id"`
	Name        	string `json:"name"`
	Rank        	int    `json:"rank"`
	MemberCount 	int    `json:"memberCount"`
}

// UserGroupRelation represents a user's Role in a Group
type UserGroupRelation struct {
	RobloxGroup 	Group `json:"group"`
	RobloxRole  	Role  `json:"role"`
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

// interiorPost is the struct that allows for User and Role to be embedded in Post
type interiorPost struct {
	User user.User `json:"user"`
	Role Role `json:"role"`
}

// Post is the decoded value of a Post on a group wall
type Post struct {
	Id int `json:"id"`
	Poster interiorPost `json:"poster"`
	Body string `json:"body"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	Group *Group
}


