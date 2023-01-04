package awssession

import (
	"os"

	_ "example.com/fileUploadApp/configparser"
	"example.com/fileUploadApp/logging"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type awscreds struct {
	s3KeyId     string
	s3SecretKey string
}

var Sess *session.Session

// getEnvVars will load the environment
// variables from .env file for
// setting up aws credentials
func getEnvVars() {

	err := godotenv.Load(".env")
	if err != nil {
		logging.ErrorLogger.Printf("Failed to load environment variables : %s", err)
		os.Exit(1)
	}
	logging.InfoLogger.Printf("Environment variables loaded successfully!")

}

// getAWSCredentials will return the
// AWS access and secret key
func getAWSCredentials() awscreds {

	//Load the environment variables
	getEnvVars()

	//capture the environment variables
	awsCredentials := awscreds{s3KeyId: os.Getenv("AWS_ACCESS_KEY"), s3SecretKey: os.Getenv("AWS_SECRET_KEY")}

	return awsCredentials

}

func init() {

	awsCredentials := getAWSCredentials()

	conf := aws.Config{
		Region:      aws.String(viper.GetString("awss3config.s3region")),
		Credentials: credentials.NewStaticCredentials(awsCredentials.s3KeyId, awsCredentials.s3SecretKey, ""),
	}
	sess, err := session.NewSession(&conf)
	if err != nil {
		logging.ErrorLogger.Println("Failed to establish AWS session", err)
		os.Exit(1)
	}
	logging.InfoLogger.Println("AWS session created successfully!")
	Sess = sess

}
