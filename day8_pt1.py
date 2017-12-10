#!/usr/bin/python

def ofile():
	var_track = []
	full = []
	with open("new_stuff", "r") as f:
		f = f.readlines()
		for line in f:
			line = line.split()
			full.append(line)
#			var = line[0]
#			if_state = " ".join(line[3:]) + ":"
#			step = line[2]
#			if line[1] == "inc":
#				action = "+ {}".format(step)
#			elif line[1] == "dec":
#				action = "- {}".format(step)
#			else:
#				print("Didnt understand what \"{}\" was".format(line[1]))
#			if_state = "{} {} {}".format(" ".join(line[3:]) + ":", var, action) 
#			add = var, if_state
#			full.append(add)
		return(full)
def check_it(check, stuff):
	print(check)
	var1 = check[0]
	var2 = check[4]
	if var2 in stuff:
		checker = "{} {} {}".format(stuff[var2], check[5], int(check[6]))
		if eval(checker):
			if check[1] == "inc":
				stuff[var1] = stuff[var1] + int(check[2])
			elif check[1] == "dec":
				stuff[var1] = stuff[var1] - int(check[2])
		else:
			pass
	return(stuff)
	
def main(guy):
	vals_dict = {}
	for line in guy:
		vals_dict[line[0]] = 0
	for line in guy:
		check_it(line, vals_dict)
	print(max(vals_dict.values()))
main(ofile())
