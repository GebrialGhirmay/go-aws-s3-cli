package config //declares the package name for this Go source file and the package name is config.

import ( //imports two packages from the Go standard library: os and testing. The os package provides functionality for interacting with the operating system, such as setting or reading environment variables. The testing package is used for writing and running unit tests in Go.
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
    // Test case 1: Environment variables are not set
    os.Unsetenv("AWS_ACCESS_KEY_ID")
    os.Unsetenv("AWS_SECRET_ACCESS_KEY")

    cfg, err := NewConfig()
    if err == nil {
        t.Errorf("Expected an error when environment variables are not set, but got nil")
        return // Exit the function to avoid nil pointer dereference
    }

    // Test case 2: Environment variables are set to valid values
    os.Setenv("AWS_ACCESS_KEY_ID", "test-access-key")
    os.Setenv("AWS_SECRET_ACCESS_KEY", "test-secret-key")
    defer os.Unsetenv("AWS_ACCESS_KEY_ID")
    defer os.Unsetenv("AWS_SECRET_ACCESS_KEY")

    cfg, err = NewConfig()
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
        return // Exit the function to avoid nil pointer dereference
    }

    // Add assertions to check the values of cfg fields
    // Test S3 bucket name
    expectedBucketName := "am1gocli"
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

