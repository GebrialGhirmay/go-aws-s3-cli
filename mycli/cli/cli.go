package cli

import (
    "go-aws-s3-cli/mycli/fileupload"
    "go-aws-s3-cli/mycli/logging" // Imports the logging package
    "os"

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
            logging.Warning("File path is required. Please provide a file path using the --file flag.")
            return
        }

        err := fileupload.UploadFile(filePath)
        if err != nil {
            logging.Error("Error uploading file: %v", err)
            return
        }

        logging.Info("File uploaded successfully")
    },
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&filePath, "file", "f", "", "Path to the file you want to upload")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        logging.Error("Error executing CLI: %v", err)
        os.Exit(1)
    }
}