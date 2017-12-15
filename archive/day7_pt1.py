#!/usr/bin/python
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
	while len(jabby) > 1:
		new_dict = {}
		for key in jabby:
			for value in jabby[key]:
				if value in jabby:
					new_dict[key] = {}
					new_dict[key][value] = jabby[value]
		jabby = new_dict
	for i in jabby:
		print(i)
main(ofile())
