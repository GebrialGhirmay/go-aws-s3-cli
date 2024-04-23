//In the userinput and upload package, implement functions to parse command-line arguments and validate user input as well as provide for the file to eb uploaded.

package upload //package upload declaration defines a new package named upload

//import statement imports the necessary packages required by the code

import (
	"bufio"   //Provides buffered I/O operations for reading user input.
	"fmt"     //Provides formatted I/O operations for printing output.
	"os"      //Provides functionality for interacting with the operating system, such as opening files.
	"strings" //Provides string manipulation functions.

	"github.com/aws/aws-sdk-go/aws"                  //The core package of the AWS SDK for Go.
	"github.com/aws/aws-sdk-go/aws/session"          // Provides functionality for creating AWS sessions.
	"github.com/aws/aws-sdk-go/service/s3"           //The package for interacting with the Amazon S3 service.
	"github.com/aws/aws-sdk-go/service/s3/s3manager" //Provides higher-level functionality for managing S3 uploads and downloads.
)

// UploadFilesFromUserInput prompts the user for file paths, uploads the files to S3, and invalidates the CloudFront distribution.

//The UploadFilesFromUserInput function is the main entry point for the file upload process.

// It takes three arguments:

//sess *session.Session: A pointer to an AWS session that will be used for interacting with AWS services.
//bucketName string: The name of the S3 bucket to upload files to.
//distributionID string: The ID of the CloudFront distribution to invalidate.

//The function first calls the getFilePaths function to prompt the user for file paths and get a slice of file paths.

//If there's an error getting the file paths, it returns the error.

//Next, it calls the uploadFilesToS3 function to upload the files to the specified S3 bucket, passing the AWS session, file paths, and bucket name.

//If there's an error uploading the files, it returns the error.

//If the files are uploaded successfully, it calls the invalidateCloudFrontDistribution function to invalidate the CloudFront distribution for the uploaded files, passing the AWS session, distribution ID, and file paths.

// The function returns any error that occurs during the CloudFront invalidation process.

func UploadFilesFromUserInput(sess *session.Session, bucketName, distributionID string) error {
	filePaths, err := getFilePaths()
	if err != nil {
		return err
	}

	err = uploadFilesToS3(sess, filePaths, bucketName)
	if err != nil {
		return err
	}

	return invalidateCloudFrontDistribution(sess, distributionID, filePaths)
}

// getFilePaths prompts the user to enter file paths and returns a slice of file paths.

//The getFilePaths function is responsible for prompting the user to enter file paths and returning a slice of file paths.
//It initializes an empty slice filePaths to store the file paths.
//It prints a prompt for the user to enter file paths or type "done" to finish.
//It enters an infinite for loop that continues until the user types "done".
//Inside the loop, it calls the getUserInput function to get the user's input for a file path.
//If the user's input is "done", it breaks out of the loop.
//Otherwise, it appends the entered file path to the filePaths slice using the append function.
//After the loop ends, if the filePaths slice is empty (no file paths were entered), it returns an error.
//If file paths were entered, it returns the filePaths slice and nil error.

func getFilePaths() ([]string, error) {
	var filePaths []string

	fmt.Println("Enter file paths (or 'done' to finish):")
	for {
		path := getUserInput("> ")
		if path == "done" {
			break
		}
		filePaths = append(filePaths, path)
	}

	if len(filePaths) == 0 {
		return nil, fmt.Errorf("no file paths provided")
	}

	return filePaths, nil
}

//The getUserInput function is a helper function that prompts the user for input and returns the input as a string.
//It creates a new bufio.Reader instance using bufio.NewReader(os.Stdin), which reads from the standard input (os.Stdin).
//It prints the provided prompt string using fmt.Print.
//It reads the user's input until a newline character ('\n') is encountered using reader.ReadString('\n').
//It ignores any error returned by ReadString (the _ blank identifier is used to discard the error value).
//It trims any leading or trailing whitespace characters from the user's input using strings.TrimSpace and returns the trimmed input string.

