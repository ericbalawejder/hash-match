### Hash Match
Produce the hash of a file's contents or check the equality of two file's contents by 
comparing their [sha256](https://pkg.go.dev/crypto/sha256) hash.

##### Usage:
```
$ go run main.go --file1=/PATH/TO/FILE
```
or for comparison:
```
$ go run main.go --file1=/PATH/TO/FILE --file2=/PATH/TO/FILE
```
Using the binary:
```
$ go build main.go 
$ ./main --file1=/PATH/TO/FILE
$ ./main --file1=/PATH/TO/FILE --file2=/PATH/TO/FILE
```
