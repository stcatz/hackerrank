t = int(raw_input())

for i in range(t):
	n = int(raw_input())

	fou = False
	for x in range(n+1):
		if (n-x)%3==0 and x%5==0:
			fou = True
			print "5"*(n-x) + "3"*x
			break

	if not fou:
		print "-1"

