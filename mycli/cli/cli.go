package cli

import (
    "fmt"
    "os"

    "go-aws-s3-cli/mycli/fileupload"

    "github.com/spf13/cobra"
)

var (
    filePath string
)

var rootCmd = &cobra.Command{
    Use:   "aws-s3-cli",
    Short: "A command-line interface for interacting with AWS S3",
    Run: func(cmd *cobra.Command, args []string) {
        if filePath == "" {
            fmt.Println("Error: File path is required. Please provide a file path using the --file flag.")
            return
        }

        err := fileupload.UploadFile(filePath)
        if err != nil {
            fmt.Println("Error uploading file:", err)
            return
        }

        fmt.Println("File uploaded successfully")
    },
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Path to the file you want to upload")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}