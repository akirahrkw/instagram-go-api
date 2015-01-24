package instagram

import (
	"fmt"
)

type LikesApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/likes/#get_media_likes
func (o *LikesApi) Get(mediaId string) ([]User, *Content, error) {
	var path = fmt.Sprintf("/media/%s/likes", mediaId)
	var item = new([]User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/likes/#post_likes
func (o *LikesApi) Post(mediaId string) (*Content, error) {
	var path = fmt.Sprintf("/media/%s/likes", mediaId)
	data, err := o.Instagram.NewRequest(nil, "POST", path, nil, true)
	return data, err
}

//http://instagram.com/developer/endpoints/likes/#delete_likes
func (o *LikesApi) Del(mediaId string) (*Content, error) {
	var path = fmt.Sprintf("/media/%s/likes", mediaId)
	data, err := o.Instagram.NewRequest(nil, "DELETE", path, nil, true)
	return data, err
}
