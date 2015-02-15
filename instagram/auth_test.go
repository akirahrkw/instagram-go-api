package instagram

import (
	"testing"
	"strings"	
)

func TestAuthorizeURL(t *testing.T) {
		var endpoint = "https://api.instagram.com/oauth/authorize"
		var result = strings.Contains(client.AuthorizeURL(), endpoint)
		if !result {
			t.Errorf("TestAuthorizeURL Error")
		}
}

func TestAuthorizeURLWithScope(t *testing.T) {
		var endpoint = "https://api.instagram.com/oauth/authorize"
		var url = client.AuthorizeURLWithScope([]string{"likes","comments","relationships"})
		var result = strings.Contains(url, endpoint)
		if !result {
			t.Errorf("TestAuthorizeURL Error")
		}
}

func TestRequestAccessToken(t *testing.T) {
	// remove this skip and set code if you want to test this api
	t.Skip("skipping")
	var code = ""

	data, err := client.RequestAccessToken(code)

	var _ = err

	if data.AccessToken == "" {
		t.Errorf("access token is wrong")
	}
}