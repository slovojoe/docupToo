package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/slovojoe/docupToo/database"
	"github.com/slovojoe/docupToo/models"
)

//Define handlers

var Documents []models.Document
//Creating a new document
func (h handler) CreateDocument(w http.ResponseWriter, r *http.Request) {
    // get the body of the  POST request
    // unmarshal this into a new Document struct
    // append this to the Documents array.     
    reqBody, _ := ioutil.ReadAll(r.Body)
	var document models.Document 
    unmarshalled:=json.Unmarshal(reqBody, &document)


	fmt.Println(unmarshalled)
	
    if result := database.Db.Create(&document); result.Error != nil {
		fmt.Println(result.Error)
	}
    json.NewEncoder(w).Encode(document)
}
