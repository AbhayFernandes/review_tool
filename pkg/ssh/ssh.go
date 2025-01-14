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

func Sign(message, filepath string) (string, error) {
	// TODO: update this based on configuration
	privateKeyBytes, err := os.ReadFile(filepath)

	if err != nil {
		// TODO: Handle gracefully
        fmt.Println("An error occured reading the private key: ", err)
		return "", err
	}

	privateKey, err := ssh.ParseRawPrivateKey(privateKeyBytes)

	if err != nil {
		// TODO: Handle gracefully
		fmt.Println("An error occured parsing the private key")
		return "", err
	}

	hash := sha256.Sum256([]byte(message))

	signer, err := ssh.NewSignerFromKey(privateKey)
	signature, err := signer.Sign(rand.Reader, hash[:])

	blob := base64.StdEncoding.EncodeToString(signature.Blob)
	return signature.Format + " " + blob, nil
}

func Verify(signature, message, publicKey string) bool {
	fmt.Print(string(publicKey))

	publicKeyAuthorizedBytes, _, _, _, err := ssh.ParseAuthorizedKey([]byte(publicKey))

	publicKeyParse, err := ssh.ParsePublicKey(publicKeyAuthorizedBytes.Marshal())
	if err != nil {
		return false
	}

	messageBytes := []byte(message)

	signatureParts := strings.Fields(signature)

	if len(signatureParts) < 2 {
		return false
	}

	blob, err := base64.StdEncoding.DecodeString(signatureParts[1])
	if err != nil {
		return false
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

func GetPublicKey(filepath string) (string, error) {
	publicKeyBytes, err := os.ReadFile(filepath)

	if err != nil {
		// TODO: Handle gracefully
        fmt.Println("An error occured reading the private key: ", err)
		return "", err
	}

	return string(publicKeyBytes), nil
}
