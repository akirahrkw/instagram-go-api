package instagram

import (
	"testing"
	"fmt"
	"os"
)
var _ = fmt.Printf
var client *Instagram
var selfId string
var mediaId string

func TestMain(m *testing.M) {
	selfId = "" // you have to set selfId
	mediaId = "889858275048983161_263873"

	// set your api setting
	client = NewClient(func(config *Config){
		config.RedirectUri = "http://localhost"
		config.ClientId = ""
		config.ClientSecret = ""
	})
	client.SetAccessToken(os.Getenv("INSTAGRAM_ACCESS_TOKEN"))
	os.Exit(m.Run())
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
