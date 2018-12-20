package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func (cli *CLI) decryptFile(encodedFile, outputFile, keyFile string) {
	key := readFile(keyFile)
	data := readFile(encodedFile)

	decrypted, err := decryptAES(key, string(data))
	if err != nil {
		log.Panic(err)
	}

	b := []byte(decrypted)
	b[0] = b[0]
	err = ioutil.WriteFile(outputFile, []byte(decrypted), 0644)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s was decrypted and written into %s", encodedFile, outputFile)
}
