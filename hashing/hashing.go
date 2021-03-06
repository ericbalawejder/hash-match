package hashing

import (
	"encoding/hex"
	"io"
	"os"
)

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
