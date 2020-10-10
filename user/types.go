package user

type User struct {
	BuildersClubMembershipType string `json:"buildersClubMembershipType"`
	UserId 					   int	  `json:"userId"`
	Username				   string `json:"username"`
	DisplayName				   string `json:"displayName"`
}
