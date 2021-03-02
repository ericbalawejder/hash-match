package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"hash"
)

// Supported hash algorithms. Blake2s and Blake2b return (hash, err)
var algorithms = map[string]func() hash.Hash{
	"md4":        func() hash.Hash { return md4.New() },
	"md5":        func() hash.Hash { return md5.New() },
	"sha1":       func() hash.Hash { return sha1.New() },
	"sha224":     func() hash.Hash { return sha256.New224() },
	"sha256":     func() hash.Hash { return sha256.New() },
	"sha384":     func() hash.Hash { return sha512.New384() },
	"sha512":     func() hash.Hash { return sha512.New() },
	"ripemd160":  func() hash.Hash { return ripemd160.New() },
	"sha3_224":   func() hash.Hash { return sha3.New224() },
	"sha3_256":   func() hash.Hash { return sha3.New256() },
	"sha3_384":   func() hash.Hash { return sha3.New384() },
	"sha3_512":   func() hash.Hash { return sha3.New512() },
	"sha512_224": func() hash.Hash { return sha512.New512_224() },
	"sha512_256": func() hash.Hash { return sha512.New512_256() },
}
