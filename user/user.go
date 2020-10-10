package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Clan-Labs/RoGo/group"
	"net/http"
	"time"
)

func(u User) GetRankInGroups() ([]group.RobloxGroupData, error) {
	URI := fmt.Sprintf("https://groups.roblox.com/v1/users/%v/groups/roles", u.UserId)
	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", URI, bytes.NewReader([]byte{}))
	if err != nil { return nil, err }
	res, err := client.Do(req)
	if err != nil { return nil, err }
	defer func(){ _ = res.Body.Close() }()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("unexpected status code recieved `%v`",
			res.StatusCode))
	}
	var jsonRes group.UserGroupData
	err = json.NewDecoder(res.Body).Decode(&jsonRes)
	if err != nil { return nil, err }
	return jsonRes.Data, nil
}
