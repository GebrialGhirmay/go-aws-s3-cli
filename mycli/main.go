package main

import (
	"fmt"

	"go-aws-s3-cli/mycli/fileupload"
)

func main() {
	err := fileupload.UploadFile("C:/Users/Gebrial Ghirmay/Desktop/Cached pages/Azerion H&M 16 4 2024/3500.html")
	if err != nil {
		fmt.Println("Error uploading file:", err)
		return
	}

	fmt.Println("File uploaded successfully")
}