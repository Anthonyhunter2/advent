#!/usr/bin/python

def file_open():
	with open("chart", "r") as f:
		f = " ".join(f.readlines()).split("\n")
		return(f)
def main(stuff):
	total = 0
	for line in stuff:
		if line:
			line = map(int, line.split("\t"))
			biggest = max(line)
			smallest = min(line)
			mean = biggest - smallest
			total = total + mean
	print(total)
def main2(stuff):
	total = 0 
	for line in stuff:
		if line:
			line = map(int, line.split("\t"))
			for ind in range(len(line)):
				for num in range(ind + 1, len(line)):
					if line[ind] % line[num] == 0:
						total = total + (line[ind] / line[num])
					elif line[num] % line[ind] == 0:
						total = total + (line[num] / line[ind])
	print(total)
main2(file_open())
