n, m = [int(i) for i in raw_input().strip().split()]

ar = [0] * n

ops = []

for x in xrange(m):
	a,b,k = [int(i) for i in raw_input().strip().split()]
	ops.append([a, k])
	ops.append([b, -k])

#This has to be a stable sort. 
ops.sort(key=lambda i:(i[0],-i[1]))


ans, result = 0, 0
for x in ops:
	result += x[1]
	if result > ans:
		ans = result

print ans

