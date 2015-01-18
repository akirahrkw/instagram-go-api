package instagram

import (
	"fmt"
)

type UserApi struct {
	Instagram *Instagram
}

func (o *UserApi) GetUser(userId string) (*User, *Content, error) {
	var path = fmt.Sprintf("/users/%s", userId)
	var item = new(User)
	data, err := o.Instagram.NewRequest("GET", path, item, true)
	return item, data, err
}

func (o *UserApi) GetSelf() (*User, *Content, error) {
	var path = "/users/self"
	var item = new(User)
	data, err := o.Instagram.NewRequest("GET", path, item, true)
	return item, data, err
}

func (o *UserApi) GetSelfFeed() {
	var path = "/users/self/feed"
	var item = new([]Media)
	data, err := o.Instagram.NewRequest("GET", path, item, true)
	return *item, data ,err
}

func (o *UserApi) GetRecentMedia(userId string) ([]Media, *Content, error) {
	var path = fmt.Sprintf("/users/%s/media/recent", userId)
	var item = new([]Media)
	data, err := o.Instagram.NewRequest("GET", path, item, true)
	return *item, data ,err
}

func (o *UserApi) GetLikedMedia() {
	var path = "/users/self/media/liked"
	var item = new([]Media)
	data, err := o.Instagram.NewRequest("GET", path, item, true)
	return *item, data ,err
}

func (o *UserApi) Search(query string) {
	//TODO
}
