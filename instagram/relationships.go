package instagram

import(
	"fmt"
	"net/url"
)

type RelationshipsApi struct {
	Instagram *Instagram
}

//http://instagram.com/developer/endpoints/relationships/#get_users_follows
func (o *RelationshipsApi) Follows(userId string) ([]User, *Content, error) {
	var path = fmt.Sprintf("/users/%s/follows", userId)
	var item = new([]User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/relationships/#get_users_followed_by
func (o *RelationshipsApi) FollowedBy(userId string) ([]User, *Content, error) {
	var path = fmt.Sprintf("/users/%s/followed-by", userId)
	var item = new([]User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/relationships/#get_incoming_requests
func (o *RelationshipsApi) RequestedBy() ([]User, *Content, error) {
	var path = "/users/self/requested-by"
	var item = new([]User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return *item, data, err
}

//http://instagram.com/developer/endpoints/relationships/#get_relationship
func (o *RelationshipsApi) GetRelationship(userId string) (*Relationship, *Content, error) {
	var path = fmt.Sprintf("/users/%s/relationship", userId)
	var item = new(Relationship)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//http://instagram.com/developer/endpoints/relationships/#post_relationship
func (o *RelationshipsApi) PostRelationship(userId string, action string) (*Relationship, *Content, error) {
	var params = url.Values{}
	params.Set("action", action)

	var path = fmt.Sprintf("/users/%s/relationship", userId)
	var item = new(Relationship)
	data, err := o.Instagram.NewRequest(item, "POST", path, params, true)
	return item, data, err
}
