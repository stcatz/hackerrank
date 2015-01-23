'''
Problem Statement

You have a rectangular board consisting of N rows, numbered from 1 to N, and M columns, numbered from 1 to M. The top left is (1,1) and the bottom right is (N,M). Initially - at time 0 - there is a coin on the top-left cell of your board. Each cell of your board contains one of these letters:

*: Exactly one of your cells has letter '*'.

U: If at time t, the coin is on the cell(i,j) and cell(i,j) has letter 'U', the coin will be on the cell(i-1,j) in time t+1, if i > 1. Otherwise there is no coin on your board at time t+1.

L: If at time t, the coin is on the cell(i,j) and cell(i,j) has letter 'L', the coin will be on the cell(i,j-1) in time t+1, if j > 1. Otherwise there is no coin on your board at time t+1.

D: If at time t, the coin is on the cell(i,j) and cell(i,j) has letter 'D', the coin will be on the cell(i+1,j) in time t+1, if i < N. Otherwise there is no coin on your board at time t+1.

R: If at time t, the coin is on the cell(i,j) and cell(i,j) has letter 'R', the coin will be on the cell(i,j+1) in time t+1, if j < M. Otherwise there is no coin on your board at time t+1.

When the coin reaches a cell that has letter '*', it will stay there permanently. When you punch on your board, your timer starts and the coin moves between cells. Before starting the game, you can make operations to change the board, such that, you are sure that at time K the coin will reach the cell having letter '*'. In each operation you can select a cell with some letter other than '*' and change the letter to 'U', 'L', 'R' or 'D'. You need to carry out as few operations as possible in order to achieve your goal. Your task is to find the minimum number of operations.

Input: 
The first line of input contains three integers N and M and K. Next N lines contains M letters which describe your board.

Output: 
Print an integer which represents the minimum number of operations required to achieve your goal. If you cannot achieve your goal, print -1 .

Constraints 
N, M <= 51 
K <= 1000

Sample input :

2 2 3  
RD  
*L
Sample output :

0
Sample input :

2 2 1  
RD  
*L
Sample output :

1
Explanation :

In the first example, you don't have to change any letter; but in the second example, you should change the letter of cell (1,1) to 'D'.

Copyright (c) 2015 HackerRank.
All Rights Reserved
Suggest Edits


RRD
DLL
*UU
'''


n, m, k = map(int, raw_input().strip().split(" "))
ar = []
max_k = m*n

dp = []
for i in xrange(n):
	dp.append([])
	for j in xrange(m):
		dp[i].append([])
		dp[i][j] = [max_k]*(k+1)

dp[0][0][0] = 0

for x in xrange(n):
	line = raw_input()
	for l in range(len(line)):
		if line[l] == '*':
			star_x, star_y = x, l
	ar.append(line)


def calc(i, j, o):
	mini = max_k
	if i-1>=0:
		if dp[i-1][j][o-1] + (0 if ar[i-1][j] == 'D' else 1) < max_k:
			mini = min(mini, dp[i-1][j][o-1] + (0 if ar[i-1][j] == 'D' else 1))

	if i+1<n:
		if dp[i+1][j][o-1] + (0 if ar[i+1][j]=='U' else 1) < max_k:
			mini = min(mini, dp[i+1][j][o-1] + (0 if ar[i+1][j]=='U' else 1))

	if j-1>=0:
		if dp[i][j-1][o-1] + (0 if ar[i][j-1]=='R' else 1) < max_k:
			mini = min(mini, dp[i][j-1][o-1] + (0 if ar[i][j-1]=='R' else 1))

	if j+1<m:
		if dp[i][j+1][o-1] + (0 if ar[i][j+1]=='L' else 1) < max_k:
			mini = min(mini, dp[i][j+1][o-1] + (0 if ar[i][j+1]=='L' else 1))

	return mini

for ki in xrange(k+1):
	for a in xrange(n):
		for b in xrange(m):
			if ki==0 and a==0 and b==0:
				continue
			else:
				dp[a][b][ki] = calc(a, b, ki)

answer = max_k

print dp[2][0]

for ki in xrange(k+1):
	answer = min(answer, dp[star_x][star_y][ki])

if answer == max_k:
	print -1
else:
	print answer
				























