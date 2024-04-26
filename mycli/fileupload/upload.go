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