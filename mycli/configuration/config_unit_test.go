//In the follow test:

//Test case 1: We call NewConfig without setting any environment variables and assert that the AWSAccessKeyID and AWSSecretAccessKey fields in the returned Config struct are empty strings. This tests the behavior when the environment variables are not set.

//Test case 2: This is similar to your original test. We set the environment variables to valid test values, call NewConfig, and assert that the AWSAccessKeyID and AWSSecretAccessKey fields contain the expected values.

//Test case 3: the modified TestNewConfig function tests the behavior of os.LookupEnv when the environment variables are set to invalid values (empty strings, in this case).In the NewConfig function from the config.go file, it's using os.LookupEnv to retrieve the values of the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables:

//By setting the environment variables to empty strings and then calling NewConfig, we're testing the behavior of os.LookupEnv when it encounters these invalid values. The assertion checks that the AWSAccessKeyID and AWSSecretAccessKey fields in the returned Config struct are empty strings, as expected when the environment variables are set to empty strings.

//After adding these test cases, the unit test will cover the following scenarios:

//Environment variables are not set
//Environment variables are set to valid values
//Environment variables are set to invalid values (empty strings)

//By testing these different scenarios ensures that the NewConfig function handles missing or invalid environment variables correctly, and I can verify the behavior of os.LookupEnv when reading the AWS credentials.

package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	// Test case 1: Environment variables are not set
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")

	cfg := NewConfig()
	if cfg.AWSAccessKeyID != "" || cfg.AWSSecretAccessKey != "" {
		t.Errorf("Expected empty AWS credentials when environment variables are not set")
	}

	// Test case 2: Environment variables are set to valid values
	os.Setenv("AWS_ACCESS_KEY_ID", "test-access-key")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "test-secret-key")
	defer os.Unsetenv("AWS_ACCESS_KEY_ID")
	defer os.Unsetenv("AWS_SECRET_ACCESS_KEY")

	cfg = NewConfig()
	if cfg.AWSAccessKeyID != "test-access-key" {
		t.Errorf("Incorrect AWSAccessKeyID: %s", cfg.AWSAccessKeyID)
	}
	if cfg.AWSSecretAccessKey != "test-secret-key" {
		t.Errorf("Incorrect AWSSecretAccessKey: %s", cfg.AWSSecretAccessKey)
	}

	// Test case 3: Environment variables are set to invalid values (empty strings)
	os.Setenv("AWS_ACCESS_KEY_ID", "")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "")

	cfg = NewConfig()
	if cfg.AWSAccessKeyID != "" || cfg.AWSSecretAccessKey != "" {
		t.Errorf("Expected empty AWS credentials when environment variables are set to empty strings")
	}

	// Test S3 bucket name
	expectedBucketName := "s3://content.lumen-research.com/cachepages/release/AM1 Go CLI Project /"
	if cfg.S3BucketName != expectedBucketName {
		t.Errorf("Incorrect S3BucketName: %s", cfg.S3BucketName)
	}

	// Test CloudFront distribution ID
	expectedCloudFrontDistID := "IELI6WX9MC9WSEFR9VFCDBVWZX"
	if cfg.CloudFrontDistID != expectedCloudFrontDistID {
		t.Errorf("Incorrect CloudFrontDistID: %s", cfg.CloudFrontDistID)
	}

	// Test log level
	expectedLogLevel := "info"
	if cfg.LogLevel != expectedLogLevel {
		t.Errorf("Incorrect LogLevel: %s", cfg.LogLevel)
	}
}
