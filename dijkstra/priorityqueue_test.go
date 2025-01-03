package dijkstra

import (
    "testing"
    "container/heap"
)

func examplequeue() PriorityQueue{
	items := map[int]int{
        100: 1, 101: 2, 102: -1,
	}

	// Create a distance queue, put the items in it, and
	// establish the distance queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for num_as, distance := range items {
		pq[i] = &Item{
            as:    as{id: num_as},
			distance: distance,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
    return pq
}

func TestPriorityqueue(t *testing.T) {
    pq := examplequeue()
    
    item1:= heap.Pop(&pq).(*Item)
    t.Logf("Priority %d - AS: %d", item1.distance, item1.as )
    if item1.distance != 1 {
        t.Fatalf("Priority error, item distance should be 1 but get %d", 
            item1.distance)
    }
    item2:= heap.Pop(&pq).(*Item)
    t.Logf("Priority %d - AS: %d", item2.distance, item2.as )
    if item2.distance != 2 {
        t.Fatalf("Priority error, item2 distance should be 1 but get %d", 
            item2.distance)
    }
    item3:= heap.Pop(&pq).(*Item)
    t.Logf("Priority %d - AS: %d", item3.distance, item3.as )
    if item3.distance != -1 {
        t.Fatalf("Priority error, item3 distance should be 1 but get %d", 
            item3.distance)
    }

    if (pq.Len() != 0){
        t.Fatalf("The length of queue should be 0 but get %d", pq.Len())
    }
}

func TestUpdate(t *testing.T){
    pq := examplequeue()
    heap.Push(&pq, &Item{
        as: as{id:103}, 
        distance: 0},
    )

    item0:= heap.Pop(&pq).(*Item)
    t.Logf("Priority %d - AS: %d", item0.distance, item0.as )
    if item0.distance != 0 {
        t.Fatalf("Priority error, item distance should be 0 but get %d", 
            item0.distance)
    }

    item1:= heap.Pop(&pq).(*Item)
    t.Logf("Priority %d - AS: %d", item1.distance, item1.as )
    if item1.distance != 1 {
        t.Fatalf("Priority error, item distance should be 1 but get %d", 
            item1.distance)
    }
    item2:= heap.Pop(&pq).(*Item)
    t.Logf("Priority %d - AS: %d", item2.distance, item2.as )
    if item2.distance != 2 {
        t.Fatalf("Priority error, item2 distance should be 1 but get %d", 
            item2.distance)
    }
    item3:= heap.Pop(&pq).(*Item)
    t.Logf("Priority %d - AS: %d", item3.distance, item3.as )
    if item3.distance != -1 {
        t.Fatalf("Priority error, item3 distance should be 1 but get %d", 
            item3.distance)
    }



}

