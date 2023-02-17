package handlers

import (
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovojoe/docupToo/storage"
	"github.com/slovojoe/docupToo/models"
)

type User models.User

var Users []User

// Creating a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// get the body of the  POST request
	// unmarshal this into a new User struct
	// append this to the Users array.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user User
	unmarshalled := json.Unmarshal(reqBody, &user)

	fmt.Println(unmarshalled)

	if result := storage.Db.Create(&user); result.Error != nil {
		fmt.Println(result.Error)
	}
	json.NewEncoder(w).Encode(user)
}

// Fetch a single user without their documents
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["userid"]

	var singleUser User

	//Get the first user whose ID matches the provided ID
	
	if result := storage.Db.First(&singleUser, key); result.Error != nil {
		fmt.Println(result.Error)
	}

	json.NewEncoder(w).Encode(singleUser)
	fmt.Printf("Got user %s", singleUser)

}

// Get user docs
func GetUsersDocs(w http.ResponseWriter, r *http.Request) {
	var users []User
	vars := mux.Vars(r)
	key := vars["userid"]

	//Passing the key parameters at .Find(&users,key) is the equivalent of SELECT * FROM users WHERE id = key;
	if results := storage.Db.Preload("Documents").Find(&users, key); results.Error != nil {
		fmt.Println(results.Error)
	}

	fmt.Println("got users")
	json.NewEncoder(w).Encode(users)

}

// Update user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// once again, we will need to parse the path parameters
	var userToUpdate User
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &userToUpdate)
	var user User
	vars := mux.Vars(r)
	id := vars["userid"]
	result := storage.Db.First(&user, id)

	//check if there is an error getting the user
	if result.Error != nil {
		fmt.Println("An error occured while fetching the desired user")
		return
	}

	//check if the user is trying to update  username
	//If the new updated username string is not empty it means they are trying to update username
	if userToUpdate.Username != "" {

		oldUser := user.Username
		storage.Db.Save(&user)
		fmt.Printf("Changed username from %s to %s/n", oldUser, userToUpdate.Username)
	}

	//check and prevent user from updating email or id
	if userToUpdate.ID != 0 || userToUpdate.Email != "" {
		fmt.Println("Updates are not allowed for User ID or Email")
		return

	}
	json.NewEncoder(w).Encode(&userToUpdate)

}

// Deleting a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the user we
	// wish to delete
	id := vars["userid"]
	var usertodelete User
	//Such the users table for a user whose ID is same as the one we specify
	result := storage.Db.First(&usertodelete, id)

	if result.Error != nil {
		fmt.Println(result.Error)

	}
	fmt.Printf("deleting user %s/n", usertodelete.Username)
	storage.Db.Delete(&usertodelete)
	fmt.Println("User deleted successfully")

	json.NewEncoder(w).Encode(usertodelete)

}
