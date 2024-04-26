// Config represents the configuration settings for the CLI

//Configuration Management:

//I would use a configuration management approach to handle settings such as AWS credentials, S3 bucket names, and other configurable parameters.

//Advantages:

//Allows for flexibility and easy configuration changes.

//Separates configuration concerns from the application logic.

//Considerations:

//Utilize configuration files or environment variables for flexibility.

//Implement a configuration parser to read and validate settings.

//Store configuration details (e.g., AWS credentials, S3 bucket names) in configuration files or environment variables.

//Avoid hardcoding configuration values within the code.

//Advantages:

//Enhances flexibility and maintainability.

//Allows for easy configuration changes without modifying code.

package config

//Once you've set these environment variables, you can load them into the Go application using the os package
import (
	"errors"
	"os"
)

type Config struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	S3BucketName       string
	CloudFrontDistID   string
	LogLevel           string
}

// the NewConfig function now uses os.LookupEnv to retrieve the values of the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables and assigns them to the respective fields in the Config struct.

// the LogLevel is set to "info", which means that log messages with a severity level of info or higher (e.g., warn, error) will be logged, while messages with a lower level (e.g., debug) will be ignored.

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
        S3BucketName:       "s3://content.lumen-research.com/cachepages/release/AM1 Go CLI Project /",
        CloudFrontDistID:   "IELI6WX9MC9WSEFR9VFCDBVWZX",
        LogLevel:           "info",
    }, nil
}

