'''
Problem Statement

Hermione Granger is lost in the Forbidden Forest while collecting some herbs for her magical potion. The forest is magical and has only 1 exit point which magically transports her back to the Hogwarts School of Wizardy and Witch Craft. 
Forest can be considered as a grid of NxM size. Each cell in the forest is either empty (represented by '.') or has a tree (represented by 'X'). Hermione can move through empty cell, but not through cells with tree on it. She can only travel LEFT, RIGHT, UP, DOWN. Her position in the forest is indicated by the marker 'M' and the location of the exit point is indicated by '*'. Top-left corner is indexed (0, 0).

.X.X......X
.X*.X.XXX.X
.XX.X.XM...
......XXXX.
In the above forest, Hermione is located at index (2, 7) and exit is at (1, 2). Each cell is indexed according to Matrix Convention

She starts her commute back to the exit and every time she encounters more than one option to move, she waves her wand and the correct path is illuminated and she proceeds in that way. It is guaranteed that there is only one path to each reachable cell from the starting cell. Can you tell us if she waved her wand exactly K times or not? Ron will be impressed if she is able to do so.

Input Format 
First line contains an integer T. T test cases follow. 
Each test case starts with two space separated integer N and M. 
Then next N lines contain a string each length of M which represents the forest. 
Last line of each single test case is integer K.

Output Format 
For each test case, if she could impress Ron then print "Impressed" otherwise print "Oops!".

All quotes used for the clarifications.

Constraints 
1 <= T <= 10 
1<=N, M<=100 
0 <= K <= 10000 
There is exactly one 'M' and one '*' in the graph. 
Exactly one path exists between 'M' and '*.'

Sample Input
'''

def has_way(ditu, i, j):
	result = []
	if i>0 and ditu[i-1][j] == '.':
		result.append([i-1,j])
	if i<len(ditu)-1 and ditu[i+1][j] == '.':
		result.append([i+1,j])
	if j>0 and ditu[i][j-1] == '.':
		result.append([i, j-1])
	if j<len(ditu[0])-1 and ditu[i][j+1] == '.':
		result.append([i, j+1])
	return result

def found(ditu, i, j):
	result = False
	if i>0 and ditu[i-1][j] == '*':
		result = True
	if i<len(ditu)-1 and ditu[i+1][j] == '*':
		result = True
	if j>0 and ditu[i][j-1] == '*':	
		result = True
	if j<len(ditu[0])-1 and ditu[i][j+1] == '*':
		result = True
	return result

def solve(ditu, turn, star, man):
	mx, my = man
	ditu[mx][my] = 'X'
	last_pos = [mx, my]
	way = has_way(ditu, mx, my)
	dirctions = []
	dirctions.append(last_pos)
	while len(way)!=0:
		next_step = way.pop()
		last_one = dirctions[len(dirctions)-1]
		while sum(map(lambda x,y: abs(x-y), next_step, last_one)) != 1:
			dirctions.pop()
			last_one = dirctions[len(dirctions)-1]
		#if next_step == [24, 14]:
		#	print dirctions, next_step, sum(map(lambda x,y: abs(x-y), next_step, last_one))
		dirctions.append(next_step)
		if found(ditu, next_step[0], next_step[1]):
			dirctions.append(next_step)
			break
		last_pos = next_step
		ditu[next_step[0]][next_step[1]] = 'O'
		way += has_way(ditu, next_step[0], next_step[1])
	return dirctions


t = int(raw_input())

for ti in xrange(t):
	n, m = map(int, raw_input().strip().split(" "))
	ar = []
	star_x, star_y, m_x, m_y = 0,0,0,0
	for ni in xrange(n):
		line = raw_input()
		ar.append([])
		for x in range(m):
			if line[x] == '*':
				star_x, star_y = ni, x
			if line[x] == 'M':
				m_x, m_y = ni, x 
			ar[ni].append(line[x])
	k = int(raw_input())

	#print ar, star_x, star_y, m_x, m_y
	dirs = solve(ar, k, [star_x, star_y], [m_x, m_y])
	count = 0

	for x in xrange(len(dirs)-1):
		point = dirs[x]
		ar[point[0]][point[1]] = '$'
		if point[0]-1 >=0 and [point[0]-1, point[1]] != dirs[x+1] and (ar[point[0]-1][point[1]] == '.' or ar[point[0]-1][point[1]] == 'O'):
			count += 1
			ar[point[0]][point[1]] = 'A'
			continue
		if point[0]+1 <= len(ar)-1 and [point[0]+1, point[1]] != dirs[x+1] and (ar[point[0]+1][point[1]] == '.' or ar[point[0]+1][point[1]] == 'O'):
			count += 1
			ar[point[0]][point[1]] = 'A'
			continue
		if point[1]-1 >=0 and [point[0], point[1]-1] != dirs[x+1] and (ar[point[0]][point[1]-1] == '.' or ar[point[0]][point[1]-1] == 'O'):
			count += 1
			ar[point[0]][point[1]] = 'A'
			continue
		if point[1]+1 <= len(ar[0])-1 and [point[0], point[1]+1] != dirs[x+1] and (ar[point[0]][point[1]+1] == '.' or ar[point[0]][point[1]+1] == 'O'):
			count += 1
			ar[point[0]][point[1]] = 'A'
			continue

	if count == k:
		print 'Impressed'
	else:
		print 'Oops!'

