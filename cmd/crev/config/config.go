package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
    User string
}

func GetConfig() Config {
    configPath := GetConfigPath()

    var conf Config
    _, err := toml.DecodeFile(configPath, &conf)

    if err != nil {
        fmt.Println("error reading config file")
    }

    return conf
}

func GetConfigPath() string {
    configPath := getConfigDir() + "/config.toml"

    if _, err := os.Stat(configPath); os.IsNotExist(err) {
        os.Create(configPath)
    }

    return configPath
}

func SaveConfig(conf Config) {
    configPath := GetConfigPath()
    f, err := os.OpenFile(configPath, os.O_RDWR, 0644)

    if err != nil {
        fmt.Println("error opening config: ", err)
    }

    tomlBytes, _ := toml.Marshal(conf)
    f.Write(tomlBytes)

    f.Close()
}
