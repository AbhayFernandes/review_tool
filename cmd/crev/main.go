package main

import (
	"fmt"
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
}
