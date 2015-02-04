n, m, k = map(int, raw_input().strip().split(" "))

bikers = []
bikes = []
distances = []

for x in xrange(n):
	bikers.append(map(int, raw_input().strip().split(" ")))

for x in xrange(m):
	bikes.append(map(int, raw_input().strip().split(" ")))


for i in xrange(n):
	for j in xrange(m):
		dis = (bikers[i][0] - bikes[j][0]) ** 2 + (bikers[i][1] - bikes[j][1]) ** 2
		distances.append([dis, i, j])

result = 0

distances.sort(key=lambda y: y[0])


print distances



print result
