package main

import (
	"os"
	"path/filepath"

	"github.com/bogdanfinn/tls-client-api/internal/tls-client-api/api"
	"github.com/justtrackio/gosoline/pkg/apiserver"
	"github.com/justtrackio/gosoline/pkg/application"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFilePath := filepath.Join(cwd, "config.dist.yml")

	application.Run(
		application.WithConfigFile(configFilePath, "yml"),
		application.WithConfigFileFlag,
		application.WithModuleFactory("tls-client-api", apiserver.New(api.DefineRouter)),
	)
}
