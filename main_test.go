package linkedList

import "testing"

func TestList_Insert(t *testing.T) {
	list := createList()

	list.insert("data1")

	if list.head.data != "data1" {
		t.Fatalf("unexpected head: %v", list.head)
	}
}

func TestList_Insert_InsertTwoDataPoints(t *testing.T) {
	list := createList()

	list.insert("data1")
	list.insert("data2")

	if list.head.nextNode.data != "data2" {
		t.Fatalf("unexpected head: %v", list.head)
	}
}
