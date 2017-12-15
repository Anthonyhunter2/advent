#!/usr/bin/python
 
with open("course", "r") as f:
    f = f.readlines()
    store = []
    try_dict = {}
    count = 0
    for line in f:
        store.append(int(line.strip("\n")))
    for each in range(len(store)):
        try_dict[each] = store[each]
count = 0 
for key, value in try_dict.items():
	count = 0
	while True:
		try:
#			print(try_dict)
			new_key = None
			new_key = key + value
			if new_key in try_dict:
				if value >= 3:
					try_dict[key] = value - 1
					count += 1
				elif value < 3:
					try_dict[key] = value + 1
					count += 1
			key = new_key
			value = try_dict[key]
		except KeyError:
			print(count + 1, "count")
			break
	break
	break

