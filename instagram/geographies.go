package instagram

import (
	"fmt"
	"net/url"
)

type GeographiesApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/geographies/#get_geographies_media_recent
func (o *GeographiesApi) RecentMedia(geoId string, params url.Values) ([]Media, *Content, error) {
	var path = fmt.Sprintf("/geographies/%s/media/recent", geoId)
	var item = new([]Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}
