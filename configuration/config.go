package config

//Once you've set environment variables, you can load them into the Go application using the os package which is imported here.

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//This defines the Config struct, which holds the configuration values for the application, such as the AWS access key ID, AWS secret access key, S3 bucket name, CloudFront distribution ID, and log level.

type Config struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	S3BucketName       string
	CloudFrontDistID   string
	LogLevel           string
}

// NewConfig function is responsible for loading the AWS access key ID and AWS secret access key from environment variables using the os.LookupEnv function. If the environment variables are not set or empty, it will try loading from .aws/credentials. 

func NewConfig() (*Config, error) {
	// Try environment variables first
	accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")

	// If environment variables are not set, try loading from .aws/credentials
	if accessKeyID == "" || secretAccessKey == "" {
		awsAccessKey, awsSecretKey, err := loadFromAWSCredentials()
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

func loadFromAWSCredentials() (string, string, error) {
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