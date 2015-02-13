package instagram

import (
	"testing"
	"net/url"
)

func TestTagsGet(t *testing.T) {
	var expected = "test"

	tag, content, err := client.Tags.Get(expected)
	isInvalidMetaData(content, err, t)

	if tag.Name != expected {
		t.Errorf("expected tag name is wrong: %s", tag.Name)
	}
}

func TestTagsRecentMedia(t *testing.T) {	
	items, content, err := client.Tags.RecentMedia("test", nil)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of liked media is 0")
	}
}

func TestTagsSearch(t *testing.T) {
	var params = url.Values{}
	params.Set("q", "test")

	items, content, err := client.Tags.Search(params)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of search tags is 0")
	}
}
