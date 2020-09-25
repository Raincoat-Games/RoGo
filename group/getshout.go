package group

//The GetShout function returns the supplied group's shout.
//Returns nil if shout doesn't exist or the request was unauthorized.
func GetShout(groupId int, cookie string) (*Shout, error) {
	group, err := Get(groupId, cookie)
	if err != nil {
		return nil, err
	}

	return group.Shout, nil
}
