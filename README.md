### Hash Match
hash_diff.go checks the equality of two file's contents by comparing their [sha256](https://pkg.go.dev/crypto/sha256) hash.

##### Usage:
```
$ go run hash_diff.go --left=/PATH/TO/FILE --right=/PATH/TO/FILE
```
Using the binary:
```
$ go build hash_diff.go 
$ ./hash_diff --left=/PATH/TO/FILE --right=/PATH/TO/FILE
```

### Hash a single file
hash.go provides the [sha256](https://pkg.go.dev/crypto/sha256) hash of a single file's content.

##### Usage:
```
$ go run hash.go --file=/PATH/TO/FILE
```
Using the binary:
```
$ go build hash.go 
$ ./hash --file=/PATH/TO/FILE
```
