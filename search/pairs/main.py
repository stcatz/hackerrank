'''
Problem Statement

Given N integers, count the number of pairs of integers whose difference is K.

Input Format 
The 1st line contains N & K (integers). 
The 2nd line contains N numbers of the set. All the N numbers are distinct.

Output Format 
An integer that tells the number of pairs of integers whose difference is K.

Constraints: 
Each integer will be greater than 0 and at least K smaller than 231-1

Sample Input #00:

5 2  
1 5 3 4 2  
Sample Output #00:

3
Sample Input #01:

10 1  
363374326 364147530 61825163 1073065718 1281246024 1399469912 428047635 491595254 879792181 1069262793 
Sample Output #01:

0
Copyright (c) 2015 HackerRank.
All Rights Reserved
Suggest Edits
'''
n, k = map(int, raw_input().strip().split(" "))
ar = map(int, raw_input().strip().split(" "))

count = 0
flag = {}

for i in xrange(n):
	flag[ar[i]] = 0

for i in xrange(n):
	if flag.has_key(ar[i]-k):
		count += 1
	if flag.has_key(ar[i]+k):
		count += 1 
		
print count/2