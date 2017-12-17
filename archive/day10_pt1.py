#!/usr/bin/python

def ofile():
    with open("steps","r") as f:
        new_list = []
        f = f.readline()   
        for i in f.split(","):
            i = i.strip("\n")
            i = int(i)
            new_list.append(i)
        return(new_list)
def loopedlist(w_list, w_ind, j_ind, jabby):
    times = (j_ind - 1)
    times2 = j_ind
    start_list = []
    rev_list = []
    start_list.append(w_list[w_ind])
    next_one = w_ind
    next_two = w_ind
    while times > 0:
        try:
            next_one = next_one + 1
            start_list.append(w_list[next_one])
            times -= 1
        except IndexError:
            next_one = -1
    for i in reversed(start_list):
        rev_list.append(i)
    for x in range(len(rev_list)):
        try:
            w_list[next_two] = rev_list[x]
            next_two = next_two + 1
        except IndexError:
            next_two = 0
            w_list[next_two] = rev_list[x]
            next_two = next_two + 1
    return(w_list)

def where_am_i (lister, ind, length, step):
    count = length + step
    while count > 0:
        try:
            ind = ind + 1
            lister[ind]
            count -= 1
        except IndexError:
            ind = -1
    return(ind)
def main(jabby):
    working_list = []
    index_num_list = 0
    step = 0
    for each in range(0,256):
        working_list.append(each)
    for x in range(len(jabby)):
        length = int(jabby[x])
        loopedlist(working_list, index_num_list, length, jabby)
        index_num_list = where_am_i(working_list, index_num_list, length, step)
        step += 1
    #    print(working_list)
    #    print(index_num_list, "index")
    print(working_list)
    print(working_list[0] * working_list[1])
        
            
        
main(ofile())

