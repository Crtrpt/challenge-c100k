package container

import (
	"testing"
)

type testFilter struct {
}

func (t *testFilter) Filter(v any) bool {
	return true
}

func TestLinkList(t *testing.T) {
	link := NewLinkList()
	link.PushHead(20)
	link.PushHead(30)
	link.PushTail(30)
	link.PushTail(30)
}
