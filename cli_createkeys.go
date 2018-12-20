package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

func (cli *CLI) generateKeyPair(privateFile, publicFile string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Panic(err)
	}

	writePrivateKey(privateFile, privateKey)
	writePublicKey(publicFile, &privateKey.PublicKey)
}
