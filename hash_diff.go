package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var left = flag.String("left", "", "a file to compare")
var right = flag.String("right", "", "a file to compare")

//$ ./hash_diff --left <filea> --right <fileb>
func main() {
	flag.Parse()
	if *left == "" || *right == "" {
		log.Fatal("--left and --right are required arguments")
	}
	// make sure left and right files exist
	_, err := os.Stat(*left)
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}
	_, err = os.Stat(*right)
	if os.IsNotExist(err) {
		log.Fatal("File does not exist.")
	}

	leftFile, err := os.Open(*left)
	if err != nil {
		log.Fatal(err)
	}
	defer leftFile.Close()

	leftHashStruct := sha256.New()
	if _, err := io.Copy(leftHashStruct, leftFile); err != nil {
		log.Fatal(err)
	}
	leftFileHash := leftHashStruct.Sum(nil)
	fmt.Printf("%x\n", leftFileHash)


	rightFile, err := os.Open(*right)
	if err != nil {
		log.Fatal(err)
	}
	defer rightFile.Close()

	rightHashStruct := sha256.New()
	if _, err := io.Copy(rightHashStruct, rightFile); err != nil {
		log.Fatal(err)
	}
	rightFileHash := rightHashStruct.Sum(nil)
	fmt.Printf("%x\n", rightFileHash)


	if string(leftFileHash) != string(rightFileHash) {
		log.Fatal("file contents are not equal")
	}
	fmt.Println("file content is equal :-)")
}
