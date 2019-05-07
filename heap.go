package main

/*ParkNode represents a vertex for our prim's algorithm.*/
type ParkNode struct {
	name   string
	weight int
}

/*ParkHeap is an implemented heap interface.*/
type ParkHeap []ParkNode

func (heap ParkHeap) Len() int           { return len(heap) }
func (heap ParkHeap) Less(i, j int) bool { return heap[i].weight < heap[j].weight }
func (heap ParkHeap) Swap(i, j int)      { heap[i], heap[j] = heap[j], heap[i] }

/*Push appends an element onto heap*/
func (heap *ParkHeap) Push(x interface{}) {
	*heap = append(*heap, x.(ParkNode))
}

/*Pop removes minimum element from heap*/
func (heap *ParkHeap) Pop() interface{} {
	old := *heap
	n := len(old)
	x := old[n-1]
	*heap = old[0 : n-1]
	return x
}

/*Peek checks minimum element in heap*/
func (heap *ParkHeap) Peek() interface{} {
	old := *heap
	n := len(old)
	return old[n-1]
}
