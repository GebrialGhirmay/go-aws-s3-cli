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

type Config struct {
	AWSAccessKeyID     string // AWS access key ID
	AWSSecretAccessKey string // AWS secret access key
	S3BucketName       string // Name of the S3 bucket to interact with
	CloudFrontDistID   string // Name of the ID of the CloudFront distribution to invalidate
	LogLevel           string // Logging level of the AWS SDK (eg. "debug", "info", "warn")
	// Add other configuration fields as needed
}
