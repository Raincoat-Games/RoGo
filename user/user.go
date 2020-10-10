package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (u User) ToUser() *User {
	return &User{
		BuildersClubMembershipType: "None",
		UserID:                     u.UserID,
		Username:                   u.Username,
	}
}

func NewUserFromID(id int) (*AdvancedUser, error) {
	URI := fmt.Sprintf("https://users.roblox.com/v1/users/%v", id)
	client := &http.Client{Timeout: time.Second*10}
	req, err := http.NewRequest("GET", URI, bytes.NewReader([]byte{}))
	if err != nil { return nil, err }
	res, err := client.Do(req)
	if err != nil { return nil, err }
	defer res.Body.Close()
	//type TempUser struct { Id int `json:"Id"`; Username string `json:"Username"`}
	var NewUser = new(AdvancedUser)
	var TempUser map[string]interface{}
	json.NewDecoder(res.Body).Decode(&TempUser)
	if username, ok := TempUser["name"]; ok {
		NewUser.Username = username.(string)
		NewUser.UserID = id
		NewUser.Description = TempUser["description"].(string)
		NewTime, err := time.Parse("2006-01-02T15:04:05Z0700", TempUser["created"].(string))
		if err != nil { return nil, err }
		NewUser.Created = NewTime
	}

	return NewUser, nil
}

