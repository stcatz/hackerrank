/***
Problem Statement

You are given a tree where each node is labeled from 1, 2, ..., n. How many similar pairs(S) are there in this tree?

A pair (A,B) is a similar pair iff

node A is the ancestor of node B
abs(A - B) <= T.
Input format: 
The first line of the input contains two integers n and T. This is followed by n-1 lines each containing two integers si and ei where node si is a parent to node ei.

Output format: 
Output a single integer which denotes the number of similar pairs in the tree

Constraints:

Sample Input:

5 2
3 2
3 1
1 4
1 5
Sample Output:

4
Explanation: 
The similar pairs are: (3, 2) (3, 1) (3, 4) (3, 5)

You can have a look at the tree image here

Copyright (c) 2015 HackerRank.
All Rights Reserved
Suggest Edits
298 hackers have submitted code
Share

Download PDF
Current Buffer (saved locally, editable)     

***/


package main

import "fmt"

type Node struct {
    index int // its value
    sons []int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}else{
		return x
	}
}

func (node *Node)pre_order(nodes []Node, min_t int) ([]int, int){
	//fmt.Println(node.index)
	var sons []int
	count := 0

	for i := 0; i < len(node.sons); i++ {
		sons = append(sons, node.sons[i])
		son, count1 := nodes[node.sons[i]].pre_order(nodes, min_t)
		for i := 0; i < len(son); i++ {
			sons = append(sons, son[i])
		}
		count += count1
	}

	//fmt.Println(node.index, sons, min_t)

	for i := 0; i < len(sons); i++ {
		if abs(node.index - sons[i]) <= min_t {
			count++
		}
	}

	return sons, count
}

func main() {
	var n, t, head int
	fmt.Scanf("%d %d", &n, &t)

	node_list := make([]Node, n+1)

	for i := 1; i < n+1; i++ {
		node_list[i].sons = make([]int, 0)
		node_list[i].index = i
	}

	for i := 0; i < n-1; i++ {
		var f, s int
		fmt.Scanf("%d %d", &f, &s)
		if i==0 {
			head = f
		}

		node_list[f].sons =append(node_list[f].sons, s)
	}

	//for i := 1; i < n+1; i++ {
	//	fmt.Println(node_list[i], head)
	//}

	_, finalcount := node_list[head].pre_order(node_list, t)

	fmt.Println(finalcount)

}











