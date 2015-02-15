package instagram

import (
	"testing"
)

func TestRelationshipsFollows(t *testing.T) {
	items, content, err := client.Relationships.Follows(selfId)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of follows is 0")
	}
}

func TestRelationshipsFollowedBy(t *testing.T) {
	items, content, err := client.Relationships.FollowedBy(selfId)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of followedBy is 0")
	}
}

func TestRelationshipsRequestedBy(t *testing.T) {
	items, content, err := client.Relationships.RequestedBy()
	isInvalidMetaData(content, err, t)

	var _ = items

	if content.Meta.Code != 200 {
		t.Errorf("Http status is not 200")
	}
}

func TestRelationshipsGetRelationship(t *testing.T) {
	var userId = "1574083"
	relationship, content, err := client.Relationships.GetRelationship(userId)
	isInvalidMetaData(content, err, t)

	if relationship.OutgoingStatus == "" || relationship.IncomingStatus == "" {
		t.Errorf("OutgoingStatus is empty")
	}
}

func TestRelationshipsPostRelationship(t *testing.T) {
	var userId = "1574083"

	// follow
	relationship, content, err := client.Relationships.PostRelationship(userId, "follow")
	isInvalidMetaData(content, err, t)

	if relationship.OutgoingStatus == "" {
		t.Errorf("OutgoingStatus is empty")
	}

	// unfollow
	relationship, content, err = client.Relationships.PostRelationship(userId, "unfollow")
	isInvalidMetaData(content, err, t)

	if relationship.OutgoingStatus == "" {
		t.Errorf("OutgoingStatus is empty")
	}
}
