//In the userinput and upload package, implement functions to parse command-line arguments and validate user input as well as provide for the file to eb uploaded.

// fileupload/fileupload.go

package fileupload

import (
	"fmt"
	"os"

	awsClient "go-aws-s3-cli/mycli/aws"
	awsConfig "go-aws-s3-cli/mycli/configuration"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	//"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadFile(filePath string) error {
	// Load the configuration
	cfg, err := awsConfig.LoadConfig()
	if err != nil {
		return err
	}

	// Create an S3 client
	//s3Client, err := awsClient.NewS3Client()
	///if err != nil {
		//return err
	//}

	// Open the file for reading
	file, err := os.Open(filePath)
	if err != nil {
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

	// Create an uploader with the S3 client and the AWS session
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
	})

	// Upload the file to S3
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(cfg.S3BucketName),
		Key:    aws.String("3500.html"),
		Body:   file,
	})
	if err != nil {
		return err
	}

	// Print the upload result
	fmt.Println("Successfully uploaded file to S3:", result.Location)

	// Invalidate the CloudFront cache for the uploaded object
	err = awsClient.InvalidateCloudFrontCache("s3://am1gocli/1.3500.html")
	if err != nil {
		return err
	}

	return nil
}
