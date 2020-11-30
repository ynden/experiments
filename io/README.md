# IO
### Context
The stream package contains (for now ) a function (Read) that retrieves a file as a chunk of bytes.  
The number of byte chunks depends on the nbParts variable passed to the function

### How to run
Inside cmd folder run:
```bash
$ go run main.go <file path>
```

Example:
```
$ go run main.go ~/Desktop/test.txt
```

You can also use 'test.txt' inside the samples folder:
```
$ go run main.go /Users/your_username/go/src/github.com/ynden/experiments/samples/test.txt
```