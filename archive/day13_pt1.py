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
    resultx = layer % ((rnge_size *2) - 2)
    if count == None:
        pass
    elif count == 0:
        return(ind)
    elif len(list_guy) == 1:
        return(ind) 
    else:
        while count > 0:
            try:
                ind += 1
      #          print(list_guy)
                list_guy[ind]
       #         print(ind)  
                count -= 1
            except IndexError:
                ind = 0
                list_guy = list(reversed(list_guy))
        # return(list_guy[ind])
        return(resultx)
total_catch = []
for x,y in ofile().items():
    scanner = where_is_scanner(y,x)
    if scanner == 0:
        total_catch.append(x*y)
print(sum(total_catch))