import collections

n, m = [int(i) for i in raw_input().strip().split()]

a = [int(i) for i in raw_input().strip().split()]
b = [int(i) for i in raw_input().strip().split()]
c = [int(i) for i in raw_input().strip().split()]



def one():
    return 1

factors = collections.defaultdict(one)
for i in range(0, m):
    factors[b[i]] *= c[i] % 1000000007


for i, factor in factors.iteritems():
    #for j in range(1, n/i + 1):
    j = 1
    while i*j <= n:
        idx = j*i-1
        a[idx] *= factor
        a[idx] %= 1000000007;
        j += 1


print ' '.join(map(str, a))
