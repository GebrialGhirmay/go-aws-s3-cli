package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/google/uuid" // Imports the UUID library (had to run go mod tidy command)
	"go-aws-s3-cli/mycli/configuration"
)

//The InvalidateCloudFrontCache function takes an objectKey string as input, which represents the key (filename) of the object for which the CloudFront cache needs to be invalidated.

func InvalidateCloudFrontCache(objectKey string) error {
	cfg, err := config.LoadConfig() //the function first loads the AWS configuration using config.LoadConfig(). This configuration contains the CloudFront distribution ID (cfg.CloudFrontDistID), which is required for creating the cache invalidation request.
	if err != nil {
		return err
	}

	sess, err := session.NewSession(&aws.Config{ //A new AWS session is created using session.NewSession with the specified AWS region ("eu-west-2").
		Region: aws.String("eu-west-2"),
	})
	if err != nil {
		return err
	}

	svc := cloudfront.New(sess) //A new AWS session is created using session.NewSession with the specified AWS region.

	// Generate a unique caller reference
	callerReference := uuid.New().String()

	//An invalidationBatch is created, which contains the unique caller reference and the path of the object to be invalidated. The path is constructed by prepending a forward slash (/) to the objectKey.

	invalidationBatch := &cloudfront.InvalidationBatch{
		CallerReference: aws.String(callerReference),
		Paths: &cloudfront.Paths{
			Quantity: aws.Int64(1),
			Items: []*string{
				aws.String("/" + objectKey),
			},
		},
	}
	//A CreateInvalidationInput is created, which includes the CloudFront distribution ID (cfg.CloudFrontDistID) and the invalidationBatch.

	input := &cloudfront.CreateInvalidationInput{
		DistributionId:    aws.String(cfg.CloudFrontDistID),
		InvalidationBatch: invalidationBatch,
	}

	//The svc.CreateInvalidation function is called with the CreateInvalidationInput to submit the cache invalidation request to CloudFront.

	_, err = svc.CreateInvalidation(input)
	if err != nil {
		return fmt.Errorf("failed to create invalidation: %v", err)
	}

	fmt.Println("Cache invalidation request submitted successfully.")
	return nil
}