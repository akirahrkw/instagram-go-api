package instagram

import (
	"fmt"
	"net/url"
)

type MediaApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/media/#get_media
func (o *MediaApi) Get(mediaId string) (*Media, *Content, error) {
	var path = fmt.Sprintf("/media/%s", mediaId)
	var item = new(Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_media_by_shortcode
func (o *MediaApi) Shortcode(shortcode string) (*Media, *Content, error) {
	var path = fmt.Sprintf("/media/shortcode/%s", shortcode)
	var item = new(Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_media_popular
func (o *MediaApi) Search(params url.Values) ([]Media, *Content, error) {
	var path = "/media/search/"
	var item = new([]Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_media_popular
func (o *MediaApi) Popular() ([]Media, *Content, error) {
	var path = "/media/popular/"
	var item = new([]Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}

