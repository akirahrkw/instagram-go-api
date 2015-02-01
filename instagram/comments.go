package instagram

import (
	"fmt"
	"net/url"
)

type CommentsApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/comments/#get_media_comments
func (o *CommentsApi) Get(mediaId string) ([]Comment, *Content, error) {
	var path = fmt.Sprintf("/media/%s/comments", mediaId)
	var item = new([]Comment)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/comments/#post_media_comments
func (o *CommentsApi) Post(mediaId string, text string) (*Content, error) {
	var params = url.Values{}
	params.Set("text", text)

	var path = fmt.Sprintf("/media/%s/comments", mediaId)
	data, err := o.Instagram.NewRequest(nil, "POST", path, params, true)
	return data, err
}

//http://instagram.com/developer/endpoints/comments/#delete_media_comments
func (o *CommentsApi) Del(mediaId string, commentId string) (*Content, error) {
	var path = fmt.Sprintf("/media/%s/comments/%s", mediaId, commentId)
	data, err := o.Instagram.NewRequest(nil, "DELETE", path, nil, true)
	return data, err
}
