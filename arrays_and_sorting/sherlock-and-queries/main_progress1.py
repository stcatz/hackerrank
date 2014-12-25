import itertools

n, m = [int(i) for i in raw_input().strip().split()]

a = [int(i) for i in raw_input().strip().split()]
b = [int(i) for i in raw_input().strip().split()]
c = [int(i) for i in raw_input().strip().split()]


for i in range(0, m):
	c[i] = c[i] % 1000000007
	factor = b[i]

	j=1
	for index = factor*j; index <= n; j++:
	#while index <= n:
		a[index-1] *= c[i]
		a[index-1] %= 1000000007

for i in a:
	print i,
	#% (10**9 + 7)
