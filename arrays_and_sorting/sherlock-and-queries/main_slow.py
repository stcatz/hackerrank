n, m = [int(i) for i in raw_input().strip().split()]

a = [int(i) for i in raw_input().strip().split()]
b = [int(i) for i in raw_input().strip().split()]
c = [int(i) for i in raw_input().strip().split()]

for i in xrange(1,m+1):
	for j in xrange(1,n+1):
		if j%b[i-1]==0 :
			a[j-1] = a[j-1] * c[i-1]

for i in a:
	print i % (10**9 + 7),