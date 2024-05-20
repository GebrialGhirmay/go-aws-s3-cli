package main

import (
    "go-aws-s3-cli/mycli/cli"
    "go-aws-s3-cli/mycli/logging" // Imports the logging package
)

func main() {
    logging.SetLogLevel(0) // Set the log level to debugLevel (0)
    cli.Execute()
}