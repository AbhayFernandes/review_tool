package commands

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/AbhayFernandes/review_tool/cmd/crev/config"
	"github.com/AbhayFernandes/review_tool/pkg/proto"
	"github.com/AbhayFernandes/review_tool/pkg/ssh"
	"google.golang.org/grpc/metadata"
)

func UploadCurrentDiff(args []string) {
    var flagSet = flag.NewFlagSet("uploadDiff", flag.ExitOnError)

	serverAddr := flagSet.String(
		"server", "crev.abhayf.com:8080",
		"The server address in the form of host:port",
	)

	sshKey := flagSet.String(
		"ssh", os.Getenv("HOME") + "/.ssh/id_ed25519",
		"The filepath to your local ssh private key. Ex: ~/.ssh/id_ed25519",
	)

	flagSet.Parse(args)

    user := config.GetConfig().User

	client, ctx, conn, cancel := getClient(serverAddr)
	defer conn.Close()
	defer cancel()

	diff, err := getDiffs()
	if err != nil {
		fmt.Println("There was an error getting diffs. Are you in a git repo? Is there a commit to upload?")
	}

	res, err := ssh.Sign(diff, *sshKey)
	md := metadata.Pairs("ssh_sig", res)
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err = client.UploadDiff(ctx, &proto.UploadDiffRequest{
		Diff: diff,
		User: user,
	})
	if err != nil {
		fmt.Println("Uploading diffs has failed: ", err.Error())
	}
}

func getDiffs() (string, error) {
	currentDir, err := getCurrentDir()
	repository := getRepository(currentDir)

	diffs, err := getPatchDiffs(repository)
	if err != nil {
		return "", nil
	}
	if len(diffs) > 0 {
		return diffs, nil
	} else {
		return "", errors.New("diff doesn't exist. No commit exists")
	}
}
