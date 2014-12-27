t = int(raw_input())

for ti in xrange(t):
	n = int(raw_input())
	ar = [int(i) for i in raw_input().strip().split()]

	result = 0

	for i in xrange(n):
		times = (i+1) * (n-i)
		if times % 2 ==1:
			result = result ^ ar[i]

	print result

