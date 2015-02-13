package instagram

import (
	"testing"
	"fmt"
	"strings"
	"os"
)
var _ = fmt.Printf
var client *Instagram
var selfId string
var mediaId string

func TestMain(m *testing.M) {
	selfId = "" // you have to set selfId
	mediaId = "889858275048983161_263873"

	client = NewClient(func(config *Config){
		config.RedirectUri = "http://localhost"
	})
	client.SetAccessToken(os.Getenv("INSTAGRAM_ACCESS_TOKEN"))
	os.Exit(m.Run())
}

func TestAuthorizeURL(t *testing.T) {
		var endpoint = "https://api.instagram.com/oauth/authorize"

		var result = strings.Contains(client.AuthorizeURL(), endpoint)
		if !result {
			t.Errorf("TestAuthorizeURL Error")
		}

		var url = client.AuthorizeURLWithScope([]string{"likes","comments","relationships"})
		result = strings.Contains(url, endpoint)
		if !result {
			t.Errorf("TestAuthorizeURL Error")
		}
}

// helper
func isInvalidMetaData(content *Content, err error, t *testing.T) {
	if content.Meta.Code > 299 {
		t.Errorf("Status Error: %v", content.Meta.Code)
	}
	if err != nil {
		t.Errorf("Response Error: %v", err)
	}
}
