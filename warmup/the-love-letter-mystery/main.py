num = input()

def love_count(val):
	length = len(val)
	count = 0
	for x in range(length/2):
		count += abs(ord(val[x]) - ord(val[length-1-x]))
	return count

for i in range(num):
	x = raw_input()
	print love_count(x)