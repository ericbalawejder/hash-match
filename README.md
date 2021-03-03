### Hash Match
Compare the contents of two files based on their [hash](https://pkg.go.dev/crypto).

##### Usage:
```
$ go run main.go --help
```
The default hashing algorithm is [sha256](https://pkg.go.dev/crypto/sha256). You can explicitly specify an
algorithm with an optional third flag `--hash=` and provide a Go supported 
[hashing algorithm](https://pkg.go.dev/crypto#Hash) with the exception of the 
[blake2s](https://pkg.go.dev/golang.org/x/crypto/blake2s) and 
[blake2b](https://pkg.go.dev/golang.org/x/crypto/blake2b) algorithms.
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
`--string` flag allows for string hashing with an optional specified algorithm.
```
$ go run main.go --string="a string to hash"
$ go run main.go --string="a string to hash" --hash=sha512
```
Using the binary:
```
$ go build main.go 
$ ./main --<flag>=
```