// getUserInput prompts the user for input and returns the input as a string.
func getUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

//uploadFilesTos3 defines a function named uploadFilesToS3 which takes three parameters:
//sess *session.Session: A pointer to an AWS session object, which represents the session used for communication with AWS services.
//filePaths []string: A slice containing paths to the files that need to be uploaded to S3.
//bucketName string: The name of the S3 bucket where the files will be uploaded.

//uploader := s3manager.NewUploader(sess): This line creates a new uploader object using the NewUploader function from the s3manager package.
// The uploader is configured to use the provided AWS session (sess) for authentication and communication with AWS S3.

// for _, filePath := range filePaths {: This line starts a loop that iterates over each file path in the filePaths slice. The underscore _ is used as a placeholder for the index value, which is not used inside the loop.

//Inside the loop, the code opens each file specified by the filePath variable using the os.Open function. If an error occurs during file opening, the function returns an error message using fmt.Errorf, indicating the file path and the error encountered. The defer statement ensures that the file is closed after the function exits, even if an error occurs.

//Inside the loop, the code initiates the file upload to the specified S3 bucket using the Upload method of the uploader object. It creates an UploadInput object specifying the S3 bucket (Bucket), the object key (Key), and the file content (Body). If an error occurs during the upload process, the function returns an error message.

//After each successful upload, the code prints a message indicating that the file has been uploaded successfully.
// }: marks the end of the loop
// return nil: This line returns nil to indicate that the function executed successfully without encountering any errors.

// uploadFilesToS3 uploads the specified HTML files to the given S3 bucket.
func uploadFilesToS3(sess *session.Session, filePaths []string, bucketName string) error {
	uploader := s3manager.NewUploader(sess)

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %v", filePath, err)
		}
		defer file.Close()

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(file.Name()),
			Body:   file,
		})
		if err != nil {
			return fmt.Errorf("failed to upload file %s: %v", filePath, err)
		}

		fmt.Printf("Uploaded file: %s\n", filePath)
	}

	return nil
}

// func invalidateCloudFrontDistribution(sess *session.Session, distributionID string, uploadedFiles []string) error {: This line defines a function named invalidateCloudFrontDistribution with the following parameters:
//sess *session.Session: A pointer to an AWS session object, representing the session used for communication with AWS services.
//distributionID string: The ID of the CloudFront distribution to invalidate.
//uploadedFiles []string: A slice containing the paths of the files that were uploaded to S3 and are associated with the CloudFront distribution.

// svc := s3.New(sess): This line creates a new S3 service client using the provided AWS session (sess). The s3.New() function initializes a new S3 service client.

// invalidationPaths:= These lines create a list of invalidation paths for the CloudFront distribution. It iterates over each file path in the uploadedFiles slice, prepends a / to each file path, and converts it to a pointer to a string (*string). These paths will be invalidated in the CloudFront distribution.

// invalidateCloudFrontDistribution invalidates the CloudFront distribution for the uploaded files.
func invalidateCloudFrontDistribution(sess *session.Session, distributionID string, uploadedFiles []string) error {
	svc := s3.New(sess)

	invalidationPaths := make([]*string, len(uploadedFiles))
	for i, file := range uploadedFiles {
		path := fmt.Sprintf("/%s", file)
		invalidationPaths[i] = aws.String(path)
	}

	input := &s3.CreateInvalidationInput{
		DistributionId: aws.String(distributionID),
		InvalidationBatch: &s3.InvalidationBatch{
			CallerReference: aws.String("cli-upload"),
			Paths: &s3.Paths{
				Quantity: aws.Int64(int64(len(invalidationPaths))),
				Items:    invalidationPaths,
			},
		},
	}

	_, err := svc.CreateInvalidation(input)
	if err != nil {
		return fmt.Errorf("failed to invalidate CloudFront distribution: %v", err)
	}

	fmt.Println("CloudFront distribution invalidated successfully.")
	return nil
}
