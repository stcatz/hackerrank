import math

t = int(raw_input())

for ti in xrange(t):
	n = int(raw_input())
	a = [int(i) for i in raw_input().strip().split()]

	result = 0
	key_map = {}

	for x in a:
		if not key_map.has_key(x):
			key_map[x] = 0
		else:
			square = int(math.sqrt(key_map[x]))
			key_map[x] += (square + 1) * 2
			result += (square + 1) * 2

	print result



'''
1        0      0       0
2        2      1       1          
3        6      3       1+2
4        12     6       1+2+3
5               10      1+2+3+4 
                        (1+i-1)(i-1)/2
'''


