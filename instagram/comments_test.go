package instagram

import (
	"testing"
)

var commentMediaId = "860587270891932468_263873"

func TestCommentsGet(t *testing.T) {
	items, content, err := client.Comments.Get(commentMediaId)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of comments is 0")
	}
}
