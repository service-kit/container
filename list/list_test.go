package list

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	lis := NewList()
	lis.Init()
	lis.PushBack(1)
	fmt.Println("PushBack 1")
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nDelete First")
	lis.Delete(lis.First())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nPushFront 2")
	lis.PushFront(2)
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nPushBack 3")
	lis.PushBack(3)
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nInsert 4 Before First")
	lis.Insert(4, lis.First())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nInsertBefore 5 Before First")
	lis.InsertBefore(5, lis.First())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nInsertAfter 6 After First")
	lis.InsertAfter(6, lis.First())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nInsertAfter 7 After End")
	lis.InsertAfter(7, lis.End())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nMoveBefore End Before")
	lis.MoveBefore(lis.End())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nMoveAfter First After")
	lis.MoveAfter(lis.First())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nMoveToFront End")
	lis.MoveToFront(lis.End())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println("\nMoveToBack First")
	lis.MoveToBack(lis.First())
	for iter := lis.First(); iter != lis.End().Next(); iter = iter.Next() {
		fmt.Printf("%v ", iter.Value())
	}
	fmt.Println()
}
