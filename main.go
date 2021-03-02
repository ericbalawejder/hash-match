package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ericbalawejder/hash-match/hashing"
)

// Quotes must be used with white space, i.e. --string="two words"
var sequence = flag.String("string", "", "a string to hash")
var file1 = flag.String("file1", "", "a file to hash")
var file2 = flag.String("file2", "", "a file to hash")
var hashAlgorithm = flag.String("hash", "sha256", "a hashing function")

func main()  {
	flag.Parse()

	// Call hashing.Supported to validate if the "--hash=" argument is valid.
	if !hashing.Supported(*hashAlgorithm) {
		log.Fatalf("%s algorithm not supported", *hashAlgorithm)
	}

	// Operate on --string flag.
	if *sequence != "" {
		hash, err := hashing.StringDigest(*sequence, *hashAlgorithm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
		return
	}

	if *file1 == "" && *file2 == ""{
		log.Fatal("--file argument required")
	}

	// Operate on file2.
	if *file1 == "" {
		hash, err := hashing.ComputeHash(*file2, *hashAlgorithm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
		return
	}

	// Operate on file1.
	if *file2 == "" {
		hash, err := hashing.ComputeHash(*file1, *hashAlgorithm)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hash)
		return
	}

	// Compare the two file hashes.
	hash1, err1 := hashing.ComputeHash(*file1, *hashAlgorithm)
	hash2, err2 := hashing.ComputeHash(*file2, *hashAlgorithm)

	// Check for errors.
	if err1 != nil && err2 != nil {
		log.Fatal("error")
	}

	if hash1 != hash2 {
		log.Fatal("file contents are NOT equal :-(")
	}
	fmt.Println("file contents are equal :-)")
}
