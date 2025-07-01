package config //declares the package name for this Go source file and the package name is config.

import ( //imports three packages from the Go standard library: os, testing and path/filepath. The os package provides functionality for interacting with the operating system, such as setting or reading environment variables. The testing package is used for writing and running unit tests in Go.The package filepath implements utility routines for manipulating filename paths in a way compatible with the target operating system-defined file paths.
	"os"
	"testing"
    "path/filepath"

)

func TestNewConfig(t *testing.T) {
	// Test case 1: Environment variables are not set and no .aws/credentials
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")

	cfg, err := NewConfig()
	if err == nil {
		t.Errorf("Expected an error when no credentials are available, but got nil")
		return
	}

	// Test case 2: Environment variables are set to valid values
	os.Setenv("AWS_ACCESS_KEY_ID", "test-access-key")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "test-secret-key")
	defer os.Unsetenv("AWS_ACCESS_KEY_ID")
	defer os.Unsetenv("AWS_SECRET_ACCESS_KEY")

	cfg, err = NewConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	// Validate environment variable values
	if cfg.AWSAccessKeyID != "test-access-key" {
		t.Errorf("Expected AWSAccessKeyID to be 'test-access-key', got '%s'", cfg.AWSAccessKeyID)
	}
	if cfg.AWSSecretAccessKey != "test-secret-key" {
		t.Errorf("Expected AWSSecretAccessKey to be 'test-secret-key', got '%s'", cfg.AWSSecretAccessKey)
	}

	// Test default values
	expectedBucketName := "am1gocli"
	if cfg.S3BucketName != expectedBucketName {
		t.Errorf("Incorrect S3BucketName: %s", cfg.S3BucketName)
	}

	expectedCloudFrontDistID := "EWF9J0XNACRVF"
	if cfg.CloudFrontDistID != expectedCloudFrontDistID {
		t.Errorf("Incorrect CloudFrontDistID: %s", cfg.CloudFrontDistID)
	}

	expectedLogLevel := "info"
	if cfg.LogLevel != expectedLogLevel {
		t.Errorf("Incorrect LogLevel: %s", cfg.LogLevel)
	}
}

func TestLoadConfigWithAWSCredentials(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("AWS_ACCESS_KEY_ID")
	os.Unsetenv("AWS_SECRET_ACCESS_KEY")

	// Create a temporary .aws/credentials file
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get home directory: %v", err)
	}

	awsDir := filepath.Join(homeDir, ".aws")
	credentialsPath := filepath.Join(awsDir, "credentials")

	// Create .aws directory if it doesn't exist
	err = os.MkdirAll(awsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create .aws directory: %v", err)
	}

	// Create test credentials file
	credentialsContent := `[default]
aws_access_key_id = file-access-key
aws_secret_access_key = file-secret-key

[other-profile]
aws_access_key_id = other-access-key
aws_secret_access_key = other-secret-key
`

	err = os.WriteFile(credentialsPath, []byte(credentialsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test credentials file: %v", err)
	}

	// Clean up after test
	defer os.Remove(credentialsPath)

	// Test loading from file
	cfg, err := LoadConfig()
	if err != nil {
		t.Errorf("Unexpected error loading from .aws/credentials: %v", err)
		return
	}

	// Validate file-based values
	if cfg.AWSAccessKeyID != "file-access-key" {
		t.Errorf("Expected AWSAccessKeyID to be 'file-access-key', got '%s'", cfg.AWSAccessKeyID)
	}
	if cfg.AWSSecretAccessKey != "file-secret-key" {
		t.Errorf("Expected AWSSecretAccessKey to be 'file-secret-key', got '%s'", cfg.AWSSecretAccessKey)
	}
}

func TestEnvironmentVariablesPrecedence(t *testing.T) {
	// Set environment variables
	os.Setenv("AWS_ACCESS_KEY_ID", "env-access-key")
	os.Setenv("AWS_SECRET_ACCESS_KEY", "env-secret-key")
	defer os.Unsetenv("AWS_ACCESS_KEY_ID")
	defer os.Unsetenv("AWS_SECRET_ACCESS_KEY")

	// Create .aws/credentials file with different values
	homeDir, err := os.UserHomeDir()
	if err != nil {
		t.Fatalf("Failed to get home directory: %v", err)
	}

	awsDir := filepath.Join(homeDir, ".aws")
	credentialsPath := filepath.Join(awsDir, "credentials")

	err = os.MkdirAll(awsDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create .aws directory: %v", err)
	}

	credentialsContent := `[default]
aws_access_key_id = file-access-key
aws_secret_access_key = file-secret-key
`

	err = os.WriteFile(credentialsPath, []byte(credentialsContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test credentials file: %v", err)
	}

	defer os.Remove(credentialsPath)

	// Test that environment variables take precedence
	cfg, err := LoadConfig()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
		return
	}

	// Should use environment variable values, not file values
	if cfg.AWSAccessKeyID != "env-access-key" {
		t.Errorf("Expected AWSAccessKeyID to be 'env-access-key' (from env), got '%s'", cfg.AWSAccessKeyID)
	}
	if cfg.AWSSecretAccessKey != "env-secret-key" {
		t.Errorf("Expected AWSSecretAccessKey to be 'env-secret-key' (from env), got '%s'", cfg.AWSSecretAccessKey)
	}
}