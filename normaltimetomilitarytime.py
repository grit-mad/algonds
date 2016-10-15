#!/bin/python

import sys


time = raw_input().strip()
if time[8]=="A":
    if int(time[0]+time[1]) == 12:
        print str(int(time[0]+time[1])-12)+"0"+time[2:-2]
    else:
        print time[:-2]
else:
    if int(time[0]+time[1]) == 12:
        print time[:-2]
    else:
        print str(int(time[0]+time[1])+12)+time[2:-2]
