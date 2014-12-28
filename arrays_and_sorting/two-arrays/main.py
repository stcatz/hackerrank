t = int(raw_input())


for ti in xrange(t):
	n, k= [int(i) for i in raw_input().strip().split()]
	a = [int(i) for i in raw_input().strip().split()]
	b = [int(i) for i in raw_input().strip().split()]

	a.sort()
	b.sort()
	b.reverse()

	flag = True

	for i in xrange(n):
		if a[i] + b[i] < k:
			flag = False

	if flag:
		print "YES"
	else:
		print "NO"
