package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	

	//"path/filepath"

	//"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/mux"
	"github.com/slovojoe/docupToo/models"
	"github.com/slovojoe/docupToo/storage"
)

// Define handlers
type Document models.Document

type UserFile struct{
	FileName string
	Base64File string
	UserID int
}
var Documents []models.Document
const MAX_UPLOAD_SIZE = 10240 * 10240 // 10MB

//Creating a new document
func CreateDocument(w http.ResponseWriter, r *http.Request, ) {
	// get the body of the  POST request
	// unmarshal this into a new Document struct
	//receive client request body containing file details
	//Base64 file, filename, 
	reqBody, _ := ioutil.ReadAll(r.Body)
	var userFile UserFile
	unmarshalled := json.Unmarshal(reqBody, &userFile)
	fmt.Println(unmarshalled)

	// decode the base64 image
	base64Image := userFile.Base64File
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		log.Fatal("An error occured decoding the image %s",err)
		return
	}

	// Determine the content type of the received  file
	mimeType := http.DetectContentType(imageData)
	fmt.Printf("file type is %s\n",mimeType)


	//Call the upload function to upload file to s3
	key := fmt.Sprintf("images/%s", userFile.FileName)
	fmt.Printf("object key is %s\n",key)
	//The upload function returns a url to the file which we will store in the db
	fileurl:=storage.UploadImage(storage.Sess,imageData,key,mimeType)

	//Create a document using details from user file
	document:= models.Document{Name: userFile.FileName,UserID: userFile.UserID, URL: fileurl,FileKey: key}
    
	//Store the document in the database
	if result := storage.Db.Create(&document); result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(document)
}

func GetSingleDoc(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userIdKey := vars["user_id"]
	docIdKey := vars["id"]

	var userdocuments Document
	result :=storage.Db.Where("id = ? AND user_id = ?", docIdKey, userIdKey).Find(&userdocuments)

	//Get a single document based on provided user id
	// SELECT * FROM userdocuments WHERE id = docIdKey AND user_id =userIdKey;
	if  result.Error != nil {
		fmt.Println(result.Error)
	}
	
	json.NewEncoder(w).Encode(userdocuments)
	fmt.Printf("SUCCESS: Got document %s", userdocuments)
}

// Update document by id
func UpdateDocument(w http.ResponseWriter, r *http.Request) {

	// once again, we will need to parse the path parameters
	var docToUpdate Document
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &docToUpdate)
	var doc Document
	vars := mux.Vars(r)
	id := vars["docid"]
	result := storage.Db.First(&doc, id)

	//check if there is an error getting the document
	if result.Error != nil {
		fmt.Println("An error occured while fetching the desired user")
		return
	}

	//check if the user is trying to update  the document name
	//If the new updated document name string is not empty it means they are trying to update username
	if docToUpdate.Name != "" {

		oldDoc := doc.Name
		storage.Db.Save(&doc)
		fmt.Printf("Changed document name from %s to %s/n", oldDoc, docToUpdate.Name)
	}

	//check and prevent user from updating doc id or owner id
	if docToUpdate.ID >= 0 || docToUpdate.UserID >= 0 {
		fmt.Println("Updates are not allowed for User ID or Email")
		return

	}
	json.NewEncoder(w).Encode(&docToUpdate)

}

// Delete document
// Deleting a document by ID
func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	// parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the doc we
	// wish to delete
	id := vars["docid"]
	var documentToDelete Document
	//Search the documents table for a document whose ID is same as the one we specify
	result := storage.Db.First(&documentToDelete, id)

	if result.Error != nil {
		fmt.Println(result.Error)

	}
	fmt.Printf("deleting document %s/n", documentToDelete.Name)
	storage.Db.Delete(&documentToDelete)
	fmt.Println("Document deleted successfully")

	json.NewEncoder(w).Encode(documentToDelete)

}
