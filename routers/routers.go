package routers

import (
	"net/http"

	"github.com/slovojoe/docupToo/database"
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
var h = handlers.New(database.Db)
// Define a slice of routes to handle all of the apps routing
var Routes = []Route{
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
        HandlerFunc: h.CreateDocument,
		
	},

	// {
	// 	Name:        "updateDocument",
	// 	Method:      "PUT",
	// 	Pattern:     "/document/update",
	// 	HandlerFunc: UpdateDocument,
	// },
	// {
	// 	Name:        "deleteteDocument",
	// 	Method:      "DELETE",
	// 	Pattern:     "/delete/{id}",
	// 	HandlerFunc: DeleteDocument,
	// },
}