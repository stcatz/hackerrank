# Enter your code here. Read input from STDIN. Print output to STDOUT
T = int(raw_input())
for i in range (0,T):
    A,B,C1 = [int(x) for x in raw_input().split(' ')]
    
    answer = 0
    # write code to compute answer
    count = A/B
    answer = count
    #answer = count + count/C1
    while count >= C1:
        answer += count/C1
        count = count/C1 + count % C1
    print answer