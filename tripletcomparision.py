#!/bin/python

import sys

a=map(int, raw_input().split())
b=map(int, raw_input().split())
alice=0
bob=0

for i in range(len(a)):
    if a[i] > b[i]:
        alice+=1
    if b[i] > a[i]:
        bob+=1
print "%d %d" %(alice,bob)
        
