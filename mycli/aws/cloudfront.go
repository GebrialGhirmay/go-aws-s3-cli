package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"go-aws-s3-cli/mycli/configuration"
)

func InvalidateCloudFrontCache(objectKey string) error {
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-2"),
	})
	if err != nil {
		return err
	}

	svc := cloudfront.New(sess)

	invalidationBatch := &cloudfront.InvalidationBatch{
		CallerReference: aws.String("go-cli-invalidation"),
		Paths: &cloudfront.Paths{
			Quantity: aws.Int64(1),
			Items: []*string{
				aws.String("/" + objectKey),
			},
		},
	}

	input := &cloudfront.CreateInvalidationInput{
		DistributionId:    aws.String(cfg.CloudFrontDistID),
		InvalidationBatch: invalidationBatch,
	}

	_, err = svc.CreateInvalidation(input)
	if err != nil {
		return fmt.Errorf("failed to create invalidation: %v", err)
	}

	fmt.Println("Cache invalidation request submitted successfully.")
	return nil
}