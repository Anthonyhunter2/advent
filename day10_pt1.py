#!/usr/bin/python

def ofile():
    with open("steps","r") as f:
        new_list = []
        f = f.readline()   
        for i in f.split():
            i = i.strip(",").strip("\n")
            i = int(i)
            new_list.append(i)
        return(new_list)
def loopedlist(w_list, w_ind, j_ind):
    print(j_ind - 1, "IM TXT")
    times = (j_ind + 1)
    times2 = j_ind
    start_list = []
    rev_list = []
    start_list.append(w_list[w_ind])
    next_one = w_ind
    next_two = w_ind
    print(times)
    while times > 0:
        try:
            next_one = next_one + 1
            start_list.append(w_list[next_one])
            times -= 1
        except IndexError:
            next_one = -1
    for i in reversed(start_list):
        rev_list.append(i)
    print(rev_list)
    for x in range(len(rev_list)):
        try:
            w_list[next_two] = rev_list[x]
            next_two = next_two + 1
        except IndexError:
            next_two = 0
            w_list[next_two] = rev_list[x]
            next_two = next_two + 1
    return(w_list)
def main(jabby):
    working_list = []
    list_for_changes = []
    index_num_list = 0
    step = 0
    cycles = len(jabby)
    for each in range(0,10):
        working_list.append(each)
   # while cycles > 0:
    for x in range(len(jabby)):
        length = int(jabby[x])
     #   cwp = index_num_list
   #     index_num_list = length + step
        print(index_num_list, working_list[index_num_list], length)
        loopedlist(working_list, index_num_list, length)
   #     else:
    #        index_num_list = index_num_list - working_list[cwp:]
        print(working_list, index_num_list, jabby, length)
        print(working_list[index_num_list], length)
   #     index_num_list += 1
    #    step += 1
main(ofile())

