'''
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

'''

m,n = map(int, raw_input().strip().split(" "))

ar = []

for x in xrange(m):
	ar.append([])
	line = raw_input()
	for c in line:
		ar[x].append(c)


hm = []
vm = []

for i in xrange(m):
	hm.append([])
	for j in xrange(n):
		hm[i].append([])
		for k in xrange(n):
			hm[i][j].append(False)

for i in xrange(m):
	for j in xrange(n):
		if ar[i][j] == 'x':
			continue
		for k in xrange(j,n):
			if k == j:
				hm[i][j][k] = True
			elif hm[i][j][k-1] == True and ar[i][k] == '.':
				hm[i][j][k] = True


for i in xrange(n):
	vm.append([])
	for j in xrange(m):
		vm[i].append([])
		for k in xrange(m):
			vm[i][j].append(False)


for i in xrange(n):
	for j in xrange(m):
		if ar[j][i] == 'x':
			continue
		for k in xrange(j,m):
			if k == j:
				vm[i][j][k] = True
			elif vm[i][j][k-1] == True and ar[k][i] == '.':
				vm[i][j][k] = True


max_result = 0

'''
for x in xrange(m):
	for y in xrange(n):
		print ar[x][y],
	print ''
'''

for i in range(m):
	for j in range(n):
		if ar[i][j] == 'x':
			continue
		for xi in range(i+1, m):
			for xj in range(j+1, n):
				if ar[xi][xj] == 'x':
					continue
				result = xi - i + xj - j
				if result > max_result and hm[i][j][xj] and hm[xi][j][xj] and vm[j][i][xi] and vm[xj][i][xi]:
					max_result = result


if max_result > 0:
	print max_result * 2
else:
	print 'impossible'

