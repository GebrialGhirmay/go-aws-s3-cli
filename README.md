# Go AWS S3 CLI

A command-line interface (CLI) application written in Go for uploading files to AWS S3 with automatic CloudFront cache invalidation. This tool is specifically designed for uploading HTML files as cache pages to Amazon S3 and managing CloudFront distribution cache invalidation.

## Features

- **File Upload to S3**: Upload files to a specified S3 bucket with HTML content type
- **CloudFront Cache Invalidation**: Automatically invalidate CloudFront cache after file upload
- **Flexible Credential Management**: Supports AWS credentials from environment variables or `~/.aws/credentials` file
- **Comprehensive Logging**: Multi-level logging with both console and file output
- **Error Handling**: Robust error handling with descriptive messages
- **Cross-Platform**: Supports Windows, macOS, and Linux

## Architecture

The application is structured with the following components:

- **CLI Interface** (`cli/`): Command-line argument parsing using Cobra
- **Configuration** (`config/`): AWS credentials and configuration management
- **File Upload** (`fileupload/`): S3 file upload functionality
- **AWS Services** (`aws/`): S3 and CloudFront service clients
- **Logging** (`logging/`): Comprehensive logging system

## Installation

### Prerequisites

- Go 1.20 or later
- AWS credentials configured (see [Configuration](#configuration))

### Build from Source

```bash
git clone https://github.com/your-username/go-aws-s3-cli.git
cd go-aws-s3-cli
go mod tidy
go build -o go-aws-s3-cli main.go
```

### Download Pre-built Binaries

Check the [Releases](https://github.com/your-username/go-aws-s3-cli/releases) page for pre-built binaries for your platform.

## Configuration

### AWS Credentials

The application supports two methods for AWS credential configuration:

#### 1. Environment Variables (Recommended)

```bash
export AWS_ACCESS_KEY_ID="your-access-key-id"
export AWS_SECRET_ACCESS_KEY="your-secret-access-key"
```

#### 2. AWS Credentials File

Create or modify `~/.aws/credentials`:

```ini
[default]
aws_access_key_id = your-access-key-id
aws_secret_access_key = your-secret-access-key
```

### Default Configuration

The application uses the following default settings:

- **AWS Region**: `eu-west-2`
- **S3 Bucket**: `am1gocli`
- **CloudFront Distribution ID**: `EWF9J0XNACRVF`
- **Content Type**: `text/html`
- **Log Level**: `info`

## Usage

### Basic Usage

```bash
# Upload a file using long flag
./go-aws-s3-cli --file /path/to/your/file.html

# Upload a file using short flag
./go-aws-s3-cli -f /path/to/your/file.html
```

### Help

```bash
./go-aws-s3-cli --help
```

## AWS Permissions

Your AWS credentials must have the following permissions:

### S3 Permissions
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:PutObject",
                "s3:PutObjectAcl"
            ],
            "Resource": "arn:aws:s3:::am1gocli/*"
        }
    ]
}
```

### CloudFront Permissions
```json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "cloudfront:CreateInvalidation"
            ],
            "Resource": "arn:aws:cloudfront::*:distribution/EWF9J0XNACRVF"
        }
    ]
}
```

## Workflow

1. **Configuration Loading**: Loads AWS credentials from environment variables or `~/.aws/credentials`
2. **File Validation**: Validates the provided file path
3. **S3 Upload**: Uploads the file to the configured S3 bucket with HTML content type
4. **CloudFront Invalidation**: Automatically invalidates the CloudFront cache for the uploaded file
5. **Logging**: Comprehensive logging throughout the process

## Logging

The application provides detailed logging with multiple levels:

- **DEBUG**: Detailed execution information
- **INFO**: General operational messages  
- **WARNING**: Non-critical issues
- **ERROR**: Critical failures

Logs are written to:
- Console output (stdout)
- Log file (`app.log`)

## Error Handling

The application handles various error scenarios:

- Missing or invalid AWS credentials
- Invalid file paths or inaccessible files
- AWS service failures (S3 upload, CloudFront invalidation)
- Network connectivity issues

All errors are logged with appropriate context and returned with descriptive messages.

## Development

### Project Structure

```
go-aws-s3-cli/
├── main.go                 # Application entry point
├── cli/
│   └── cli.go             # CLI interface and command handling
├── config/
│   ├── config.go          # Configuration structure and loading
│   ├── config_loader.go   # Alternative configuration loader
│   └── config_unit_test.go # Unit tests for configuration
├── aws/
│   ├── s3.go              # S3 client configuration
│   └── cloudfront.go      # CloudFront cache invalidation
├── fileupload/
│   └── fileupload.go      # File upload logic
├── logging/
│   └── logging.go         # Logging system
├── .github/workflows/
│   ├── release.yml        # Linux build and release
│   └── windows-build.yml  # Windows build workflow
└── README.md
```

### Dependencies

The project uses the following key dependencies:

- `github.com/aws/aws-sdk-go` - AWS SDK for Go
- `github.com/spf13/cobra` - CLI framework
- `github.com/google/uuid` - UUID generation for CloudFront invalidation

### Running Tests

```bash
go test ./...
```

### Building for Different Platforms

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o go-aws-s3-cli-linux main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o go-aws-s3-cli.exe main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o go-aws-s3-cli-macos main.go
```

## CI/CD

The project includes GitHub Actions workflows for automated building and releasing:

- **Linux Build**: Triggers on changes to `main.go`
- **Windows Build**: Triggers on release creation
- **Automatic Releases**: Binaries are automatically attached to GitHub releases


### v1.0.0
- Initial release
- Basic S3 file upload functionality
- CloudFront cache invalidation
- Environment variable and AWS credentials file support
- Comprehensive logging system