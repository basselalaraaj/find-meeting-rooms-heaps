package main

import (
	"container/heap"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(meetingTimes [][]int) int {
	if len(meetingTimes) == 0 {
		return 0
	}

	sort.Slice(meetingTimes, func(i, j int) bool {
		return meetingTimes[i][0] < meetingTimes[j][0]
	})

	freeRooms := &IntHeap{meetingTimes[0][1]}
	heap.Init(freeRooms)

	for _, meeting := range meetingTimes[1:] {
		if (*freeRooms)[0] <= meeting[0] {
			heap.Pop(freeRooms)
		}
		heap.Push(freeRooms, meeting[1])
	}
	return freeRooms.Len()
}

func main() {
	meeting_times := [][]int{{3, 4}, {3, 9}, {5, 11}, {2, 8}, {8, 20}, {11, 15}}
	println(minMeetingRooms(meeting_times))
}
