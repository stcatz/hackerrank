/***
Problem Statement

Chinese Version
Russian Version

There are N bikers present in a city (shaped as a grid) having M bikes. All the bikers want to participate in the HackerRace competition, but unfortunately only K bikers can be accommodated in the race. Jack is organizing the HackerRace and wants to start the race as soon as possible. He can instruct any biker to move towards any bike in the city. In order to minimize the time to start the race, Jack instructs the bikers in such a way that first K bikes are acquired in the minimum time.

Every biker moves with a unit speed and one bike can be acquired by only one biker. A biker can proceed in any direction. Consider distance between bike and bikers as euclidean distance.

Jack would like to know the square of required time to start the race as soon as possible.

Input Format 
The first line contains three integers - N,M,K separated by a single space. 
Following N lines will contain N pairs of integers denoting the co-ordinates of N bikers. Each pair of integers is separated by a single space. The next M pairs of lines will denote the co-ordinates of M bikes.

Output Format 
A single line containing the square of required time.

Constraints 

Sample Input #00:

3 3 2
0 1
0 2
0 3
100 1
200 2 
300 3
Sample Output #00:

40000
Explanation #00: 
There's need for two bikers for the race. The first biker (0,1) will be able to reach the first bike (100,1) in 100 time units. The second biker (0,2) will be able to reach the second bike (200,2) in 200 time units. This is the most optimal solution and will take 200 time units. So output will be 2002 = 40000.

Copyright (c) 2015 HackerRank.
All Rights Reserved
Suggest Edits
***/

package main

import "fmt"
import "sort"

func main() {
	var n, m, k int
	var bikes, bikers [][2]int
	var distance []int

	fmt.Scanf("%d %d %d", &n, &m, &k)

	bikers = make([][2]int, n)
	bikes = make([][2]int, m)
	distance = make([]int, n*m)
	for i := 0; i < n+m; i++ {
		var x,y int
		fmt.Scanf("%d %d", &x, &y)
		if i<n {
			bikers[i][0] = x
			bikers[i][1] = y
		}else{
			bikes[i-n][0] = x
			bikes[i-n][1] = y
		}
	}

	for i := 0; i < n*m; i++ {
		x := i/n
		y := i%n

		dis := int((bikers[x][0] - bikes[y][0]) * (bikers[x][0] - bikes[y][0]) + (bikers[x][1] - bikes[y][1]) * (bikers[x][1] - bikes[y][1]))
		distance[i] = dis
	}
	fmt.Println(distance)
	sort.Ints(distance)
	fmt.Println(distance)
	
}
