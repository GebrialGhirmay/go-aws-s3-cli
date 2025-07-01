// config_loader.go creates a function that loads the configuration data from a config file and returns a Config struct.

package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//LoadConfig function, which is very similar to the NewConfig function in config.go. It also loads the AWS access key ID and AWS secret access key from environment variables using os.LookupEnv or if the environment variables are not set to try loading from .aws/credentials

//Both NewConfig (in the config file) and LoadConfig return a pointer to a Config struct (from config.go) with the loaded values for the AWS access key ID, AWS secret access key, and default values for the S3 bucket name, CloudFront distribution ID, and log level.

func LoadConfig() (*Config, error) {
	// Try environment variables first
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// If environment variables are not set, try loading from .aws/credentials
	if accessKeyID == "" || secretAccessKey == "" {
		awsAccessKey, awsSecretKey, err := loadCredentialsFromFile()
		if err != nil {
			return nil, fmt.Errorf("failed to load credentials from environment variables or .aws/credentials: %v", err)
		}
		if accessKeyID == "" {
			accessKeyID = awsAccessKey
		}
		if secretAccessKey == "" {
			secretAccessKey = awsSecretKey
		}
	}

	// Final validation
	if accessKeyID == "" {
		return nil, errors.New("no AWS_ACCESS_KEY_ID found in environment variables or .aws/credentials")
	}
	if secretAccessKey == "" {
		return nil, errors.New("no AWS_SECRET_ACCESS_KEY found in environment variables or .aws/credentials")
	}

	return &Config{
		AWSAccessKeyID:     accessKeyID,
		AWSSecretAccessKey: secretAccessKey,
		S3BucketName:       "am1gocli",
		CloudFrontDistID:   "EWF9J0XNACRVF",
		LogLevel:           "info",
	}, nil
}

func loadCredentialsFromFile() (string, string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", "", fmt.Errorf("failed to get user home directory: %v", err)
	}

	credentialsPath := filepath.Join(homeDir, ".aws", "credentials")
	
	file, err := os.Open(credentialsPath)
	if err != nil {
		return "", "", fmt.Errorf("failed to open .aws/credentials file: %v", err)
	}
	defer file.Close()

	var accessKeyID, secretAccessKey string
	var inDefaultProfile bool
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		
		// Check for profile section
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			profileName := strings.Trim(line, "[]")
			inDefaultProfile = (profileName == "default")
			continue
		}
		
		// Parse key-value pairs only if we're in the default profile
		if inDefaultProfile && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				
				switch key {
				case "aws_access_key_id":
					accessKeyID = value
				case "aws_secret_access_key":
					secretAccessKey = value
				}
			}
		}
	}
	
	if err := scanner.Err(); err != nil {
		return "", "", fmt.Errorf("error reading .aws/credentials file: %v", err)
	}
	
	if accessKeyID == "" || secretAccessKey == "" {
		return "", "", errors.New("aws_access_key_id or aws_secret_access_key not found in [default] profile")
	}
	
	return accessKeyID, secretAccessKey, nil
}