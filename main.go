package main

import (
	"flag"

	mocksrv "github.com/sergonas/mocksrv/pkg"
)

var (
	configPath = flag.String("configPath", "config/config.yaml", "path to config file")
	config     mocksrv.ConfigRoot
)

func main() {
	flag.Parse()
	config = mocksrv.Parse(*configPath)
	handler := mocksrv.Handler(config)

	mocksrv.Run(config, handler)
}
