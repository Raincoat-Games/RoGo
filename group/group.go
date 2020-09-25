//Package group provides functions to interact with Roblox groups.
package group

const endpoint = "https://groups.roblox.com/v1/groups/"

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
	return nil, nil
}
