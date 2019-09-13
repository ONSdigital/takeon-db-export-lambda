package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var region = os.Getenv("AWS_REGION")

func main() {

	lambda.Start(handle)

}

func handle() {
	fmt.Println("Starting the application...")
	var gqlEndpoint = os.Getenv("GRAPHQL_ENDPOINT")
	var dataExport = ""
	response, err := http.Get(gqlEndpoint)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		dataExport = string(data)
		// fmt.Println(response)
	}

	// cdbExport := make(chan string)

	// callGraphqlEndpoint(cdbExport)

	// dbExport := <- cdbExport

	var bucketFilenamePrefix = "takeon-data-export-"

	currentTime := time.Now().Format("2019-09-11-00:00:00")
	// fmt.Println(t.Format("20060102150405"))

	fmt.Printf("Region: %q\n", region)
	config := &aws.Config{
		Region: aws.String(region),
	}

	sess := session.New(config)

	uploader := s3manager.NewUploader(sess)

	bucket := os.Getenv("S3_BUCKET")
	filename := strings.Join([]string{bucketFilenamePrefix, currentTime}, "")
	fmt.Printf("Bucket filename: %q\n", filename)

	reader := strings.NewReader(dataExport)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   reader,
	})

	if err != nil {
		fmt.Printf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to s3 bucket %q\n", filename, bucket)
}

func callGraphqlEndpoint(cdbExport chan string) {
	var gqlEndpoint = os.Getenv("GRAPHQL_ENDPOINT")
	response, err := http.Get(gqlEndpoint)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		cdbExport <- string(data)
		// fmt.Println(response)
	}
}
