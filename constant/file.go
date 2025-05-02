package constant

import "os"

type FilePath = string

var dir string

func init() {
	dir, _ = os.Getwd()
}

var (
	developmentConfig = dir + "/config/environment/development"
	productionConfig  = dir + "/config/environment/production"
	testConfig        = dir + "/config/environment/test"
)

type configFilePath struct {
	DevelopmentConfig FilePath
	ProductionConfig  FilePath
	TestConfig        FilePath
}

var Path = configFilePath{
	developmentConfig,
	productionConfig,
	testConfig,
}
