
/**
package go_test

type ListNode struct{ 
	val int 
	next *ListNode
}

func sort(head *ListNode) *ListNode {

	if head == nil || head.next == nil{ 
		return head
	}
	//split list into two halves 

	mid := getMid(head)
	right := mid.next
	mid.next = nil

	leftSorted := sort(head)
	rightSorted := sort(right)

	return  merge(leftSorted, rightSorted)
}

func getMid(head *ListNode) *ListNode { 
	slow , fast := head , head.next

	for fast != nil &&  fast.next != nil { 
		slow = slow.next
		fast = fast.next.next
	}
	return slow 
}


func merge(a , b *ListNode) *ListNode { 
	dummy := &ListNode{}
	tail := dummy 

	for a != nil &&  b != nil { 
		if a.val <= b.val { 
			tail.next = a
			a = a.next
		} else { 
			tail.next= b
			b = b.next
		}
		tail = tail.next
	}

	if a != nil { 
		tail.next = a
	} else { 
		tail.next = b
	}
	return dummy.next
}
*/
