import itertools

n, m = [int(i) for i in raw_input().strip().split()]

a = [int(i) for i in raw_input().strip().split()]
b = [int(i) for i in raw_input().strip().split()]
c = [int(i) for i in raw_input().strip().split()]

for i in xrange(1,n+1):
	selectors = [i%x==0 for x in b]
	c1 = list(itertools.compress(c, selectors))

	a[i-1] *= reduce(lambda x,y :x*y,c1)

for i in a:
	print i % (10**9 + 7),