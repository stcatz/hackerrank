n = int(raw_input())
a = [int(i) for i in raw_input().strip().split()]

pmean = sum(map(lambda x,y: x*y, a, list(range(1,n+1))))

max_result = pmean
sum_a = sum(a)

for i in range(n):
	#print pmean
	pmean = pmean - sum_a + (a[i] * (n))
	if pmean > max_result:
		max_result = pmean

print max_result
	

