package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/AbhayFernandes/review_tool/pkg/proto"
)

func main() {
    serverAddr := flag.String(
        "server", "crev.abhayf.com:8080",
        "The server address in the form of host:port",
    )

    user := flag.String(
        "user", "ferna355",
        "Your MSU NetID without the @msu.edu. Ex: ferna355",
    )

    flag.Parse()

    client, ctx, conn, cancel := getClient(serverAddr)
    defer conn.Close()
    defer cancel()

    diff, err := getDiffs()
    if (err != nil) {
        fmt.Println("There was an error getting diffs. Are you in a git repo? Is there a commit to upload?")
    }

    _, err = client.UploadDiff(ctx, &proto.UploadDiffRequest{
        Diff: diff,
        User: *user,
    }); if (err != nil) {
        fmt.Println("Uploading diffs has failed: ", err.Error())
    }
}

func getDiffs() (string, error) {
    currentDir := getCurrentDir()
    repository := getRepository(currentDir)

    diffs, err := getPatchDiffs(repository)
    if err != nil {
        panic(err)
    }
    if len(diffs) > 0 {
        return diffs, nil
    } else {
        return "", errors.New("diff doesn't exist. No commit exists")
    }
}
