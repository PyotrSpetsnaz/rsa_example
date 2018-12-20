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
	fmt.Println("	createkeys [-private <PRIVATE>] [-public <PUBLIC>] - Create a public/private key pair and write them to files PRIVATE and PUBLIC")
	fmt.Println("	sign -file <FILE> -key <KEY> [-output <OUTPUT>] - Sign file FILE with key KEY and write it into file OUTPUT")
	fmt.Println("	encryptkey -aes <AES> -rsa <RSA> [-output <OUTPUT>] - Encrypt AES key with RSA and write it into file OUTPUT")
	fmt.Println("	decryptkey -aes <AES> -rsa <RSA> [-output <OUTPUT>] - Decrypt AES key with RSA and write it into file OUTPUT")
	fmt.Println("	encrypt -file <FILE> -aes <AES> [-output <OUTPUT> [-fmt <FORMAT>]] - Encrypt FILE with AES key AES and write it to file OUTPUT.FORMAT")
	fmt.Println("	decrypt -file <FILE> -aes <AES> [-output <OUTPUT> [-fmt <FORMAT>]] - Decrypt FILE with AES key AES, and write it to file OUTPUT.FORMAT")

}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()

	// TODO: write better usage for args
	generateKeysCmd := flag.NewFlagSet("createkeys", flag.ExitOnError)
	privateKey := generateKeysCmd.String("private", "private_key.pem", "The file in which the private key will be written")
	publicKey := generateKeysCmd.String("public", "public_key.pem", "The file in which the public key will be written")

	signCmd := flag.NewFlagSet("sign", flag.ExitOnError)
	fileToBeSigned := signCmd.String("file", "", "The file you want to sign")
	keyForSign := signCmd.String("key", "", "Encryption key")
	signedOutput := signCmd.String("output", "signed.txt", "Signed file")

	encryptKeyCmd := flag.NewFlagSet("encryptkey", flag.ExitOnError)
	aesToEncrypt := encryptKeyCmd.String("aes", "", "AES key that needs to be encrypted")
	rsaToEncrypt := encryptKeyCmd.String("rsa", "", "RSA key for encryption")
	aesEncryptedOutput := encryptKeyCmd.String("output", "encryptedAES.txt", "The file in which the encrypted AES key will be written")

	decryptKeyCmd := flag.NewFlagSet("decryptkey", flag.ExitOnError)
	aesToDecrypt := decryptKeyCmd.String("aes", "", "Encrypted AES key that need to be decrypted")
	rsaToDecrypt := decryptKeyCmd.String("rsa", "", "RSA key for decryption")
	aesDecryptedOutput := decryptKeyCmd.String("output", "decryptedAES.txt", "The file in which the decrypted AES key will be written")

	encryptFileCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	fileToEncrypt := encryptFileCmd.String("file", "", "File you want to encrypt")
	encryptionKey := encryptFileCmd.String("aes", "", "Other party aes key")
	encryptedFileOutput := encryptFileCmd.String("output", "encrypted", "The file in which encrypted information will be written.")
	encryptedFileOutputFmt := encryptFileCmd.String("fmt", "txt", "Format of the output file")

	decryptFileCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	encodedFile := decryptFileCmd.String("file", "", "File with encoded data")
	decryptionKey := decryptFileCmd.String("aes", "", "Key for decryption")
	decodedFileOutput := decryptFileCmd.String("output", "decrypted", "The file in which decrypted information will be written.")
	decodedFileOutputFmt := decryptFileCmd.String("fmt", "txt", "Format of the output file")

	switch os.Args[1] {
	case "createkeys":
		err := generateKeysCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "sign":
		err := signCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "encryptkey":
		err := encryptKeyCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "decryptkey" :
		err := decryptKeyCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
	}
	case "encrypt":
		err := encryptFileCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "decrypt":
		err := decryptFileCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if generateKeysCmd.Parsed() {
		if *privateKey == "" || *publicKey == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.generateKeyPair(*privateKey, *publicKey)
	}

	if signCmd.Parsed() {
		if *fileToBeSigned == "" || *signedOutput == "" || *keyForSign == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.sign(*fileToBeSigned, *signedOutput, *keyForSign)
	}

	if encryptKeyCmd.Parsed() {
		if *aesToEncrypt == "" || *rsaToEncrypt == "" || *aesEncryptedOutput == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.encryptKey(*aesToEncrypt, *rsaToEncrypt, *aesEncryptedOutput)
	}

	if decryptKeyCmd.Parsed() {
		if *aesToDecrypt == "" || *rsaToDecrypt == "" || *aesDecryptedOutput == "" {
			cli.printUsage()
			os.Exit(1)
		}
		cli.decryptKey(*aesToDecrypt, *rsaToDecrypt, *aesDecryptedOutput)
	}

	if encryptFileCmd.Parsed() {
		if *fileToEncrypt == "" || *encryptionKey == "" || *encryptedFileOutput == "" || *encryptedFileOutputFmt == "" {
			cli.printUsage()
			os.Exit(1)
		}
		outputFile := fmt.Sprintf("%s.%s", *encryptedFileOutput, *encryptedFileOutputFmt)
		cli.encryptFile(*fileToEncrypt, outputFile,  *encryptionKey)
	}

	if decryptFileCmd.Parsed() {
		if *encodedFile == "" || *decryptionKey == "" || *decodedFileOutput == "" || *decodedFileOutputFmt == "" {
			cli.printUsage()
			os.Exit(1)
		}
		outputFile := fmt.Sprintf("%s.%s", *decodedFileOutput, *decodedFileOutputFmt)
		cli.decryptFile(*encodedFile, outputFile, *decryptionKey)
	}
}
