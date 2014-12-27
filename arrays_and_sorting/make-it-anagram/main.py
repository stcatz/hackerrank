a = raw_input()
b = raw_input()

c = a+b
c = ''.join(set(list(c)))

result = 0

for x in c:
	n1, n2 = a.count(x), b.count(x)
	result += abs(n1-n2)

print result
	