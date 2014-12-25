from collections import deque

n, k, q = [int(i) for i in raw_input().strip().split()]
a = [int(i) for i in raw_input().strip().split()]

qs = []
for x in xrange(q):
	qs.append(int(raw_input()))

a = deque(a)

for x in xrange(k):
	a.rotate(1)

for x in xrange(q):
	print a[qs[x]]


