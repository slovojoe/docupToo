package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/slovojoe/docupToo/database"
	"github.com/slovojoe/docupToo/handlers"
)

// Create a Route struct defining all the parameters a route should have
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//type Routes []Route
//var h = handlers.New(database.Db)
// Define a slice of routes to handle all of the apps routing
var Routes = []Route{

	///Users
	{
		Name:        "createUser",
		Method:      "POST",
		Pattern:     "/users/create",
        HandlerFunc: handlers.CreateUser,
		
	},

	//Gets only a user without their documents
	{
		Name:        "getUser",
		Method:      "GET",
		Pattern:     "/users/get/{userid}",
        HandlerFunc: handlers.GetUser,
		
	},

	//Returns a user with all their documents
	{
		Name:        "getUserDocs",
		Method:      "GET",
		Pattern:     "/users/getuserdocs/{userid}",
        HandlerFunc: handlers.GetUsersDocs,
		
	},
	{
		Name:        "updateUser",
		Method:      "PUT",
		Pattern:     "/users/update/{userid}",
        HandlerFunc: handlers.UpdateUser,
		
	},
	{
		Name:        "deleteUser",
		Method:      "DELETE",
		Pattern:     "/users/delete/{userid}",
        HandlerFunc: handlers.DeleteUser,
		
	},

	//Documents

	// {Name: "getDoc",
	// 	Method:  "GET",
	// 	Pattern: "/doc/{id}",

	// 	HandlerFunc: getDoc,
	// },
	// {
	// 	Name:        "getDocuments",
	// 	Method:      "GET",
	// 	Pattern:     "/documents/{userID}",
	// 	HandlerFunc: getDocuments,
	// },
	{
		Name:        "createDocument",
		Method:      "POST",
		Pattern:     "/document/create",
        HandlerFunc: handlers.CreateDocument,
		
	},

	{
		Name:        "updateDocument",
		Method:      "PUT",
		Pattern:     "/documents/update/{docid}",
		HandlerFunc: handlers.UpdateDocument,
	},
	{
		Name:        "deleteteDocument",
		Method:      "DELETE",
		Pattern:     "/documents/delete/{docid}",
		HandlerFunc: handlers.DeleteDocument,
	},
}

//Loop through the specified routes
func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range Routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}