package utils

import (
	"time"

	"github.com/pelletier/go-toml"
)

//Configuration - configuration structure
type Configuration struct {
	LogFilePath                 string
	Port                        string
	ServedURL                   string
	ReadTimeout                 time.Duration
	WriteTimeout                time.Duration
	FileSaveExtensionList       string
	ScaledImageRestoreExtension string
	ScaledImageH                uint
	ScaledImageW                uint
	DataBasePath                string
	FileSavePath                string
	PreviewFileFolder           string
}

//LoadConfiguration - load configuration file
func LoadConfiguration(path string) *Configuration {
	config, err := toml.LoadFile(path)

	if err != nil {
		panic(err)
	}

	configuration := Configuration{}
	config.Unmarshal(&configuration)

	return &configuration
}
