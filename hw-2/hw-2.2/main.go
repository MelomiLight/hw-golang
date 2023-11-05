package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}
/*
Сложность этой функции составляет O(m + n), где m и n - это длины списков list1 и list2, 
так как я прохожу по обоим спискам один раз, сравниваю их значения и потом строю новый отсортированный список.
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var startHead *ListNode = new(ListNode)
	var preHead *ListNode = startHead

	// Пока есть элементы в обоих списках, сравниваю значения и добавляю минимальное значение в объединенный список
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			startHead.Next = list1
			list1 = list1.Next
		} else {
			startHead.Next = list2
			list2 = list2.Next
		}
		startHead = startHead.Next
	}

	// Если один из списков закончился, добавляю оставшиеся элементы другого списка в объединенный список
	if list1 == nil {
		startHead.Next = list2
	} else if list2 == nil {
		startHead.Next = list1
	}

	return preHead.Next
}
//Просто для принта листа
func printLinkedList(head *ListNode) {
    current := head
    for current != nil {
        fmt.Printf("%d -> ", current.Val)
        current = current.Next
    }
    fmt.Println("nil")
}

func main() {
  
    list1 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}
    list2 := &ListNode{2, &ListNode{4, &ListNode{6, nil}}}


    mergedList := mergeTwoLists(list1, list2)

  
    fmt.Println("Merged List:")
    printLinkedList(mergedList)
}