package main

import (
	"flag"

	"github.com/AbhayFernandes/review_tool/cmd/crev/commands"
)

func main() {
    var (
        // define global flags here
        _ = flag.Bool("test", false, "test")
    )

    flag.Parse()
    args := flag.Args()

    rootCmd := ""

    if len(args) > 1 {
        rootCmd, args = args[0], args[1:]
    } else {
        args = []string{""}
    }

    switch rootCmd {
    case "help":
        commands.Help(args)
    case "login":
        commands.Login(args)
    case "signup":
        commands.Signup(args)
    default:
        commands.UploadCurrentDiff(args)
    }
}

