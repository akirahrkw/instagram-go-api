package instagram

import (
	"fmt"
	"net/url"
	// "strconv"
)

type MediaApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/media/#get_media
func (o *MediaApi) GetMedia(mediaId string) (*Media, *Content, error) {
	var path = fmt.Sprintf("/media/%s", mediaId)
	var item = new(Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_media_by_shortcode
func (o *MediaApi) GetShortcode(shortcode string) (*Media, *Content, error) {
	var path = fmt.Sprintf("/media/shortcode/%s", shortcode)
	var item = new(Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_media_popular
func (o *MediaApi) GetSearch(lat string, lng string, distance string, minTimestamp string, maxTimestamp string) ([]Media, *Content, error) {

	var params = url.Values{}
	if lat != "" {
		params.Set("lat", lat)
	}
	if lng != "" {
		params.Set("lng", lng)
	}
	if distance != "" {
		params.Set("distance", distance)
	}
	if minTimestamp != "" {
		params.Set("min_timestamp", minTimestamp)
	}
	if maxTimestamp != "" {
		params.Set("max_timestamp", maxTimestamp)
	}

	var path = "/media/search/"
	var item = new([]Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/media/#get_media_popular
func (o *MediaApi) GetPopular() ([]Media, *Content, error) {
	var path = "/media/popular/"
	var item = new([]Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}

