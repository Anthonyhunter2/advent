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

def gen_points(end):
	_moves = itertools.cycle(moves)
	n = 1
	pos = 0,0
	times_to_move = 1
	call_back = {}
	while n < end:
		for _ in range(2):
			move = next(_moves)
			for _ in range(times_to_move):
				call_back[n] = pos
				pos = move(*pos)
				n += 1
		times_to_move +=1
	for i in call_back:
		if i == end:
			print(call_back[i])
gen_points(277678)
		
