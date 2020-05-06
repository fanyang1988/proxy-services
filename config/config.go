package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config cfg for cmd
type Config struct {
	URL string `json:"url"`
}

// LoadCfg load cfg from file
func LoadCfg(path string) Config {
	jsonBytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	res := Config{}
	err = json.Unmarshal(jsonBytes, &res)
	if err != nil {
		panic(err)
	}

	return res
}
