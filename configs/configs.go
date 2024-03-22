package configs

import (
	"encoding/json"
	"os"
)

type Configs struct {
	Port           string `json:"port`
	DataSource     string `json:"data_source"`
	FileLcs        string `json:"file_lcs"`
	Url            string `json:"url,omitempty"`
	Token          string `json:"token,omitempty"`
	TokenExpired   string `json:"token_expired,omitempty"`
	RefreshExpired int    `json:"refresh_expires,omitempty`
	ReferenceToken string `json:"reference,omitempty"`
}

var config *Configs

func Get() *Configs {
	return config
}
func LoadConfig(path string) {
	configFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()

	byteValue, err := os.ReadFile(configFile.Name())
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}
}
