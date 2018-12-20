package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func (cli *CLI) encryptFile(fileToEncrypt, outputFile, keyFile string) {
	key := readFile(keyFile)
	data := readFile(fileToEncrypt)

	encrypted, err := encryptAES(key, data)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(outputFile, []byte(encrypted), 0644)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s was encrypted and written into %s", fileToEncrypt, outputFile)

	//dec, _ := decryptAES(key, encrypted)
	//_ = ioutil.WriteFile("test.jpg", []byte(dec), 0644)
}

