#!/usr/bin/python
import sys
def ofile():
    starter = {}
    with open(sys.argv[1]) as f:
        f = f.readlines()
        for line in f:
            line = line.split(":")
            starter[int(line[0])] = int(line[1].strip("\n"))
    return(starter)
def where_is_scanner(rnge_size, layer):
    ind = 0
    count = layer
    list_guy = range(rnge_size)
    resultx = layer % ((rnge_size * 2) -2)
    return(resultx)
z = 3000000
test_case = ofile()
while True:
    total_catch = []
    for x,y in test_case.items():
        scanner = where_is_scanner(y,(x + z))
        if scanner == 0:
            total_catch.append(1)
            break
    if sum(total_catch) == 0:
        print("I WON on Try: {}".format(z))
        break
    z += 1