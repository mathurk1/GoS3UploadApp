package main

import (
	"os"
	"path/filepath"
	"strings"

	"example.com/fileUploadApp/awssession"
	_ "example.com/fileUploadApp/configparser"
	"example.com/fileUploadApp/logging"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

// getListOfFiles will accept a directory
// path and return a slice of all .csv
// files in the directory as strings
func getListOfFiles(inputDir string) []os.DirEntry {

	// create an empty slice
	var listOfFiles []os.DirEntry

	// read files in the directory
	files, err := os.ReadDir(inputDir)
	if err != nil {
		logging.ErrorLogger.Println(err)
		os.Exit(1)
	}

	// if extension is .csv, append
	// the filename to slice
	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".csv") {
			listOfFiles = append(listOfFiles, f)
		}
	}

	return listOfFiles
}

// openFile will open the passed filename
// and return the filebytes
func openFile(f string) *os.File {

	file, err := os.Open(f)
	if err != nil {
		logging.ErrorLogger.Println("Failed to open file", err)
		os.Exit(1)
	}

	return file

}

// uploadFile takes filebytes and uploads to the
// specified S3 bucket
func uploadFile(f *os.File, svc *s3manager.Uploader, fn string) {

	//upload file to s3
	_, err := svc.Upload(&s3manager.UploadInput{
		Bucket: aws.String(viper.GetString("awss3config.s3bucket")),
		Key:    aws.String(fn),
		Body:   f,
	})
	if err != nil {
		logging.ErrorLogger.Println("error", err)
		os.Exit(1)
	}
	logging.InfoLogger.Printf("file successfully uploaded %s to s3", fn)

}

// createS3FilePath will return the location
// to which the csv file will be uploaded
func createS3FilePath(f string) string {

	var dp string
	paths := viper.GetStringSlice("fileconfig.s3directory")
	if len(paths) != 0 {
		paths = append(paths, f)
		dp = strings.Join(paths, "/")
	} else {
		dp = f
	}

	return dp

}

func main() {

	// get list of csvs to be processed
	inputDir := viper.GetString("fileconfig.inputdirectory")
	listOfFiles := getListOfFiles(inputDir)
	logging.InfoLogger.Printf("list of %d files has been read", len(listOfFiles))

	// check if csv files are present to be uploaded
	if len(listOfFiles) == 0 {
		logging.InfoLogger.Println("No files to process, exiting program!")
		os.Exit(0)
	}

	// setup connection to s3 bucket
	svc := s3manager.NewUploader(awssession.Sess)
	logging.InfoLogger.Println("New AWS session connection setup")

	// process files and uploade to s3
	for _, file := range listOfFiles {
		logging.InfoLogger.Printf("processing file %s\n", file.Name())
		s3UploadPath := createS3FilePath(file.Name())
		f := openFile(filepath.Join(inputDir, file.Name()))
		uploadFile(f, svc, s3UploadPath)
		f.Close()
	}

}
