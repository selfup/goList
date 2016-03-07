package linkedList

import "testing"

func TestList_CreateNode_CreatingANodeOnTheList(t *testing.T) {
	list := createList()

	list.insert("data1")

	if list.head.data != "data1" {
		t.Fatalf("unexpected head: %v", list.head)
	}
}
