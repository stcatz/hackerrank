/***
Problem Statement

Mr K has a rectangular land (m x n). There are marshes in the land where the fence cannot hold. Mr K wants you to find the perimeter of the largest rectangular fence that can be built on this land.

Input format

First line contains m and n. The next m lines contains n characters describing the state of the land. 'x' (ascii value: 120) if it is a marsh and '.' ( ascii value:46) otherwise.

Constraints

2<=m,n<=500

Output Format

Output contains a single integer - the largest perimeter. If the rectangular fence cannot be built, print "impossible" (without quotes)

Sample Input:1

4 5
.....
.x.x.
.....
.....
Output

14
Fence can be put up across the entire land owned by Mr K. The perimeter is 2 * (4-1) + 2 * (5-1) = 14.

Sample Input:2

2 2
.x
x.
Output

impossible
We need minimum of 4 points to place the 4 corners of the fence. Hence, impossible.

Sample Input:3

2 5
.....
xxxx.
Output

impossible
Copyright (c) 2015 HackerRank.
All Rights Reserved
Suggest Edits
352 hackers have submitted code
Share

Download PDF
Current Buffer (saved locally, editable)  
***/


package main

import "fmt"

const (
	MAX = 500
)

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)

	ar := make([]string, m)

	type flag [MAX][MAX][MAX]bool

	var hm, vm flag

	for i := 0; i < m; i++ {
		fmt.Scanln(&ar[i])
	}

	for i := 0; i < MAX; i++ {
		for j := 0; j < MAX; j++ {
			for k := 0; k < MAX; k++ {
				hm[i][j][k] = false
				vm[i][j][k] = false
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ar[i][j] == 'x' {
				continue
			}
			for k := j; k < n; k++ {
				if k == j {
					hm[i][j][k] = true
				}else if hm[i][j][k-1] && ar[i][k] == '.' {
					hm[i][j][k] = true
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ar[j][i] == 'x' {
				continue
			}
			for k := j; k < m; k++ {
				if k == j {
					vm[i][j][k] = true
				}else if vm[i][j][k-1] && ar[k][i] == '.' {
					vm[i][j][k] = true
				}
			}
		}
	}


	max_result := 0



	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ar[i][j] == 'x' {
				continue
			}

			for xi := i+1; xi < m; xi++ {
				for xj := j+1; xj < n; xj++ {
					if ar[xi][xj] == 'x' {
						continue
					}
					result := xi - i + xj - j
					if result > max_result && hm[i][j][xj] && hm[xi][j][xj] && vm[j][i][xi] && vm[xj][i][xi] {
						max_result = result
					}
					
				}
				
			}
			
		}
		
	}


	if max_result > 0 {
		fmt.Println(max_result)
	}else{
		fmt.Println("impossible")
	}

}