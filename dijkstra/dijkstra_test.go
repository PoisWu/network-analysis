package dijkstra

import (
    "testing"
)

func creatExampleGraph() graph {
    as1 := 101
    
    as2 := 102
    as3 := 103
    as4 := 104
    as5 := 105
    as6 := 106

    mygraph := make(graph)

    mygraph[as1] = append(mygraph[as1], as2)
    mygraph[as1] = append(mygraph[as1], as3)

    mygraph[as2] = append(mygraph[as2], as1)
    mygraph[as2] = append(mygraph[as2], as3)
    mygraph[as2] = append(mygraph[as2], as5)


    mygraph[as3] = append(mygraph[as3], as1)
    mygraph[as3] = append(mygraph[as3], as2)
    mygraph[as3] = append(mygraph[as3], as6)

    mygraph[as4] = append(mygraph[as4], as6)

    mygraph[as5] = append(mygraph[as5], as2)

    mygraph[as6] = append(mygraph[as6], as3)
    mygraph[as6] = append(mygraph[as6], as4)
    return mygraph
}

func TestDijkstra(t *testing.T){
    mygraph := creatExampleGraph()

    dist, prev := dijkstra(mygraph)

    for k:= range(dist){
        t.Logf("Dist from as1 to %d is %d", k, dist[k])
    }
    for k:= range(prev){
        t.Logf("Prev from as1 to %d is %d", k, prev[k])
    }
    
}
