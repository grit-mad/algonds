#!/bin/python

import sys


n = int(raw_input().strip())
arr = map(int,raw_input().strip().split(' '))
positive_count=0
zero_count=0
negative_count=0
for i in arr:
    if i>0:
        positive_count+=1
    elif i==0:
        zero_count+=1
    else:
        negative_count+=1
print positive_count/(n*1.0)
print negative_count/(n*1.0)
print zero_count/(n*1.0)

