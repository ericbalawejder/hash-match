package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ericbalawejder/hash-match/hashing"
)

var file1 = flag.String("file1", "", "a file to hash")
var file2 = flag.String("file2", "", "a file to hash")

func main()  {
	flag.Parse()
	if *file1 == "" && *file2 == ""{
		log.Fatal("--file argument required")
	}

	// Operate on file2.
	if *file1 == "" {
		hash, err := hashing.Sha256Hash(*file2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
		return
	}

	// Operate on file1.
	if *file2 == "" {
		hash, err := hashing.Sha256Hash(*file1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
		return
	}

	// Compare the two file hashes.
	hash1, err1 := hashing.Sha256Hash(*file1)
	hash2, err2 := hashing.Sha256Hash(*file2)

	// Check for errors.
	if err1 != nil && err2 != nil {
		log.Fatal("error")
	}

	if hash1 != hash2 {
		log.Fatal("file contents are NOT equal :-(")
	}
	fmt.Println("file contents are equal :-)")
}
