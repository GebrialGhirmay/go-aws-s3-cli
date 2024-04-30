package aws

import (
	"fmt"
	config "go-aws-s3-cli/mycli/configuration"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewS3Client, calls the config.LoadConfig() function to retrieve the configuration values. The config package is imported as config "go-aws-s3-cli/mycli/configuration".

func NewS3Client() (*s3.S3, error) {
	cfg, err := config.LoadConfig() //LoadConfig function (defined in config_loader.go) returns a pointer to a Config struct that contains the loaded AWS access key ID and AWS secret access key from the environment variables.
	if err != nil {
		return nil, err
	}

	// Prints AWS credentials (to validate that it is reading them correctly from the env. variables)
	fmt.Println("AWS Access Key ID:", cfg.AWSAccessKeyID)
	fmt.Println("AWS Secret Access Key:", cfg.AWSSecretAccessKey)

	awsCredentials := credentials.NewStaticCredentials(cfg.AWSAccessKeyID, cfg.AWSSecretAccessKey, "")
	//NewS3Client function then uses the loaded AWS credentials from the Config struct to create an AWS credentials object (awsCredentials) using the credentials.NewStaticCredentials function from the AWS SDK for Go

	awsSession, err := session.NewSession(&aws.Config{ //The awsCredentials object is then used to create a new AWS session (awsSession) using the session.NewSession function from the AWS SDK for Go. The session is configured with the desired AWS region ("eu-west-2" in this case) and the loaded AWS credentials.
		Region:      aws.String("eu-west-2"),
		Credentials: awsCredentials,
	})
	if err != nil {
		return nil, err
	}

	s3Client := s3.New(awsSession) //an S3 client (s3Client) is created using the s3.New function from the AWS SDK for Go, passing the awsSession as an argument.
	return s3Client, nil
}
