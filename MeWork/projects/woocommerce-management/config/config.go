// config/config.go
package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	WooCommerceURL            string
	WooCommerceConsumerKey    string
	WooCommerceConsumerSecret string
}

// LoadConfig loads the application configuration from the given file path
func LoadConfig(filePath string) (*Config, error) {
	viper.SetConfigFile(filePath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{
		WooCommerceURL:            viper.GetString("woocommerce.url"),
		WooCommerceConsumerKey:    viper.GetString("woocommerce.consumer_key"),
		WooCommerceConsumerSecret: viper.GetString("woocommerce.consumer_secret"),
	}

	log.Println("Configuration loaded successfully")
	return config, nil
}
