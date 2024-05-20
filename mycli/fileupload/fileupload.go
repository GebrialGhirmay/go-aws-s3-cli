//In the userinput and upload package, implement functions to parse command-line arguments and validate user input as well as provide for the file to eb uploaded.

// fileupload/fileupload.go

package fileupload


import (
	"fmt"
	"go-aws-s3-cli/mycli/logging" // Import the logger package
	"os"
	"path/filepath"

	awsConfig "go-aws-s3-cli/mycli/configuration"
	awsClient "go-aws-s3-cli/mycli/aws"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(filePath string) error {
	// Load the configuration
	cfg, err := awsConfig.LoadConfig()
	if err != nil {
		logging.Logger.Printf("Error loading configuration: %v", err) // Use the logger instance
		return err
	}

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
		logging.Logger.Printf("Error opening file: %v", err) // Use the logger instance
		return err
	}

	defer file.Close()
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	if err != nil {
		return err
	}

	// Create an uploader with the AWS session
	uploader := s3manager.NewUploader(sess)

	// Extract the filename from the file path
	_, filename := filepath.Split(filePath)

	// Upload the file to S3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.S3BucketName),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		return err
	}

	// Print the upload result
	fmt.Println("Successfully uploaded file to S3:", result.Location)

	// Invalidate the CloudFront cache for the uploaded object
	
	err = awsClient.InvalidateCloudFrontCache(filename)
	if err != nil {
		return err
	}

	return nil
}