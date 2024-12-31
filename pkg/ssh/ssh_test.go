package ssh_test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"

	"github.com/AbhayFernandes/review_tool/pkg/ssh"
	gossh "golang.org/x/crypto/ssh"
)

func genereteSSHKeys(t *testing.T) (string, string) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	der := x509.MarshalPKCS1PrivateKey(privateKey)

	// Create a PEM block
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: der,
	}

	privateKeyPEM := pem.EncodeToMemory(block)

	sshPubKey, err := gossh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		t.Fatalf("Failed to generate public key from private key: %v", err)
	}

	pubKeyBytes := gossh.MarshalAuthorizedKey(sshPubKey)

	tmpFile, _ := os.CreateTemp("", "sshkey_test_*.pem")

	t.Cleanup(func() {
		os.Remove(tmpFile.Name())
	})

	if _, err := tmpFile.Write(privateKeyPEM); err != nil {
		t.Fatalf("failed to write private key PEM to file: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	pubFile, err := os.CreateTemp("", "sshkey_test_*.pub")
	if err != nil {
		t.Fatalf("failed to create temp file for public key: %v", err)
	}

	// Ensure removal after the test completes.
	t.Cleanup(func() {
		os.Remove(pubFile.Name())
	})

	if _, err := pubFile.Write(pubKeyBytes); err != nil {
		t.Fatalf("failed to write public key to file: %v", err)
	}
	if err := pubFile.Close(); err != nil {
		t.Fatalf("failed to close public key file: %v", err)
	}

	return tmpFile.Name(), pubFile.Name()
}

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	diffs := "Gladys"

	priv, pub := genereteSSHKeys(t)

	sig := ssh.Sign(diffs, priv)

	publicKey := ssh.GetPublicKey(pub)
	if (ssh.Verify(sig, diffs, publicKey)) != true {
		t.Fatalf("The ssh code did not correctly verify ssh sigs it should have.")
	}
}

func TestSign_EmptyMessage(t *testing.T) {
	diffs := ""

	priv, pub := genereteSSHKeys(t)
	sig := ssh.Sign(diffs, priv)

	publicKey := ssh.GetPublicKey(pub)
	if (ssh.Verify(sig, diffs, publicKey)) != true {
		t.Fatalf("The ssh code did not correctly verify ssh sigs it should have.")
	}
}

func TestSign_LongMessage(t *testing.T) {
	diffs := "ThisIsAVeryLongMessageThatExceedsNormalLength"

	priv, pub := genereteSSHKeys(t)
	sig := ssh.Sign(diffs, priv)

	publicKey := ssh.GetPublicKey(pub)
	if (ssh.Verify(sig, diffs, publicKey)) != true {
		t.Fatalf("The ssh code did not correctly verify ssh sigs it should have.")
	}
}

func TestVerify_InvalidSignature(t *testing.T) {
	diffs := "Gladys"
	invalidSig :=
		"ssh-ed25519 EzoRFp0RHKlcc9o7pfcmvOjNEuPGcIxjLdvOdXW0/0hgVjwngA/oEcZPsbDJA7527ZzhnoiFndqMbHd1jBMFDw=="

	_, pub := genereteSSHKeys(t)

	publicKey := ssh.GetPublicKey(pub)
	if (ssh.Verify(invalidSig, diffs, publicKey)) != false {
		t.Fatalf("The ssh code did not correctly identify an invalid signature.")
	}
}

func TestVerify_TwoSeperateSigsInvalid(t *testing.T) {
	diffs := "ThisIsAVeryLongMessageThatExceedsNormalLength"

	priv, _ := genereteSSHKeys(t)
	sig := ssh.Sign(diffs, priv)

	_, pub := genereteSSHKeys(t)

	publicKey := ssh.GetPublicKey(pub)
	if (ssh.Verify(sig, diffs, publicKey)) != false {
		t.Fatalf("The ssh code did not correctly verify ssh sigs were different.")
	}
}
