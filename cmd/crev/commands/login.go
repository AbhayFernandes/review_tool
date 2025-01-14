package commands

import (
	"flag"
	"fmt"
	"os"

	"github.com/AbhayFernandes/review_tool/cmd/crev/config"
	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"github.com/AbhayFernandes/review_tool/pkg/ssh"
)

func Login(args []string) {
    var flagSet = flag.NewFlagSet("uploadDiff", flag.ExitOnError)

	serverAddr := flagSet.String(
		"server", "crev.abhayf.com:8080",
		"The server address in the form of host:port",
	)

    nonce := flagSet.String(
        "nonce", "",
        "The nonce to sign. You should get this from the website",
    )

    flagSet.Parse(args)

	client, ctx, conn, cancel := getClient(serverAddr)
	defer conn.Close()
	defer cancel()
    
    // TODO: Update this to read the private key from conifg
    signedNonce, err := ssh.Sign(*nonce, os.Getenv("HOME") + "/.ssh/id_ed25519")

    _, err = client.VerifySession(ctx, &proto.VerifySessionRequest{
        SignedNonce: signedNonce,
        User: config.GetConfig().User,
	})

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error verifying session: ", err.Error())
    }
}
