package main

import (
	"fmt"
	"go-aws-s3-cli/mycli/aws"

	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	s3Client, err := aws.NewS3Client()
	if err != nil {
		fmt.Println("Error creating S3 client:", err)
		return
	}

	bucketName := "am1gocli"

	result, err := s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: &bucketName,
	})
	if err != nil {
		fmt.Println("Error listing S3 objects:", err)
		return
	}

	fmt.Printf("Objects in S3 bucket '%s':\n", bucketName)
	for _, obj := range result.Contents {
		fmt.Println("-", *obj.Key)
	}
}
