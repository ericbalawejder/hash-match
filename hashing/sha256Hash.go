package hashing

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func Sha256Hash(filePath string) (string, error) {
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
	hash := sha256.New()

	// Copy the file in the hashing interface and check for any error.
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	// Get the 32 byte hashing. Define [:16] for 16 bytes.
	hashInBytes := hash.Sum(nil)//[:16]

	// Convert the bytes to a string
	sha256Hash := hex.EncodeToString(hashInBytes)
	return sha256Hash, nil
}
