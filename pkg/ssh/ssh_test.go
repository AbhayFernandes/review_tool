package ssh_test

import (
    "testing"

	"github.com/AbhayFernandes/review_tool/pkg/ssh"
)


// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
    diffs := "Gladys"

    sig := ssh.Sign(diffs, "/home/abhay/.ssh/id_ed25519")

    publicKey := ssh.GetPublicKey("/home/abhay/.ssh/id_ed25519.pub")
    if (ssh.Verify(sig, diffs, publicKey)) != true {
        t.Fatalf("The ssh code did not correctly verify ssh sigs it should have.")
    }
}

