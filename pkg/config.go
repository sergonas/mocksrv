package mocksrv

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	Proxy   = "PROXY"
	Respond = "RESPOND"
	Timeout = "TIMEOUT"
)

type ConfigRoot struct {
	Port       string `yaml:"port"`
	ConfigPort string `yaml:"configPort"`
	Roots      []Root `yaml:"roots"`
}

type Root struct {
	Name     string   `yaml:"name"`
	Method   string   `yaml:"method"`
	Path     string   `yaml:"path"`
	Response Response `yaml:"response"`
}

type Response struct {
	Code    int               `yaml:"code"`
	Body    string            `yaml:"body"`
	Headers map[string]string `yaml:"headers"`
}

func Parse(configPath string) ConfigRoot {
	configData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	root := ConfigRoot{}
	err = yaml.Unmarshal(configData, &root)
	if err != nil {
		log.Fatal(err)
	}

	return root
}
