package commands

import (
	"flag"
	"fmt"

	"github.com/AbhayFernandes/review_tool/cmd/crev/config"
)

func Signup(args []string) {
    flagSet := flag.NewFlagSet("signup", flag.ExitOnError)

	username := flagSet.String(
		"username", "None",
		"The username you want to sign up with",
	)

    flagSet.Parse(args)

    if (*username == "None") {
        fmt.Println("You must provide a username")
        return
    }

    configVal := config.GetConfig()

    // save the config
    configVal.User = *username

    config.SaveConfig(configVal)
}
