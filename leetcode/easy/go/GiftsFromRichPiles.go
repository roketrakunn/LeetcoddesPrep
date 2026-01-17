package go_test

import (
	"container/heap"
	"math"
)


type MaxHeap []int

func (h MaxHeap) Len() int { 
	return  len(h)
}

func (h MaxHeap) Less(i,j int) bool {
	return h[i] > h[j]

}

func (h MaxHeap) Swap(i , j int ) { h[i]  , h[j] = h[j] , h[i]}

func (h *MaxHeap) Push(x any) { 
	*h = append(*h , x.(int))
}

func (h *MaxHeap) Pop() any { 
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]

	return x
}

func pickGifts(gifts []int, k int) int64 {

	h := MaxHeap(gifts)
	heap.Init(&h)

	for range k { 
		x := heap.Pop(&h).(int)
		y := int( math.Sqrt(float64(x)))
		heap.Push(&h , y)
	}

	var sum int64 

	for h.Len() > 0 { 
		sum +=int64(heap.Pop(&h).(int))
	}
	return sum
    
}
