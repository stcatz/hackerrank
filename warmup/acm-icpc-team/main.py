N, M = raw_input().split()

matrix = []
for x in xrange(int(N)):
	matrix.append(int(raw_input(),2))

max_topic = 0
max_team  = 0

for i in xrange(int(N)):
	for j in xrange(i, int(N)):
		if( bin(matrix[i] | matrix[j]).count('1') > max_topic ):
			max_team = 1
			max_topic =  bin(matrix[i] | matrix[j]).count('1')
		elif (  bin(matrix[i] | matrix[j]).count('1') == max_topic ):
			max_team += 1


print max_topic
print max_team