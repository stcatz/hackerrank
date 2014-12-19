def insertion_sort(l):
    count = 0
    for i in xrange(1, len(l)):
        j = i-1 
        key = l[i]
        while (l[j] > key) and (j >= 0):
           count += 1
           l[j+1] = l[j]
           j -= 1
        l[j+1] = key
    print count


m = input()
ar = [int(i) for i in raw_input().strip().split()]
insertion_sort(ar)