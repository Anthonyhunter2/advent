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
            new_key = None
            new_key = key + try_dict[key]
            if new_key in try_dict:
               try_dict[key] = value + 1
               count += 1
            key = new_key
            value = try_dict[key]
        except KeyError:
            print(count + 1, "count")
            break
    break
    break

