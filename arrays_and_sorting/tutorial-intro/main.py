v = int(raw_input())
n = int(raw_input())

a = [int(x) for x in raw_input().split()]

for x in xrange(n):
	if a[x] == v:
		print x