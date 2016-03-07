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
		t.Fatalf("unexpected node: %v", list.head.nextNode.data)
	}
}

func TestList_Include_CanFindData(t *testing.T) {
	list := createList()

	list.insert("data1")
	list.insert("data2")
	list.insert("data3")

	search := list.include("data2")

	if search != true {
		t.Fatalf("unexpected bool: %v", search)
	}
}
