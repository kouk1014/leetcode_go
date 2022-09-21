package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	test()
}

func test() {
	v1 := []int{9, 9, 9, 9, 9, 9, 9}
	v2 := []int{9, 9, 9, 9}
	l1 := &ListNode{}
	l2 := &ListNode{}
	cl1 := l1
	cl2 := l2
	for _, v := range v1 {
		cl1.Next = &ListNode{}
		cl1 = cl1.Next
		cl1.Val = v
	}
	for _, v := range v2 {
		cl2.Next = &ListNode{}
		cl2 = cl2.Next
		cl2.Val = v
	}
	l1 = l1.Next
	l2 = l2.Next
	x := addTwoNumbers(l1, l2)
	fmt.Println(x)
}

//思路
//同步循环链表,对每个node的值相加存入l1
//直到一个链表结尾
//若l1 l2 长度相同 返回l1
//若l1 长度大于l2  返回l1
//若l1 长度小于l2  l1最后节点指向l2剩余节点 ，返回l1
//期间要注意进位
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cl1 := l1
	cl2 := l2
	ten := false
	for cl1 != nil {
		v := cl1.Val + cl2.Val
		if ten {
			v += 1
			ten = false
		}
		if v >= 10 {
			ten = true
			v = v - 10
		}
		cl1.Val = v
		if cl1.Next == nil || cl2.Next == nil {
			if cl1.Next == nil && cl2.Next == nil {
				if ten {
					cl1.Next = &ListNode{
						Val: 1,
					}
				}
				return l1
			}
			if cl1.Next == nil && cl2.Next != nil {
				cl1.Next = cl2.Next
				for ten {
					if cl2.Next == nil {
						cl2.Next = &ListNode{Val: 1}
						break
					}
					cl2 = cl2.Next
					cl2.Val += 1
					ten = false
					if cl2.Val >= 10 {
						ten = true
						cl2.Val -= 10
					}
				}
				return l1
			}
			for ten {
				if cl1.Next == nil {
					cl1.Next = &ListNode{Val: 1}
					break
				}
				cl1 = cl1.Next
				cl1.Val += 1
				ten = false
				if cl1.Val >= 10 {
					ten = true
					cl1.Val -= 10
				}
			}
			return l1
		}

		cl1 = cl1.Next
		cl2 = cl2.Next

	}
	return l1
}
