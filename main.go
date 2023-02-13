package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slovojoe/docupToo/constants"
	"github.com/slovojoe/docupToo/database"
	"github.com/slovojoe/docupToo/models"
	"github.com/slovojoe/docupToo/routers"
)

//Adding dummy user and document data
var(
	documents = []models.Document{
	{Name: "Drivers License", Body: "DL", UserID: 1},
}
	users=[]models.User{
	{Username: "Kratos", Email: "Kratos@gmail.com", Password: "Keratosis"},
})






func main() {
	database.ConnectDB()

	muxRouter := mux.NewRouter().StrictSlash(true)
	router := routers.AddRoutes(muxRouter)
	err := http.ListenAndServe(constants.CONN_HOST+":"+constants.CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}



    for i := range documents{database.Db.Create(&documents[i])}
	for i := range users{database.Db.Create(&users[i])}

}
