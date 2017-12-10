#!/usr/bin/python

def ofile():
	var_track = []
	full = []
	with open("new_stuff", "r") as f:
		f = f.readlines()
		for line in f:
			line = line.split()
			full.append(line)
		return(full)
def main(guy):
	vals_dict = {}
	count = 1
	tracking = {}
	for line in guy:
		vals_dict[line[0]] = 0
	for line in guy:
	#	check_it(line, vals_dict)
		var2 = line[4]
		var1 = line[0]
		if var2 in vals_dict:
			checker = "{} {} {}".format(vals_dict[var2], line[5], int(line[6]))
			if eval(checker):
				if line[1] == "inc":
					vals_dict[var1] = vals_dict[var1] + int(line[2])
					count += 1
				elif line[1] == "dec":
					vals_dict[var1] = vals_dict[var1] - int(line[2])
					count += 1
			else:
				pass
		tracking[count] = max(vals_dict.values())
	print(max(vals_dict.values()))
	print(max(tracking.values()))
main(ofile())
