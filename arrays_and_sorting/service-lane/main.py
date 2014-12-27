n, t = [int(i) for i in raw_input().strip().split()]
width = [int(i) for i in raw_input().strip().split()]

for x in xrange(t):
	first, second = [int(i) for i in raw_input().strip().split()]
	print min(width[first:second+1])