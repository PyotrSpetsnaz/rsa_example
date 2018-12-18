package main

import (
	"bufio"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
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

func (cli *CLI) sign(input, output, key string) {
	k := readFile(key)
	message := readFile(input)

	pemKey, _ := pem.Decode(k)

	privateKey, err := x509.ParsePKCS1PrivateKey(pemKey.Bytes)
	if err != nil {
		log.Panic(err)
	}

	hash := sha256.New()
	hash.Write(message)
	hashed := hash.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		log.Panic(err)
	}
	hexSign := fmt.Sprintf("%X\n", signature)

	err = ioutil.WriteFile(output, append(message, []byte(hexSign)...), 0644)
	if err != nil {
		log.Panic(err)
	}

	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, hashed, signature)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Signature was successfully verified and written in %s\n", output)
}
