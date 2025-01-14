package config

import (
	"os"
	"runtime"
)

func getConfigDir() (string) {
    os := runtime.GOOS

    switch os {
        case "linux": 
            return getConfigDirLinux()
        case "windows":
            return getConfigDirWindows()
        case "darwin": 
            return getConfigDirDarwin()
    }

    return ""
}

func getConfigDirLinux() string {
    xdgConfigDir := os.Getenv("XDG_CONFIG_HOME")
    if xdgConfigDir == "" {
        xdgConfigDir = os.Getenv("HOME") + "/.config"
    }

    dir := xdgConfigDir + "/crev"

    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.Mkdir(dir, 0755)
    }

    return dir
}

func getConfigDirDarwin() string {
    xdgConfigDir := os.Getenv("XDG_CONFIG_HOME")
    if xdgConfigDir == "" {
        xdgConfigDir = os.Getenv("HOME") + "/Library/Application Support"
    }

    dir := xdgConfigDir + "/crev"

    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.Mkdir(dir, 0755)
    }

    return dir
}

func getConfigDirWindows() string {
    appData := os.Getenv("APPDATA")

    dir := appData + "\\crev"
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.Mkdir(dir, 0755)
    }

    return dir
}
