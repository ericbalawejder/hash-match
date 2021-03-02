package hashing

import (
	"encoding/hex"
)

func StringDigest(sequence string, algorithm string) (string, error) {
	// Open a new hashing interface to write to.
	h:= algorithms[algorithm]()

	// Coerce the string to bytes.
	h.Write([]byte(sequence))

	// Get the 32 byte hashing. Define [:16] for 16 bytes.
	hashInBytes := h.Sum(nil)//[:16]

	// Convert the bytes to a string
	s := hex.EncodeToString(hashInBytes)
	return s, nil
}
