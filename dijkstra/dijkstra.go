package dijkstra


import(
    "container/heap"
)

// I need a structure represent a graph
// number of using ASN: 90,000
// from 1 - 65534

// I want to build a shortest path diagram from a given AS.

// I need a parser to read from the file. 

type as int

// graph is a hash map of ASID: [other AS connect to this as]
type graph map[as][]as 

func dijkstra(g graph, as_source as) (map[as]int,map[as]as) {
    res_dist := make(map[as]int)
    res_prev := make(map[as]as)
    as2item := make(map[as]*Item)
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    for v := range (g) {
        var item Item
        item.as = v
        item.distance = -1
        if v == as_source {
            item.distance = 0
        }
        heap.Push(&pq, &item, )
        as2item[v] = &item
    }

    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)

        res_dist[item.as] = item.distance
        d := res_dist[item.as]
        for _, v := range(g[item.as]){
            v_item := as2item[v]
            v_origin_dist := v_item.distance
            if (v_origin_dist == -1  || v_origin_dist > d + 1){
                pq.update(as2item[v], v, d+1)
                res_prev[v] = item.as
            }
        }
    }
    return res_dist, res_prev
}
