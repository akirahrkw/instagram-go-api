package instagram

import (
	"fmt"
	"net/url"
	"strconv"
)

type UserApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/users/#get_users
func (o *UserApi) GetUser(userId string) (*User, *Content, error) {
	var path = fmt.Sprintf("/users/%s", userId)
	var item = new(User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/users/#get_users
func (o *UserApi) GetSelf() (*User, *Content, error) {
	var path = "/users/self"
	var item = new(User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/users/#get_users_feed
func (o *UserApi) GetSelfFeed(count int, minId string, maxId string) ([]Media, *Content, error) {
	var path = "/users/self/feed"
	var item = new([]Media)

	var params = url.Values{}
	params.Set("count", strconv.Itoa(count))
	if minId != "" {
		params.Set("min_id", minId)
	}
	if maxId != "" {
		params.Set("max_id", maxId)
	}

	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content ,err
}

//http://instagram.com/developer/endpoints/users/#get_users_media_recent_with_client_id
func (o *UserApi) GetRecentMedia(userId string, count int, minId string, maxId string) ([]Media, *Content, error) {
	var path = fmt.Sprintf("/users/%s/media/recent", userId)
	var item = new([]Media)

	var params = url.Values{}
	params.Set("count", strconv.Itoa(count))
	if minId != "" {
		params.Set("min_id", minId)
	}
	if maxId != "" {
		params.Set("max_id", maxId)
	}

	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content ,err
}

//http://instagram.com/developer/endpoints/users/#get_users_feed_liked
func (o *UserApi) GetLikedMedia(count int) ([]Media, *Content, error) {
	var path = "/users/self/media/liked"
	var item = new([]Media)

	var params = url.Values{}
	params.Set("count", strconv.Itoa(count))

	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content ,err
}

//http://instagram.com/developer/endpoints/users/#get_users_search
func (o *UserApi) Search(query string, count int) ([]User, *Content, error) {
	var path = "/users/search"
	var item = new([]User)
	var params = url.Values{}
	params.Set("q", query)
	params.Set("count", strconv.Itoa(count))

	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content ,err

}
