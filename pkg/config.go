package mocksrv

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

const (
	ProxyType   = "PROXY"
	RespondType = "RESPOND"
	TimeoutType = "TIMEOUT"
)

type ConfigRoot struct {
	Port       string  `yaml:"port"`
	ConfigPort string  `yaml:"configPort"`
	Proxies    []Proxy `yaml:"proxies"`
	Roots      []Root  `yaml:"roots"`
}

type Root struct {
	Method   string   `yaml:"method"`
	Path     string   `yaml:"path"`
	Response Response `yaml:"response"`
}

type Proxy struct {
	Id   string `yaml:"id"`
	Host string `yaml:"host"`
}

type Response struct {
	Type    string            `yaml:"type"`
	Headers map[string]string `yaml:"headers"`

	//Proxy
	ProxyPath string `yaml:"proxyPath"`
	ProxyId   string `yaml:"proxyId"`

	//Response
	Code int    `yaml:"code"`
	Body string `yaml:"body"`
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
