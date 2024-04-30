package config

//Once you've set environment variables, you can load them into the Go application using the os package which is imported here.

import (
	"errors"
	"os"
)

//This defines the Config struct, which holds the configuration values for the application, such as the AWS access key ID, AWS secret access key, S3 bucket name, CloudFront distribution ID, and log level.

type Config struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	S3BucketName       string
	CloudFrontDistID   string
	LogLevel           string
}

// NewConfig function is responsible for loading the AWS access key ID and AWS secret access key from environment variables using the os.LookupEnv function. If the environment variables are not set or empty, it returns an error.

// NewConfig (like LoadConfig in configloader file) returns a pointer to a Config struct with the loaded values for the AWS access key ID, AWS secret access key, and default values for the S3 bucket name, CloudFront distribution ID, and log level.

func NewConfig() (*Config, error) {
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
		LogLevel:           "info", // the LogLevel is set to "info", which means that log messages with a severity level of info or higher (e.g., warn, error) will be logged, while messages with a lower level (e.g., debug) will be ignored.
	}, nil
}
