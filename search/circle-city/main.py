'''
Problem Statement

Roy lives in a city that is circular in shape on a 2D plane. The city center is located at origin (0,0) and it has suburbs lying on the lattice points (points with integer coordinates). The city Police Department Headquarters can only protect those suburbs which are located strictly inside the city. The suburbs located on the border of the city are still unprotected. So the police department decides to build at most k additional police stations at some suburbs. Each of these police stations can protect the suburb it is located in.

Given the radius of the city, Roy has to determine if it is possible to protect all the suburbs.

Input Format 
The first line of input contains integer t, t testcases follow. 
Each of next t lines contains two space separated integers: r, the square of the radius of the city and k, the maximum number of police stations the headquarters is willing to build.

Constraints
Output Format 
For each test case, print in a new line possible if it is possible to protect all the suburbs, otherwise print impossible.

Sample Input

5
1 3
1 4
4 4
25 11
25 12
Sample Output

impossible
possible
possible
impossible
possible
Timelimits

Timelimits for this challenge are given here

Copyright (c) 2015 HackerRank.
All Rights Reserved
Suggest Edits
Choose a translation

1032 hackers have submitted code
Share

Download PDF
Current Buffer (saved locally, editable)     

'''

import math

t = int(raw_input())

for x in xrange(t):
	r, k = map(int, raw_input().strip().split(" "))

	radius = math.sqrt(r)
	if radius != int(radius):
		radius += 1

	edge = 0
	for i in xrange(0,int(radius)):
		other = math.sqrt(r - (i ** 2))
		if other == int(other):
			edge += 1

	if edge*4 <= k:
		print 'possible'
	else:
		print 'impossible'
