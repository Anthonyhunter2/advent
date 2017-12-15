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
def loopedlist(w_list, jabby, w_ind, j_ind):
#    times = jabby[j_ind]
    rev_list = []
    ind1 = w_list[w_ind]
    ind2 = jabby[j_ind]
    for a in reversed(w_list[w_ind:w_list[w_ind + jabby[j_ind]]]):
        rev_list.append(a)
    print(rev_list)
    for i in w_list[w_ind:w_list[w_ind + jabby[j_ind]]]:
        print(i)
 #   while times > 0:
  #      try:
   #         
    #        ind = ind + 1
    #        stuff[ind] = stuff[ind] +1
    #        times -= 1
    #    except IndexError:
    #        ind = -1
   # return(stuff)
    
def main(jabby):
    working_list = []
    list_for_changes = []
    index_num_step = 0
    index_num_list = 1
    step = 0
    for each in range(0,5):
        working_list.append(each)
    print(working_list, index_num_list, jabby, index_num_step)
    print(working_list[index_num_list], jabby[index_num_step])
    loopedlist(working_list, jabby, index_num_list, index_num_step)
#    print(reversed(working_list[index_num_list:jabby[index_num_step]]))
   # for a in reversed(working_list[index_num_list:jabby[index_num_step]]):
    #    print(a)
   # for a in working_list[index_num_list:jabby[index_num_step]+1]:
main(ofile())

