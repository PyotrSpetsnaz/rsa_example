package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI responsible fo processing command line arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreatekeys -private PRIVATE -public PUBLIC\t-\tCreate a public/private key pair and write them to files PRIVATE and PUBLIC")
	//fmt.Println("\tparsekey -file FILE\t-\tParse public key from FILE")
	fmt.Println("\tsign -file FILE -output OUTPUT -key KEY \t-\tSign file FILE with key KEY and write it into file OUTPUT")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	generateKeysCmd := flag.NewFlagSet("createkeys", flag.ExitOnError)
	signCmd := flag.NewFlagSet("sign", flag.ExitOnError)
	//parseKeyCmd := flag.NewFlagSet("parsekey", flag.ExitOnError)

	privateKey := generateKeysCmd.String("private", "private_key.pem", "The file in which the private key will be written")
	publicKey := generateKeysCmd.String("public", "public_key.pem", "The file in which the private key will be written.0")
	fileToBeSigned := signCmd.String("file", "message.txt", "The file you want to sign")
	signedOutput := signCmd.String("output", "signed.txt", "Signed file")
	keyForSign := signCmd.String("key", "private_key.pem", "Encryption key")
	//parseKeyPublic := parseKeyCmd.String("file", "public.pem", "Parse .pem file with public key")

	switch os.Args[1] {
	case "createkeys":
		err := generateKeysCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	//case "parsekey":
	//	err := parseKeyCmd.Parse(os.Args[2:])
	//	if err != nil {
	//		log.Panic(err)
	//	}
	case "sign":
		err := signCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if generateKeysCmd.Parsed() {
		cli.generateKeyPair(*privateKey, *publicKey)
	}

	if signCmd.Parsed() {
		cli.sign(*fileToBeSigned, *signedOutput, *keyForSign)
	}

	//if parseKeyCmd.Parsed() {
	//	cli.parseKey()
	//}
}
