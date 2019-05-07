package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

/*Response is similar to DistanceMatrixResponse from google maps api*/
type Response struct {
	Origins      []string `json:"origin_addresses"`
	Destinations []string `json:"destination_addresses"`
	Rows         []struct {
		Elements []struct {
			Status   string `json:"status"`
			Distance struct {
				Text  string `json:"text"`
				Value int    `json:"value"`
			} `json:"distance"`
			Duration struct {
				Value int    `json:"value"`
				Text  string `json:"text"`
			} `json:"duration"`
			Durintraffic struct {
				Value int    `json:"value"`
				Text  string `json:"text"`
			} `json:"duration_in_traffic"`
		} `json:"elements"`
	} `json:"rows"`
}

/*Edge is representative of origin and destination of an edge + weight for adjacency graphing*/
type Edge struct {
	Root   string
	Dest   string
	Weight int
}

/*Prims reads file of distance matrix response and generates MST, printing to stdout*/
func Prims(filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("failed to read file")
		log.Fatal()
	}
	var resp Response
	err = json.Unmarshal(b, &resp)
	if err != nil {
		fmt.Printf("json broke %v+", err)
		return
	}
	//Creating an adjacency matrix for all coordinates to others, including
	//adjMatrix := make([][]int, 10)
	namesMap := make(map[string]int)
	for x := 0; x < 10; x++ {
		namesMap[resp.Origins[x]] = x
	}
	adjMatrix := make([][]Edge, 10)
	for i := range adjMatrix {
		for x := 0; x < 10; x++ {
			if x != i {
				adjMatrix[i] = append(adjMatrix[i], Edge{
					Root:   resp.Origins[i],
					Dest:   resp.Destinations[x],
					Weight: resp.Rows[i].Elements[x].Distance.Value,
				})
			}
		}
	}

	var parkHeap ParkHeap
	inputSet := make(map[string]int)
	for i := range namesMap {
		inputSet[i] = math.MaxInt32
	}
	inputSet[resp.Origins[0]] = 0
	parkHeap.Push(ParkNode{
		resp.Origins[0],
		inputSet[resp.Origins[0]],
	})

	type TreeNode struct {
		name     string
		children []string
	}

	//Let's keep a set of origins mapped to edges
	tree := make(map[string]string)
	for i := range namesMap {
		tree[i] = "empty"
	}

	mstSet := make(map[string]bool)

	sumWeights := 0
	setSize := 10
	for setSize > 0 {

		currPark := parkHeap.Pop().(ParkNode)
		//Adding into tree
		if !mstSet[currPark.name] {
			mstSet[currPark.name] = true
			setSize--
			sumWeights += currPark.weight
			//For all neighbors of currPark.name
			for _, edge := range adjMatrix[namesMap[currPark.name]] {
				//If neighbor edge weight is less than current node value, reduce key
				if inputSet[edge.Dest] > edge.Weight && !mstSet[edge.Dest] {
					//Change weight in graph itself
					inputSet[edge.Dest] = edge.Weight
					//To store the parent
					if !mstSet[edge.Dest] {
						tree[edge.Dest] = edge.Root
					}

				}
				//If not already in tree, push onto heap
				if !mstSet[edge.Dest] {
					parkHeap.Push(ParkNode{
						name:   edge.Dest,
						weight: inputSet[edge.Dest],
					})
				}

			}
		}
	}

	FinalTree := make(map[string][]string)

	//For all nodes to be treated as roots
	for i := range namesMap {
		//Check all other nodes to see if they are children of FinalTree
		for j := range namesMap {
			//If parent node it points to (tree[j] is finaltree, append it as a child)
			if i == tree[j] {
				FinalTree[i] = append(FinalTree[i], j)
			}
		}
	}
	for i, j := range FinalTree {
		fmt.Println(i+" children: ", j)
	}
	fmt.Println(sumWeights)

}
