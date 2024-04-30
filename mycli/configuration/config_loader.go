// config_loader.go creates a function that loads the configuration data from the a config file and returns a Config struct.

package config

import (
	"errors"
	"os"
)

//LoadConfig function, which is very similar to the NewConfig function in config.go. It also loads the AWS access key ID and AWS secret access key from environment variables using os.LookupEnv.

//Both NewConfig (in the config file) and LoadConfig return a pointer to a Config struct (from config.go) with the loaded values for the AWS access key ID, AWS secret access key, and default values for the S3 bucket name, CloudFront distribution ID, and log level.

func LoadConfig() (*Config, error) {
	accessKeyID, exists := os.LookupEnv("AWS_ACCESS_KEY_ID")
	if !exists || accessKeyID == "" {
		return nil, errors.New("no value provided for AWS_ACCESS_KEY_ID")
	}

	secretAccessKey, exists := os.LookupEnv("AWS_SECRET_ACCESS_KEY")
	if !exists || secretAccessKey == "" {
		return nil, errors.New("no value provided for AWS_SECRET_ACCESS_KEY")
	}

	return &Config{
		AWSAccessKeyID:     accessKeyID,
		AWSSecretAccessKey: secretAccessKey,
		S3BucketName:       "am1gocli",
		CloudFrontDistID:   "IELI6WX9MC9WSEFR9VFCDBVWZX",
		LogLevel:           "info",
	}, nil
}
