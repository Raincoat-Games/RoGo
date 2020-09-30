package requests

import (
	"bytes"
	"github.com/Clan-Labs/RoGo/account"
	"github.com/Clan-Labs/RoGo/auth"
	"net/http"
)

func NewAuthorizedRequest(acc *account.Account, endpoint string, method string, reqB *bytes.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, endpoint, reqB)
	if err != nil { return nil, err }
	csrf, err := auth.GetCSRF(acc)
	if err != nil { return nil, err }
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-CSRF-TOKEN", csrf)
	return req, nil
}
