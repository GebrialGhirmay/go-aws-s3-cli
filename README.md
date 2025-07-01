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
git clone https://github.com/GebrialGhirmay/go-aws-s3-cli.git
cd go-aws-s3-cli
go mod tidy
go build -o go-aws-s3-cli main.go
```

### Download Pre-built Binaries

Pre-built binaries are available for multiple platforms on the [Releases](https://github.com/GebrialGhirmay/go-aws-s3-cli/releases) page.

#### Available Platforms:
- **Linux**: `amd64` and `arm64` architectures
- **Windows**: `amd64` architecture  
- **macOS**: `amd64` (Intel) and `arm64` (Apple Silicon) architectures

#### Download and Install:

**Linux/macOS:**
```bash
# Download the appropriate binary for your platform
# Example for Linux amd64:
curl -L -o go-aws-s3-cli.tar.gz https://github.com/GebrialGhirmay/go-aws-s3-cli/releases/latest/download/go-aws-s3-cli-linux-amd64.tar.gz

# Extract the binary
tar -xzf go-aws-s3-cli.tar.gz

# Make it executable
chmod +x go-aws-s3-cli-linux-amd64

# Move to a directory in your PATH (optional)
sudo mv go-aws-s3-cli-linux-amd64 /usr/local/bin/go-aws-s3-cli
```

**Windows:**
1. Go to the [Releases](https://github.com/GebrialGhirmay/go-aws-s3-cli/releases) page
2. Download `go-aws-s3-cli-windows-amd64.exe.zip`
3. Extract the ZIP file
4. Run `go-aws-s3-cli-windows-amd64.exe` from Command Prompt or PowerShell

**Using GitHub CLI:**
```bash
# Install GitHub CLI first, then:
gh release download --repo GebrialGhirmay/go-aws-s3-cli --pattern "*linux-amd64*"
```

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

## Creating Releases

### For Maintainers

To create a new release and trigger the build workflow:

1. **Create and push a version tag:**
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```

2. **Or create a release through GitHub UI:**
   - Go to your repository on GitHub
   - Click "Releases" → "Create a new release"
   - Create a new tag (e.g., `v1.0.0`)
   - Add release title and description
   - Click "Publish release"

3. **Manual trigger (if needed):**
   - Go to "Actions" tab in your repository
   - Select "Build and Release Multi-Platform"
   - Click "Run workflow"

The workflow will automatically:
- Build binaries for all supported platforms
- Create compressed archives (`.tar.gz` for Unix, `.zip` for Windows)
- Upload them to the GitHub release
- Generate release notes

### For Users

#### Finding Releases

1. **GitHub Releases Page**: Navigate to `https://github.com/GebrialGhirmay/go-aws-s3-cli/releases`
2. **Latest Release**: Use the "Latest" badge or go directly to `/releases/latest`
3. **Specific Version**: Each release is tagged with a version number (e.g., v1.0.0)

#### Downloading Executables

**Method 1: Direct Download**
- Visit the releases page
- Find the latest release
- Download the appropriate file for your platform:
  - `go-aws-s3-cli-linux-amd64.tar.gz` - Linux 64-bit
  - `go-aws-s3-cli-linux-arm64.tar.gz` - Linux ARM64
  - `go-aws-s3-cli-windows-amd64.exe.zip` - Windows 64-bit
  - `go-aws-s3-cli-darwin-amd64.tar.gz` - macOS Intel
  - `go-aws-s3-cli-darwin-arm64.tar.gz` - macOS Apple Silicon

**Method 2: Using curl/wget**
```bash
# Get the latest release for Linux amd64
curl -L -o go-aws-s3-cli.tar.gz \
  "https://github.com/GebrialGhirmay/go-aws-s3-cli/releases/latest/download/go-aws-s3-cli-linux-amd64.tar.gz"
```

**Method 3: Using GitHub CLI**
```bash
gh release download --repo GebrialGhirmay/go-aws-s3-cli
```

