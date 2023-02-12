package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/slovojoe/docupToo/constants"
	"github.com/slovojoe/docupToo/database"
	"github.com/slovojoe/docupToo/models"
	"github.com/slovojoe/docupToo/routers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Adding dummy user and document data
var(
	documents = []models.Document{
	{Name: "Drivers License", Body: "DL", UserID: 1},
}
	users=[]models.User{
	{Username: "Kratos", Email: "Kratos@gmail.com", Password: "Keratosis"},
})




//Loop through the specified routes
func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routers.Routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}

func main() {
	database.ConnectDB()

	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	err := http.ListenAndServe(constants.CONN_HOST+":"+constants.CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}



    for i := range documents{db.Create(&documents[i])}
	for i := range users{db.Create(&users[i])}

}
