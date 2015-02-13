package instagram

import (
	"testing"
	"net/url"
)

func TestMediaGet(t *testing.T) {
	var expected = mediaId

	media, content, err := client.Media.Get(expected)
	isInvalidMetaData(content, err, t)

	if media.Id != expected {
		t.Errorf("expected media_id is wrong: %s", media.Id)
	}
}

func TestMediaShortcode(t *testing.T) {
	var expected = mediaId
	var shortcode = "xZaWLwvIp5"
	
	media, content, err := client.Media.Shortcode(shortcode)
	isInvalidMetaData(content, err, t)

	if media.Id != expected {
		t.Errorf("expected media_id is wrong: %s", media.Id)
	}
}

func TestMediaPopular(t *testing.T) {
	items, content, err := client.Media.Popular()
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of recent media is 0")
	}
}

func TestMediaSearch(t *testing.T) {
	var params = url.Values{}
	params.Set("lat", "1.284962")
	params.Set("lng", "103.858699")
	params.Set("distance", "1000")

	items, content, err := client.Media.Search(params)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of recent media is 0")
	}
}
