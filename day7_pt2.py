#!/usr/bin/python

def ofile ():
	bigger_stuff = {}
	with open("things","r") as f:
		f = f.readlines()
		for line in f:
			line = line.split()
			weight = "".join(line[1]).strip("(").strip(")")
			key = "{} - {}".format(line[0], weight)
			if "->" in line:
				stuff = {}
				for each in line[line.index("->")+1:]:
					stuff[each.strip(",")] = {}
				value = stuff
			else:
				value = {}
			bigger_stuff[key] = value
	return(bigger_stuff)

def num_dict(jabby):
	name_num = {}
	count = 0
	
	for k in jabby:
		name = k.split()[0]
		cot = int(k.split()[2])
		name_num[name] = cot
	while count < 5:
		check = []
		for key in jabby:
			divide = len(jabby[key])
			nme = key.split()[0]
			if divide != 0:
				total = 0
				vals = []
				check2 = []
				for value in jabby[key]:
						vals.append(int(name_num[value]))
				if len(set(vals)) <= 1:
					if name_num[nme] == int(key.split()[2]) + int(vals[0]) * divide:
						pass
					else:
						name_num[nme] = name_num[nme] + int(vals[0]) * divide 
				else: 
					print("\t {} {}".format(key, name_num[key.split()[0]]))
					check.append(key)
					for x in jabby[key]:
						print(x, name_num[x], count)
		count += 1
		print(len(check))
num_dict(ofile())
