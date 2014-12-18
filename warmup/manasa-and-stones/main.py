t = int(raw_input())

for x in range(t):
	n, a, b = [ int(raw_input()) for i in range(3)]
	result = []
	for i in range(n):
		calc = a*(n-1-i) + b*i
		if not result.count(calc):
			result.append(calc)
	result.sort()
	print " ".join([ str(x) for x in result])