package dependencies

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name    string `mapstructure:"name"`
		Address string `mapstructure:"address"`
		Key     string `mapstructure:"key"`
		Port    int    `mapstructure:"port"`
		Debug   bool   `mapstructure:"debug"`
	} `mapstructure:"app"`
}

func LoadConfig(path string) (*Config, error) {
	// Specify the config file path
	viper.SetConfigFile(path)
	viper.SetConfigType("json")

	// Enable reading from environment variables
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Convert `.` to `_` in env vars

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Manually replace placeholders like ${ENV_VAR} with actual environment variable values
	configMap := viper.AllSettings() // Get all config as a map
	replacePlaceholders(configMap)   // Replace placeholders in the map

	// Write back the modified config to Viper
	for key, value := range configMap {
		viper.Set(key, value)
	}

	// Unmarshal into the Config struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &config, nil
}

// Replace placeholders in a map recursively
func replacePlaceholders(configMap map[string]interface{}) {
	for key, value := range configMap {
		switch v := value.(type) {
		case string:
			if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "}") {
				envVar := strings.TrimSuffix(strings.TrimPrefix(v, "${"), "}")
				configMap[key] = getEnv(envVar, v) // Replace with env var value or keep as-is
			}
		case map[string]interface{}:
			replacePlaceholders(v) // Recurse for nested maps
		}
	}
}

// Helper function to get an environment variable value
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
