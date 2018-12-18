package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func (cli *CLI) generateKeyPair(private, public string) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panic(err)
	}

	publicKey := &privateKey.PublicKey

	pemPrivateFile, err := os.Create(private)
	if err != nil {
		log.Panic(err)
	}
	defer pemPrivateFile.Close()

	pemOpenFile, err := os.Create(public)
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
	defer pemOpenFile.Close()

	var pemPrivateBlock = &pem.Block{
		Type: 	"RSA PRIVATE KEY",
		Bytes:	x509.MarshalPKCS1PrivateKey(privateKey),
	}

	var pemPublicBlock = &pem.Block{
		Type:	"RSA PUBLIC KEY",
		Bytes:	x509.MarshalPKCS1PublicKey(publicKey),
	}

	err = pem.Encode(pemPrivateFile, pemPrivateBlock)
	if err != nil {
		log.Panic(err)
	}

	err = pem.Encode(pemOpenFile, pemPublicBlock)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Private key was written in %s\nPublic key was written in %s\n", private, public)
}
