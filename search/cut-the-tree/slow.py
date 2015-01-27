'''
Problem Statement

Atul is into graph theory, and he is learning about trees nowadays. He observed that the removal of an edge from a given tree T will result in formation of two separate trees T1 and T2.

Each vertex of the tree T is assigned a positive integer. Your task is to remove an edge, such that, the *Tree_diff* of the resultant trees is minimized. *Tree_diff* is defined as

 F(T) = Sum of numbers written on each vertex of a Tree T
 Tree_diff(T) = abs(F(T1) - F(T2))
Input Format 
The first line will contain an integer N, i.e., the number of vertices in the tree.
The next line will contain N integers separated by a single space, i.e., the values assigned to each of the vertices.
The next (N-1) lines contain pair of integers separated by a single space and denote the edges of the tree.
In the above input, the vertices are numbered from 1 to N.

Output Format 
A single line containing the minimum value of *Tree_diff*.

Constraints 


Sample Input

6
100 200 100 500 100 600
1 2
2 3
2 5
4 5
5 6
Sample Output

400
Explanation

Originally, we can represent tree as

         1(100)
          \
           2(200)
          / \
    (100)5   3(100)
        / \
  (500)4   6(600)
Cutting the edge at 1 2 would result in Tree_diff = 1500-100 = 1400 
Cutting the edge at 2 3 would result in Tree_diff = 1500-100 = 1400 
Cutting the edge at 2 5 would result in Tree_diff = 1200-400 = 800 
Cutting the edge at 4 5 would result in Tree_diff = 1100-500 = 600 
Cutting the edge at 5 6 would result in Tree_diff = 1000-600 = 400

Hence, the answer is 400.

Copyright (c) 2015 HackerRank.
All Rights Reserved
Suggest Edits
Choose a translation

704 hackers have submitted code
Share

Download PDF
Current Buffer (saved locally, editable)     

'''

class Node(object):
	"""docstring for ClassName"""
	def __init__(self, value=None, parent = None):
		super(Node, self).__init__()
		self.value = value
		self.parent = parent
		self.sons = []


n = int(raw_input())
ar = map(int, raw_input().strip().split(" "))
edge = []
for x in xrange(n-1):
	edge.append(map(int, raw_input().strip().split(" ")))

min_sum = 2**32 - 1
all_sum = sum(ar)
a=[0]
a = a + (map(lambda x: Node(x), ar))

for x in edge:
	a[x[0]].sons.append(x[1])
	a[x[1]].sons.append(x[0])

for x in edge:
	sumed = []
	sumed.append(x[1])

	other_sons = filter(lambda y: y!=x[0], a[x[1]].sons)
	while len(other_sons) != 0:
		sumed += other_sons
		tmp = []
		for m in other_sons:
			tmp += a[m].sons

		temp = filter(lambda z: sumed.count(z) == 0, tmp)
		other_sons = temp
	
	this_sum = 0

	for i in sumed:
		this_sum += a[i].value

	if min_sum > abs(all_sum-this_sum - this_sum):
		min_sum = abs(all_sum-this_sum - this_sum)

print min_sum

