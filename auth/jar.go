package auth

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

//NewJar creates a new *cookiejar.Jar with the supplied *http.Cookie
func NewJar(securityCookie *http.Cookie, endpoint string) (*cookiejar.Jar, error) {
	//Create cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil { return nil, err }

	//Parse URL
	URI, err := url.Parse(endpoint)
	if err != nil { return nil, err }

	//Create cookies
	jar.SetCookies(URI, []*http.Cookie{securityCookie})

	return jar, nil
}
