import os 

current_dir = os.getcwd()

dir_path = current_dir
res = []

for path in os.listdir(dir_path):
    if os.path.isfile(os.path.join(dir_path, path)):
        res.append(path)

print(res)        