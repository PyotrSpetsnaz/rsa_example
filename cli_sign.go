package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func (cli *CLI) sign(input, output, keyFile string) {
	privateKey := readPrivateKey(keyFile)
	message := readFile(input)

	hash := sha256.New()
	hash.Write(message)
	hashed := hash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		log.Panic(err)
	}
	hexSign := fmt.Sprintf("%X\n", signature)

	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashed, signature)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(output, append(message, []byte(hexSign)...), 0644)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Signature was successfully verified and written in %s\n", output)
}
