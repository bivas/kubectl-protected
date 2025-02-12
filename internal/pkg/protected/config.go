package protected

import (
	"io"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of the configuration file
type Config struct {
	Protected []string `yaml:"protected"`
}

// ReadProtectedPatterns reads the protected patterns from a YAML configuration file
func ReadProtectedPatterns(reader io.Reader) ([]string, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return config.Protected, nil
}
