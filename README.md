## GO S3 Upload App

This is my first Go lang project. This application uploads a set of csv files from a local directory to a specifiec S3 bucket

So far, the main concepts explored in this Go project is:
- creating, using and reusing Go packages and modules
- setting up logging at a project level
- parsing config files using viper

To run the program:
1. update the `config.yaml` file with required information
2. create a `.env` file at the root of the project with the S3 Access and Secret keys
3. simply run `go run main.go` from the terminal
