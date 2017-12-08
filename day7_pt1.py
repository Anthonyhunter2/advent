#!/usr/bin/python
from re import search
def ofile():
	bigger_stuff = {}
	with open("things","r") as f:
		f = f.readlines()
		for line in f:
			line = line.split()
			key = line[0]
#			weight = "".join(line[1]).strip("(").strip(")")
			if "->" in line:
				stuff = {}
				for each in line[line.index("->")+1:]:
					stuff[each.strip(",")] = {}
				value = stuff
			else:
				value = {}
			bigger_stuff[key] = value
		return(bigger_stuff)
def main(jabby):
	for k,v in jabby.items():
		if not v:
			del jabby[k]
	for k,v in jabby.items():
		for x in v:	
			if x in jabby:
				jabby[k][x] = jabby[x]
				del jabby[x]
	for i in jabby:
		print(i)
	return(jabby)
#ofile()			
main(ofile())
