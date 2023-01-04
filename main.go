package main

import (
	"os"
	"strings"

	"example.com/fileUploadApp/awssession"
	"example.com/fileUploadApp/logging"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// getListOfFiles will accept a directory
// path and return a slice of all .csv
// files in the directory as strings
func getListOfFiles(inputDir string) []os.DirEntry {

	var listOfFiles []os.DirEntry

	// read files in the directory
	files, err := os.ReadDir(inputDir)
	if err != nil {
		logging.ErrorLogger.Println(err)
	}

	// if extension is .csv, append
	// the filename to a slice
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".csv") {
			listOfFiles = append(listOfFiles, f)
		}
	}

	return listOfFiles
}

func main() {

	// get list of csvs to be processed
	listOfFiles := getListOfFiles("./inputFiles")
	logging.InfoLogger.Printf("list of %d files has been read", len(listOfFiles))

	//setup connection to s3 bucket
	svc := s3manager.NewUploader(awssession.Sess)
	logging.InfoLogger.Println("New AWS session connection setup")

	//open file to be uploaded
	file, err := os.Open("./inputFiles/features_data_set_1.csv")
	if err != nil {
		logging.ErrorLogger.Println("Failed to open file", err)
		os.Exit(1)
	}
	defer file.Close()

	//upload file to s3
	_, err = svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awssession.Bucket),
		Key:    aws.String("jan2023test1.csv"),
		Body:   file,
	})
	if err != nil {
		logging.ErrorLogger.Println("error", err)
		os.Exit(1)
	}

	//print success message
	logging.InfoLogger.Println("file successfully uploaded to s3")
}
