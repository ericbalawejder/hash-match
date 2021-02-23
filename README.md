### Hash Match
Produce the hash of a file's contents or check the equality of two file's contents by 
comparing their [hash](https://pkg.go.dev/crypto).

##### Usage:
```
$ go run main.go --help
```
The default hashing algorithm is [sha256](https://pkg.go.dev/crypto/sha256). You can explicitly specify an
algorithm with an optional third flag `--hash=` and provide a 
[Go supported hashing algorithm](https://pkg.go.dev/crypto#Hash).
```
$ go run main.go --file1=/PATH/TO/FILE
```
With optional specified hashing algorithm:
```
$ go run main.go --file1=/PATH/TO/FILE --hash=md5
```
or for equality comparison:
```
$ go run main.go --file1=/PATH/TO/FILE --file2=/PATH/TO/FILE --hash=md5
```
Using the binary:
```
$ go build main.go 
$ ./main --file1=/PATH/TO/FILE
$ ./main --file1=/PATH/TO/FILE --file2=/PATH/TO/FILE
$ ./main --file1=/PATH/TO/FILE --file2=/PATH/TO/FILE --hash=md5
```
