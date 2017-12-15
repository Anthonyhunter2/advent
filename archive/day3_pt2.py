#!/usr/bin/python
import itertools


def move_right(x,y):
    return (x+1,y)
def move_left(x,y):
    return (x-1,y)
def move_down(x,y):
    return (x,y-1)
def move_up(x,y):
    return (x,y+1)

moves = [move_right, move_up, move_left, move_down]
def around((x,y), grid):

    right = (x+1,y)
    right_up = (x+1,y+1)
    right_down = (x+1,y-1)
    left = (x-1,y)
    left_up = (x-1,y+1)
    left_down = (x-1,y-1)
    up = (x,y+1)
    down = (x,y-1)
    possible_loc = [right,right_up,right_down,left,left_up,left_down,up,down]
    sum_list = []
    for z in grid:
       for d in possible_loc:
          if d == z:
             sum_list.append(grid[z])
    return(sum(sum_list))

def gen_points(end):
    _moves = itertools.cycle(moves)
    n = 1
    pos = 0,0
    times_to_move = 1
    call_back = {}
    while n < end:
#    while True:
        for _ in range(2):
            move = next(_moves)
            for _ in range(times_to_move):
                call_back[pos] = n
                pos = move(*pos)
                n = around(pos, call_back)
        times_to_move +=1
        for i in sorted(call_back.values()):
           if i > end:
              print(i)
              break
gen_points(277678)
