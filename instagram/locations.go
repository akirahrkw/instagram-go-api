package instagram

import (
	"fmt"
	"net/url"
)

type LocationsApi struct {
	Instagram *Instagram	
}

//http://instagram.com/developer/endpoints/locations/#get_locations
func (o *LocationsApi) Get(locationId string) (*Location, *Content, error) {
	var path = fmt.Sprintf("/locations/%s", locationId)
	var item = new(Location)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/locations/#get_locations_media_recent
func (o *LocationsApi) RecentMedia(locationId string, params url.Values) ([]Media, *Content, error) {
	var path = fmt.Sprintf("/locations/%s/media/recent", locationId)
	var item = new([]Media)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/locations/#get_locations_search
func (o *LocationsApi) Search(params url.Values) ([]Location, *Content, error) {
	var path = "/locations/search"
	var item = new([]Location)
	data, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, data, err
}