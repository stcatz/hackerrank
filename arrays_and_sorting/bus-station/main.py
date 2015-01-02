import math

n = int(raw_input())
a = [int(i) for i in raw_input().strip().split()]

sum_a = sum(a)
max_a = max(a)

result = []

for i in xrange(1, int(math.sqrt(sum_a))+1):
	if sum_a % i == 0:
		if i >= max_a:	
			result.append(i)
		if sum_a/i != i and sum_a/i >= max_a:
			result.append(sum_a/i)

result.sort()

for x in result:
	sum_b = 0
	flag = True
	for i in xrange(n):
		sum_b += a[i]
		if sum_b < x:
			continue
		elif sum_b == x:
			sum_b = 0
		elif sum_b > x:
			flag = False
			break
	if flag:
		print x,
