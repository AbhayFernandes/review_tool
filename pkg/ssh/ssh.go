package ssh

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func Sign(message, filepath string) string {
	// TODO: update this based on configuration
	privateKeyBytes, err := os.ReadFile(filepath)

	if err != nil {
		// TODO: Handle gracefully
		fmt.Println("An error occured reading the private key")
		panic(err)
	}

	privateKey, err := ssh.ParseRawPrivateKey(privateKeyBytes)

	if err != nil {
		// TODO: Handle gracefully
		fmt.Println("An error occured parsing the private key")
		panic(err)
	}

	hash := sha256.Sum256([]byte(message))

	signer, err := ssh.NewSignerFromKey(privateKey)
	signature, err := signer.Sign(rand.Reader, hash[:])

	blob := base64.StdEncoding.EncodeToString(signature.Blob)
	return signature.Format + " " + blob
}

func Verify(signature, message, publicKey string) bool {
	fmt.Print(string(publicKey))

	publicKeyAuthorizedBytes, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey))

	publicKeyParse, err := ssh.ParsePublicKey(publicKeyAuthorizedBytes.Marshal())
	if err != nil {
		panic(err)
	}

	messageBytes := []byte(message)

	signatureParts := strings.Fields(signature)

	blob, err := base64.StdEncoding.DecodeString(signatureParts[1])
	if err != nil {
		panic(err)
	}

	sig := &ssh.Signature{
		Format: signatureParts[0],
		Blob:   blob,
	}

	hash := sha256.Sum256(messageBytes)
	err = publicKeyParse.Verify(hash[:], sig)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetPublicKey(filepath string) string {
	publicKeyBytes, err := os.ReadFile(filepath)

	if err != nil {
		// TODO: Handle gracefully
		fmt.Println("An error occured reading the private key")
		panic(err)
	}

	return string(publicKeyBytes)
}
