package main

import (
	//"context"
	//"bytes"
	"fmt"
	"log"
	"net/http"

	//	"os"

	// "net/http"

	//	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/gorilla/mux"

	//"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/s3"
	//	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/slovojoe/docupToo/constants"
	"github.com/slovojoe/docupToo/routers"
	"github.com/slovojoe/docupToo/storage"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/s3"
	// "github.com/gorilla/mux"
	// "github.com/slovojoe/docupToo/constants"
	// "github.com/slovojoe/docupToo/models"
	// "github.com/slovojoe/docupToo/routers"
	// "github.com/slovojoe/docupToo/storage"
)

//Adding dummy user and document data
// var(
// 	documents = []models.Document{
// 	{Name: "Drivers License", Body: "DL", UserID: 1},
// }
// 	users=[]models.User{
// 	{Username: "Kratos", Email: "Kratos@gmail.com", Password: "Keratosis"},
// })

//This function lists all items stored in a aws s3 bucket
func ListImages(sess *session.Session){
	service:=s3.New(sess)
	response, err:=service.ListObjectsV2(
		&s3.ListObjectsV2Input{Bucket: aws.String(constants.AWS_BUCKET_NAME)},

	)
	if err!=nil{
		log.Printf("AN error occured fetching items: %s",err)
	}

	//If there is no error loop through the response
	for _,item:=range response.Contents{
		log.Printf("Name : %s\n",*item.Key)
		

	}


}

//Downloading items from s3 bucket
func DownloadItems(){
	
}

func main() {
	//start a new aws session by calling the create session func
	storage.CreateAWSSession()
	storage.ConnectDB()

	muxRouter := mux.NewRouter().StrictSlash(true)
	router := routers.AddRoutes(muxRouter)
	err := http.ListenAndServe(constants.CONN_HOST+":"+constants.CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}

	// for i := range documents{database.Db.Create(&documents[i])}
	// for i := range users{database.Db.Create(&users[i])}


	
	

	

	//Create a new s3 service instance
	service := s3.New(storage.Sess)
	//List all buckets available in the aws account
	result, err := service.ListBuckets(nil)
	if err != nil {
		fmt.Printf("Error listing buckets %s", err)
	}

	for _, bucket := range result.Buckets {
		log.Printf("Bucket : %s\n", aws.StringValue(bucket.Name))
	}

	//Item upload
	//UploadImage(sess)
	ListImages(
		storage.Sess)

}
