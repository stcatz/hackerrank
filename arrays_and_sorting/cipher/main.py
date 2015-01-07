n, k = [int(i) for i in raw_input().strip().split()]

s = int(raw_input(), 2)
result = 0b0

while len(bin(result)) < 2*n+2:
	result = result ^ (s ^ (s << 1))
	s = s << k
print bin(result)[len(bin(result))-n:len(bin(result))]






'''
s = raw_input()

result = '0'*(k-1) + "0"*n + '0'*(k-1)

result = list(result)

for x in xrange(n):
	if s[x] == '1':
		if result[x:x+k].count("1") % 2 != 1:
			result[x+k-1] = '1'
	if s[x] == '0':
		if result[x:x+k].count("1") % 2 != 0:
			result[x+k-1] = '1'

print "".join(result[k-1:n+k-1])
'''
