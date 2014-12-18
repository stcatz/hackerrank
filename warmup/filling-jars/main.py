N, M = raw_input().split()

result = 0

for i in range(int(M)):
	a, b, k = raw_input().split()
	result += int(k) * ( int(b) - int(a) + 1 )