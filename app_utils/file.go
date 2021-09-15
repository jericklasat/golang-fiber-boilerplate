package app_utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func UploadFile(header *multipart.FileHeader) string {
	file, err := header.Open();
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer file.Close()
	
	fmt.Printf("Uploaded File: %+v\n", header.Filename)
	fmt.Printf("File Size: %+v\n", header.Size)
	fmt.Printf("MIME Header: %+v\n", header.Header)

	sec := strconv.Itoa(int(time.Now().Unix()));

	var filename string = sec + "_" + header.Filename
	
	dst, err := os.Create("public/images/" + filename)
	if err != nil {
		fmt.Println("Error creating file")
		fmt.Println(err)
	}
	defer dst.Close()
	
	if _, err := io.Copy(dst, file); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Uploaded File")
	return filename
}