# Python vs Golang, Read JSON file and transform to namedtuple (Py) or struct (Go)

| Language         | Runtime[sec]  | Memory[MB] |
|------------------|---------------|------------|
|------------------|---------------|------------|
| Python           | 7.71          | 1304       |
| Golang           | **1.55**      | **184**    |

Approximately:
- Golang is 5 times faster (20% of Python Runtime)
- Golang uses 7 times less memory (14% of Python memory usage)

## Output run with python
```
PS C:\projects\github.com\PatrickVienne\golang-examples\parsers\json_reader> python .\parsejson.py
1304
15
('Took: ', 7.713, 'sec')
```

## Output run with golang
```
PS C:\projects\github.com\PatrickVienne\golang-examples\parsers\json_reader> go run .\parsejson.go .\employees.json
Parsin File: .\employees.json
Alloc = 184 MiB TotalAlloc = 257 MiB    Sys = 272 MiB   NumGC = 3
Alloc = 184 MiB TotalAlloc = 257 MiB    Sys = 272 MiB   NumGC = 3
Alloc = 0 MiB   TotalAlloc = 257 MiB    Sys = 272 MiB   NumGC = 4
emp Struct: main.employee{Name:"", Age:0}
2021/10/11 22:47:45 Parsin File: .\employees.json: 1.555656s
2021/10/11 22:47:45 Parsin File: .\employees.json: 1.6046545s
```

