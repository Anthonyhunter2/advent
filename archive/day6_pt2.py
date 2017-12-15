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
def main(guy, ):
	count = 0
	startingpoint = None
	new_guy = []
	done_al = []
	while count < 500000:
		startingpoint = guy.index(max(guy))
		distrib = guy[guy.index(max(guy))]	
		new_guy = loopedlist(guy, distrib, startingpoint)
		count += 1	
		guy_list = "".join(str(x) for x in new_guy)
		if guy_list not in done_al:
			done_al.append(guy_list)
		else:
			print(len(done_al) - done_al.index(guy_list))
			break
	return(count)
print(main(list1))

