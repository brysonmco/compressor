package main

import (
	"aidanwoods.dev/go-paseto"
	"log"
)

func main() {
	privateKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := privateKey.Public()

	log.Printf("Private Key: %s", privateKey.ExportHex())
	log.Printf("Public Key: %s", publicKey.ExportHex())
}
