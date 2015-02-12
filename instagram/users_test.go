package instagram

import (
	"testing"
	"net/url"
)

func TestUsersGet(t *testing.T) {
	var expected = selfId

	user, content, err := client.Users.Get(expected)
	isInvalidMetaData(content, err, t)

	if user.Id != expected {
		t.Errorf("expected user_id is wrong: %s", user.Id)
	}
}

func TestUsersSelf(t *testing.T) {
	var expected = selfId

	user, content, err := client.Users.Self()
	isInvalidMetaData(content, err, t)

	if user.Id != expected {
		t.Errorf("expected user_id is wrong: %s", user.Id)
	}
}

func TestUsersSelfFeed(t *testing.T) {
	var params = url.Values{}
	params.Set("count", "10")

	items, content, err := client.Users.SelfFeed(params)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of feeds is 0")
	}
}

func TestUsersRecentMedia(t *testing.T) {
	items, content, err := client.Users.RecentMedia(selfId, nil)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of recent media is 0")
	}
}

func TestUsersLikedMedia(t *testing.T) {
	items, content, err := client.Users.LikedMedia(nil)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of liked media is 0")
	}
}

func TestUsersSearch(t *testing.T) {
	items, content, err := client.Users.Search("japan", 5)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of search result is 0")
	}
}

