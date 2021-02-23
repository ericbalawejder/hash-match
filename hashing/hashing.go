package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
	"hash"
	"io"
	"os"
)

// Supported hash algorithms. Blake2s and Blake2b return (hash, err)
var algorithms = map[string]func() hash.Hash{
	"md4":    func() hash.Hash { return md4.New() },
	"md5":    func() hash.Hash { return md5.New() },
	"sha1": func() hash.Hash { return sha1.New() },
	"sha224": func() hash.Hash { return sha256.New224() },
	"sha256": func() hash.Hash { return sha256.New() },
	"sha384": func() hash.Hash { return sha512.New384() },
	"sha512": func() hash.Hash { return sha512.New() },
	"ripemd160": func() hash.Hash { return ripemd160.New() },
	"sha3_224": func() hash.Hash { return sha3.New224() },
	"sha3_256": func() hash.Hash { return sha3.New256() },
	"sha3_384": func() hash.Hash { return sha3.New384() },
	"sha3_512": func() hash.Hash { return sha3.New512() },
	"sha512_224": func() hash.Hash { return sha512.New512_224() },
	"sha512_256": func() hash.Hash { return sha512.New512_256() },
}

func Supported(algorithm string) bool {
	return algorithms[algorithm] != nil
}

func ComputeHash(filePath string, algorithm string) (string, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return "", err
	}

	// Open the argument and check for any error.
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	// Close file when the current function returns.
	defer file.Close()

	// Open a new hashing interface to write to.
	h := algorithms[algorithm]()

	// Copy the file in the hashing interface and check for any error.
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	// Get the 32 byte hashing. Define [:16] for 16 bytes.
	hashInBytes := h.Sum(nil) //[:16]

	// Convert the bytes to a string
	s := hex.EncodeToString(hashInBytes)
	return s, nil
}
