package instagram

import (
	"testing"
	"net/url"
)

func TestLocationsGet(t *testing.T) {
	var expected = "1"

	location, content, err := client.Locations.Get(expected)
	isInvalidMetaData(content, err, t)

	if location.Id != expected {
		t.Errorf("expected location id is wrong: %v", location.Id)
	}
}

func TestLocationsRecentMedia(t *testing.T) {
	items, content, err := client.Locations.RecentMedia("1", nil)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of location's recent media is 0")
	}
}

func TestLocationsSearch(t *testing.T) {
	var params = url.Values{}
	params.Set("lat", "48.858844")
	params.Set("lng", "2.294351")
	params.Set("distance", "1000")	

	items, content, err := client.Locations.Search(params)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of search locations is 0")
	}
}
