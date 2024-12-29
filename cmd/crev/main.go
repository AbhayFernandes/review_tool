package main

import (
	// "flag"
	"fmt"

	"github.com/AbhayFernandes/review_tool/pkg/ssh"
)

func main() {
    // serverAddr := flag.String(
    //     "server", "localhost:8080",
    //     "The server address in the form of host:port",
    // )
    // flag.Parse()

    // client, ctx, conn, cancel := getClient(serverAddr)
    // defer conn.Close()
    // defer cancel()

    // fmt.Println(sayHello(client, ctx))

    currentDir := getCurrentDir()
    repository := getRepository(currentDir)

    diffs, err := getPatchDiffs(repository)
    if err != nil {
        panic(err)
    }
    fmt.Println(diffs)

    sig := ssh.Sign(diffs, "/home/abhay/.ssh/id_ed25519")
    fmt.Println(sig)

    publicKey := ssh.GetPublicKey("/home/abhay/.ssh/id_ed25519.pub")
    fmt.Print(ssh.Verify(sig, diffs, publicKey))
}

