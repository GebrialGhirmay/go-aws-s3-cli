package config

// Config represents the configuration settings for the CLI
type Config struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	S3BucketName       string
	CloudFrontDistID   string
	LogLevel           string
	// Add other configuration fields as needed
}
