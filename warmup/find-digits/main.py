num = input()

for i in range(num):
	x = raw_input()
	count = 0
	for i in x:
		if( int(i)!=0 and int(x) % int(i) == 0):
			count += 1
	print count
