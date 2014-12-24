t = int(raw_input())

for i in range(t):
	n = int(raw_input())
	ar = [int(i) for i in raw_input().strip().split()]
	left = 0
	right = sum(ar) - ar[0]

	flag = False

	for index in range(n-1):
		left += ar[index]
		right -= ar[index+1]
		if  left == right:
			flag = True
			break
		elif left > right:
			break

	if len(ar)==1:
		flag=True

	if flag:
		print "YES"
	else:
		print "NO"
