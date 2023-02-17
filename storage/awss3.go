package storage

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/slovojoe/docupToo/constants"
)

//Create a global aws session variable to be used across the app
var Sess *session.Session
var SessErr error
//Creating a new AWS session
func CreateAWSSession()*session.Session{
    accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
    Sess, SessErr = session.NewSession(&aws.Config{
        Region:      aws.String("us-west-2"),
        Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
    })
    if err != nil {
        log.Fatal(err.Error())
        return nil
    }
	log.Println("AWS S3 session created successfully")
    return Sess
}

//This function uploads images to aws s3 bucket
func UploadImage(sess *session.Session, fileBytes []byte, fileKey string,contentType string) (string) {
	//Use the OS package to open an image file stored in the project assets directory
	///TODO: Implement asset selection from local storage


	//define an s3 upload manager passing the session parameter to it
	uploader := s3manager.NewUploader(sess)
	result, uploaderr := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String(constants.AWS_BUCKET_NAME),
		Key:    aws.String(fileKey),
		Body:   bytes.NewReader(fileBytes),
		ContentType:aws.String(contentType) ,
	})
	//Check for errors uploading to s3
	if uploaderr != nil {
		log.Fatal("An upload error occured %s", uploaderr.Error())
	}

	//If no error
	log.Printf("Upload successful %s\n", result)
	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", constants.AWS_BUCKET_NAME, fileKey)
	return url


}