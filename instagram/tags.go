package instagram

import (
	"fmt"
	"net/url"
)

type TagsApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/tags/#get_tags
func (o *TagsApi) Get(name string) (*Tag, *Content, error) {
	var path = fmt.Sprintf("/tags/%s", name)
	var item = new(Tag)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/tags/#get_tags_media_recent
func (o *TagsApi) RecentMedia(name string, params url.Values) ([]Media, *Content, error) {
	var path = fmt.Sprintf("/tags/%s/media/recent", name)
	var item = new([]Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/tags/#get_tags_search
func (o *TagsApi) Search(params url.Values) ([]Tag, *Content, error) {
	var path = "/tags/search"
	var item = new([]Tag)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}