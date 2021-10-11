import json
import time
import collections
import os
import psutil

EmployeeTp = collections.namedtuple("employeetp", "name age")


class Employee():
    def __init__(self, name, age):
        self.name = name
        self.age = age

start = time.time()
with open("employees.json") as f:
    # employees = [Employee(d["Name"], d["Age"]) for d in content]
    employees = [EmployeeTp(d["Name"], d["Age"]) for d in json.load(f)]
    print(psutil.Process(os.getpid()).memory_info()[0] / 1024 ** 2)
    employees = None
    print(psutil.Process(os.getpid()).memory_info()[0] / 1024 ** 2)
    
print("Took: ", round(time.time() - start, 4), "sec")













