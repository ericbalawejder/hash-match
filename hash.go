package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var file = flag.String("file", "", "a file to hash")

func main()  {
	flag.Parse()
	if *file == "" {
		log.Fatal("--file argument required")
	}

	_, err := os.Stat(*file)
	if os.IsNotExist(err) {
		log.Fatal("file does not exist.")
	}

	openFile, err := os.Open(*file)
	if err != nil {
		log.Fatal(err)
	}
	defer openFile.Close()

	hashStruct := sha256.New()
	if _, err := io.Copy(hashStruct, openFile); err != nil {
		log.Fatal(err)
	}
	fileHash := hashStruct.Sum(nil)
	fmt.Printf("%x\n", fileHash)
}
