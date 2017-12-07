#!/usr/bin/python
def file_guy():
	with open("step", "r") as f:
		store = []
		for line in f.readlines():
			line = line.split("\t")
			for each in line:
				store.append(int(each))
		return(store)
def loopedlist(stuff,times, ind):
	stuff[ind] = 0
	while times > 0:
		try:
			ind = ind + 1
			stuff[ind] = stuff[ind] +1
			times -= 1
		except IndexError:
			ind = -1 
	return(stuff)
list1 = file_guy()
#list2 = loopedlist(file_guy(), 14,5)
list2 = file_guy()
firstpoint = list2.index(max(list2)) 
def main(guy, test, sp):
	count = 0
	startingpoint = None
	new_guy = []
	done_al = []
	#while compare != test:# and startingpoint != sp:
	while count < 500000:
		startingpoint = guy.index(max(guy))
		distrib = guy[guy.index(max(guy))]	
		new_guy = loopedlist(guy, distrib, startingpoint)
		count += 1	
		if "".join(str(x) for x in new_guy) not in done_al:
			done_al.append("".join(str(x) for x in new_guy))
		else:
			break
	return(count)
print(main(list1, list2, firstpoint))

