package main

import (
	"flag"
	"fmt"
)

func main() {
    serverAddr := flag.String(
        "server", "localhost:8080",
        "The server address in the form of host:port",
    )
    flag.Parse()

    client, ctx := getClient(serverAddr)
    fmt.Println(sayHello(client, ctx))

    currentDir := getCurrentDir()
    repository := getRepository(currentDir)
    diffs, err := getPatchDiffs(repository)
    if err != nil {
        panic(err)
    }
    fmt.Println(diffs)
}

