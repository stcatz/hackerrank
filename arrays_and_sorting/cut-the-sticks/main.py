n = int(raw_input())

ar = [int(i) for i in raw_input().strip().split()]

while len(ar) != 0:
	cut = min(ar)
	print len(ar)
	ar = [i-cut for i in ar]
	ar = filter(lambda x: x!=0, ar)
	