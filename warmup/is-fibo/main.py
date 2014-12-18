num = input()

def isFibo(val):
	if val<4:
		return "IsFibo"
	else:
		first,second,result = 0,1,0
		flag = 0
		while(result < val):
			result = first + second
			first = second
			second = result
		if(result == val):
			flag = 1
		if(flag == 1):
			return "IsFibo"
		else:
			return "IsNotFibo"
		

for i in range(num):
	x = input()
	print isFibo(x)