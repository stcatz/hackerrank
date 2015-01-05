n = int(raw_input())
ar = [int(i) for i in raw_input().strip().split()]

sar = sorted(ar)

swap = map(lambda x,y: x-y, sar, ar)

indexes = []
is_swap = True

#print swap

i, j = 0, len(swap)-1

while swap[i]==0 and i<=j:
	i += 1

if i == j:
	print "yes"
else:
	while j>=i and swap[j]==0:
		j-=1

	indexes.append(i+1)
	indexes.append(j+1)

	while swap[i]==-swap[j] and i < j:
		i += 1
		if i==j:
			break
		j -= 1
		if swap[i] !=0 or swap[j] != 0:
			is_swap = False

	if i==j and is_swap:
		print "yes"
		print "swap", indexes[0], indexes[1]
	elif i==j and  not is_swap:
		print "yes"
		print "reverse", indexes[0], indexes[1]
	else:
		print "no"

