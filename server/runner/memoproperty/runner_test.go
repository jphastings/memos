package memoproperty_test

import (
	"reflect"
	"testing"

	storepb "github.com/usememos/memos/proto/gen/store"
	"github.com/usememos/memos/server/runner/memoproperty"
)

func TestGetMemoPropertyFromContent(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    *storepb.MemoPayload_Property
	}{
		{"Tag surrounded by spaces",
			"hello #tag world",
			&storepb.MemoPayload_Property{Tags: []string{"tag"}}},
		{"Tag followed by punctuation",
			"hello #world.",
			&storepb.MemoPayload_Property{Tags: []string{"world"}}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			property, err := memoproperty.GetMemoPropertyFromContent(test.content)
			if err != nil {
				t.Error(err)
				return
			}

			if !reflect.DeepEqual(property, test.want) {
				t.Errorf("Fail extracting properties: (expected) %v != %s (actual)",
					test.want, property)
			}
		})
	}
}
