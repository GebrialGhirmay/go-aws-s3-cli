// A cli.go file will orchestrate the interaction between different modules. Would need to define the command-line interface using a library like cobra and invoke relevant functions from other modules based on user input.

package cli

import (
    "fmt"
    "os"

    "go-aws-s3-cli/mycli/fileupload"

    "github.com/spf13/cobra" //imports the fileupload package and the github.com/spf13/cobra library for creating the command-line interface (to use this CLI, this package will need to be installed - followed by the command go mod tidy).
)

var (
    filePath string //A filePath variable is declared to store the file path provided by the user.
)

var rootCmd = &cobra.Command{ //A rootCmd is defined using cobra.Command. This is the main command for the CLI.
    Use:   "aws-s3-cli",
    Short: "A command-line interface for interacting with AWS S3",
    Run: func(cmd *cobra.Command, args []string) { //The rootCmd.Run function calls the fileupload.UploadFile function, passing the filePath provided by the user.
        err := fileupload.UploadFile(filePath)
        if err != nil {
            fmt.Println("Error uploading file:", err)
            return
        }

        fmt.Println("File uploaded successfully")
    },
}

func init() { //In the init function, a persistent flag (--file or -f) is defined to allow the user to specify the file path.
    rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Path to the file you want to upload")
}

func Execute() { //The Execute function is called to start the CLI and execute the rootCmd.
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
