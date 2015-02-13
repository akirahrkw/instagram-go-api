package instagram

import (
	"testing"
)

var likedMediaId = "889858275048983161_263873"

func TestLikesGet(t *testing.T) {
	items, content, err := client.Likes.Get(likedMediaId)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of users is 0")
	}
}

func TestLikesPost(t *testing.T) {	
	content, err := client.Likes.Post(likedMediaId)
	isInvalidMetaData(content, err, t)

	if content.Meta.Code != 200 {
		t.Errorf("Http status is not 200")
	}
}

func TestLikesDel(t *testing.T) {
	content, err := client.Likes.Del(likedMediaId)
	isInvalidMetaData(content, err, t)

	if content.Meta.Code != 200 {
		t.Errorf("Http status is not 200")
	}
}
