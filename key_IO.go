package main

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
)

func readFile(file string) []byte {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fInfo, err := f.Stat()
	size := fInfo.Size()

	bytes := make([]byte, size)
	buf := bufio.NewReader(f)
	_, err = buf.Read(bytes)
	if err != nil {
		log.Panic(err)
	}

	return bytes
}

func writePublicKey(filename string, key *rsa.PublicKey) {
	pemFile, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer pemFile.Close()

	var block = &pem.Block{
		Type:	"RSA PUBLIC KEY",
		Bytes:	x509.MarshalPKCS1PublicKey(key),
	}

	err = pem.Encode(pemFile, block)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Public key was written in %s\n", filename)
}

func writePrivateKey(filename string, key *rsa.PrivateKey) {
	pemFile, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer pemFile.Close()

	var block = &pem.Block{
		Type:	"RSA PRIVATE KEY",
		Bytes:	x509.MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(pemFile, block)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Private key was written in %s\n", filename)
}

func readPublicKey(filename string) *rsa.PublicKey {
	keyBytes := readFile(filename)
	pemBlock, _ := pem.Decode(keyBytes)
	key, err := x509.ParsePKCS1PublicKey(pemBlock.Bytes)
	if err != nil {
		log.Panic(err)
	}
	return key
}

func readPrivateKey(filename string) *rsa.PrivateKey {
	keyBytes := readFile(filename)
	pemBlock, _ := pem.Decode(keyBytes)
	key, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		log.Panic(err)
	}
	return key
}