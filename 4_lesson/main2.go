package main

import (
	"fmt"
)

func main(){
	a := []int{1,2,3,3,3}
	s := createList(a)
	
	// qw := s

	// for i:=0; i < len(a); i++{
	// 	fmt.Println(qw)
	// 	qw = qw.Next
	// }
	
	printListNode(*s, len(a))

	q := deleteDuplicates(s)

	printListNode(*q, 3)

}

type ListNode struct {
	Val int
	Next *ListNode
}


func createList(arr []int) *ListNode{
	w := ListNode{}

	if len(arr) != 0{
		l := arr[1:]
		w := ListNode{Val: arr[0], Next: createList(l)}
		return &w
	}

	return &w
}


func printListNode(s ListNode, alen int){
	qw := s

	for i:=0; i < alen; i++{
		fmt.Println(qw)
		qw = *qw.Next
	}

	fmt.Println("-----end----")
}


// func deleteDuplicates(arr ListNode, alen int) int{
// 	qw := arr
// 	for i:=0; i < alen; i++{
// 		//fmt.Println(qw.Val)
// 		e := qw.Val
// 		qw2 := qw
// 		for i:=0; i < alen; i++{
// 			if e == qw2.Next.Val{
// 				qw.Next = qw.Next.Next
// 				alen -= 1
// 			}else{
// 				qw2 = *qw2.Next
// 			}
// 		}
// 		alen -= 1
// 		qw = *qw.Next
// 	}
// 	printListNode(qw, 3)
// 	return 0
// }


func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }

    curr := head

    for curr.Next != nil {
        if curr.Val == curr.Next.Val {
            curr.Next = curr.Next.Next
        } else {
            curr = curr.Next
        }
    }
    return head
}
