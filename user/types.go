package user

import "time"

//The User struct provides information about a Roblox user.
type User struct {
	// BuildersClubMembershipType always returns None, and there is no premium endpoint which doesnt require auth
	BuildersClubMembershipType string `json:"buildersClubMembershipType"`
	UserID                     int    `json:"userId"`
	Username                   string `json:"username"`
}

type AdvancedUser struct {
	User
	Description				   string `json:"description"`
	tempCreated				   string `json:"created"`
	Created						time.Time
}
