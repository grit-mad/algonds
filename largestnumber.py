# Enter your code here. Read input from STDIN. Print output to STDOUT
length_of_aray=int(raw_input())
array= map(int, raw_input().split())
maxm=0
for item in range (0,length_of_aray-1):
    if array[item] >= array[item+1]:
        if array[item] > maxm:
            maxm=array[item]
    else:
        maxm2=array[item+1]
        if array[item+1] > maxm2:
            maxm2=array[item+1]
        
if (maxm2>=maxm):
    print maxm2
else:
    print maxm
