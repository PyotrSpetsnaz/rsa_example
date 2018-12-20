package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
)

func (cli *CLI) decryptKey(aesFile, rsaFile, outputFile string) {
	aesK := readFile(aesFile)
	rsaK := readPrivateKey(rsaFile)
	data, err := rsa.DecryptPKCS1v15(rand.Reader, rsaK, aesK)
	if err != nil {
		log.Panic(err)
	}
	ioutil.WriteFile(outputFile, data, 0644)

	fmt.Printf("AES key was successfully decrypted and written into %s", outputFile)

}