package config

// Config represents the configuration settings for the CLI

//Configuration Management:

//You would use a configuration management approach to handle settings such as AWS credentials, S3 bucket names, and other configurable parameters.

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

type Config struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	S3BucketName       string
	CloudFrontDistID   string
	LogLevel           string
	// Add other configuration fields as needed
}
