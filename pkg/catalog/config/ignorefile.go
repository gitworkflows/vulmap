package config

import (
	"os"

	"github.com/projectdiscovery/gologger"
	"gopkg.in/yaml.v2"
)

// IgnoreFile is an internal vulmap template blocking configuration file
type IgnoreFile struct {
	Tags  []string `yaml:"tags"`
	Files []string `yaml:"files"`
}

// ReadIgnoreFile reads the vulmap ignore file returning blocked tags and paths
func ReadIgnoreFile() IgnoreFile {
	file, err := os.Open(DefaultConfig.GetIgnoreFilePath())
	if err != nil {
		gologger.Error().Msgf("Could not read vulmap-ignore file: %s\n", err)
		return IgnoreFile{}
	}
	defer file.Close()

	ignore := IgnoreFile{}
	if err := yaml.NewDecoder(file).Decode(&ignore); err != nil {
		gologger.Error().Msgf("Could not parse vulmap-ignore file: %s\n", err)
		return IgnoreFile{}
	}
	return ignore
}
