package main

import (
	"fmt"
	"math"
	"time"
)

// leaner search
func linearSearch(haystack []int, needle int) bool {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}
	return false
}

// binary search <sorted array>
func binarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack)

	for low < high {
		mid := low - (high-low)/2
		value := haystack[mid]

		if value == needle {
			return true
		} else if value > needle {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return false
}

// binary search recursive version
[5,6]
func recursiveLinearSearch(haystack []int, needle int) bool {
  if len(haystack) == 0 {
    return false
  } else {
    value := recursiveLinearSearch(len(haystack) - 1, needle)
    if value == needle {
      return true
    }
  }


}
func main() {
}

// second version of binary search
func binarySearch1(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := haystack[mid]
		if guess == needle {
			return true
		} else if guess > needle {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}

// a = append(a[:i], a[i+1:]...)
// selection sort
func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

type RingBuffer struct {
	data       []*Data
	size       int
	lastInsert int
	nextRead   int
	emitTime   time.Time
}

type Data struct {
	Stamp time.Time
	Value string
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		data:       make([]*Data, size),
		size:       size,
		lastInsert: -1,
	}
}

func (r *RingBuffer) Insert(input Data) {
	r.lastInsert = (r.lastInsert + 1) % r.size
	r.data[r.lastInsert] = &input

	if r.nextRead == r.lastInsert {
		r.nextRead = (r.nextRead + 1) % r.size
	}
}

func (r *RingBuffer) Emit() []*Data {
	output := []*Data{}
	for {
		if r.data[r.nextRead] != nil {
			output = append(output, r.data[r.nextRead])
			r.data[r.nextRead] = nil
		}
		if r.nextRead == r.lastInsert || r.lastInsert == -1 {
			break
		}
		r.nextRead = (r.nextRead + 1) % r.size
	}
	return output
} // ring buffer ends here

// two crystal ball problem
func twoCrystalBalls(breaks []bool) int {
	jumpPoint := int(math.Floor(math.Sqrt(float64(len(breaks)))))
	i := jumpPoint
	for ; i < len(breaks); i += jumpPoint {
		if breaks[i] {
			break
		}
	}

	i -= jumpPoint

	for j := 0; j <= jumpPoint && i < len(breaks); i, j = i+1, j+1 {
		if breaks[i] {
			return i
		}
	}
	return -1
}

// bubble sort
func bubbleSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// queue linked list "first in first out (FIFO)"
type Node struct {
	value any
	next  *Node
}

type Queue struct {
	head   *Node
	tail   *Node
	Length int
}

func NewQueue() *Queue {
	return &Queue{head: nil, tail: nil, Length: 0}
}

func (q *Queue) enqueue(item any) {
	node := &Node{value: item, next: nil}
	q.Length++
	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}
	q.tail.next = node
	q.tail = node
}

func (q *Queue) deque() any {
	if q.Length == 0 {
		return "there is nothing to remove"
	} else {
		q.Length--
	}
	ch := q.head
	q.head = ch.next

	ch.next = nil
	return ch.value
}

func (q *Queue) peek() {
	if q.head != nil {
		fmt.Println("peeking", q.head.value)
		return
	}
	fmt.Println("there is no item in the queue")
}

// linked list stack (lifo)
type sNode struct {
	value    any
	previous *sNode
}

type Stack struct {
	head   *sNode
	Length int
}

func NewStack() *Stack {
	return &Stack{
		head:   nil,
		Length: 0,
	}
}

func (s *Stack) push(value any) {
	node := &sNode{value: value, previous: nil}
	s.Length++
	if s.head == nil {
		s.head = node
		return
	}
	node.previous = s.head
	s.head = node
}

func (s *Stack) pop() any {
	if s.Length == 0 {
		s.head = nil
		return s.head
	}
	s.Length--
	h := s.head
	s.head = h.previous
	return h.value
}

func (s *Stack) peek() {
	if s.head == nil {
		fmt.Println("nothing in the stack")
		return
	}
	fmt.Println(s.head.value)
}
