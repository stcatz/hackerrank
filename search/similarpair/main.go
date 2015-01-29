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

var n, t, head int

var dep, tree []int
var node_list [][]int

func query(index, left, right, begin, end int) int {
	//fmt.Println(left, right, begin, end)
	if right < begin || left > end {
		return 0
	}
	if left <= begin && right >= end {
		return tree[index]
	}
	mid := begin + (end-begin)/2

	return query(index*2, left, right, begin, mid) + query(index*2+1, left, right, mid+1, end)
}

func update(index, node, val, begin, end int) {
	if node < begin || node > end {
		return
	}
	if begin <= node && node <= end {
		tree[index] += val
	}
	if begin == node && end == node {
		return
	}

	mid := begin + (end-begin)/2
	update(index*2, node, val, begin, mid)
	update(index*2+1, node, val, mid+1, end)
}

func similarPairs(root int) int {
	var l, r int
	if root-t > 1 {
		l = root - t
	} else {
		l = 1
	}
	if root+t > n {
		r = n
	} else {
		r = root + t
	}

	sim := query(1, l, r, 1, n)
	//fmt.Println(sim)
	update(1, root, 1, 1, n)
	for i := 0; i < len(node_list[root]); i++ {
		sim += similarPairs(node_list[root][i])
	}
	update(1, root, -1, 1, n)
	return sim
}

func main() {
	fmt.Scanf("%d %d", &n, &t)
	dep = make([]int, n+1)
	node_list = make([][]int, n+1)
	tree = make([]int, n*4)

	for i := 0; i < n-1; i++ {
		var f, s int
		fmt.Scanf("%d %d", &f, &s)
		dep[s]++
		node_list[f] = append(node_list[f], s)
	}

	for i := 0; i < len(dep); i++ {
		if dep[i] == 0 {
			head = i
		}
	}

	result := similarPairs(head)

	fmt.Println(result)

}
