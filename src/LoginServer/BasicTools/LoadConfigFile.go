package BasicTools

import (
	"github.com/BurntSushi/toml"
	"github.com/RotaicX/RotaicXBasicTool"
)

type TomlConfig struct {
	DataBase dataBase `toml:"DataBase"`
}

type dataBase struct {
	Host         string
	Port         string
	UserName     string
	Password     string
	DataBaseName string
}

func LoadConfigFile(Path string) TomlConfig {
	var config TomlConfig
	if _, err := toml.DecodeFile(Path, &config); err != nil {
		RotaicxBasicTool.Rlog.Errorln(err.Error())
		panic(err)
	}
	return config
}
