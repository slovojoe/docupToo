package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovojoe/docupToo/database"
	"github.com/slovojoe/docupToo/models"
)

// Define handlers
type Document models.Document

var Documents []models.Document

// Creating a new document
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	// get the body of the  POST request
	// unmarshal this into a new Document struct
	// append this to the Documents array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var document models.Document
	unmarshalled := json.Unmarshal(reqBody, &document)

	fmt.Println(unmarshalled)

	if result := database.Db.Create(&document); result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(document)
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
	result := database.Db.First(&doc, id)

	//check if there is an error getting the document
	if result.Error != nil {
		fmt.Println("An error occured while fetching the desired user")
		return
	}

	//check if the user is trying to update  the document name
	//If the new updated document name string is not empty it means they are trying to update username
	if docToUpdate.Name != "" {

		oldDoc := doc.Name
		database.Db.Save(&doc)
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
	result := database.Db.First(&documentToDelete, id)

	if result.Error != nil {
		fmt.Println(result.Error)

	}
	fmt.Printf("deleting document %s/n", documentToDelete.Name)
	database.Db.Delete(&documentToDelete)
	fmt.Println("Document deleted successfully")

	json.NewEncoder(w).Encode(documentToDelete)

}
