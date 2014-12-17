string = raw_input()
 
found = False
# Write the code to find the required palindrome and then assign the variable 'found' a value of True or False

length = len(string)
if length%2 == 1:
	max_odd = 1
else:
	max_odd = 0

result = {}

for x in string:
	if result.has_key(x):
		result[x] += 1
	else:
		result[x] = 1

odd = 0
for i in result.keys():
	if result[i]%2==1:
		odd += 1

if odd > max_odd:
	found = False
else:
	found = True

    

if not found:
    print("NO")
else:
    print("YES")
