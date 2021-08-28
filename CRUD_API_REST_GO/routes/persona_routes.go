package routes

import (
	"github.com/alejandrosubero/crud-api/controllers"
	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/persona/api").Subrouter()
	subRoute.HandleFunc("/all", controllers.GetALL).Methods("GEt")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete/{id}", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GEt")
	subRoute.HandleFunc("/", controllers.Base).Methods("GEt")
	subRoute.HandleFunc("/dowload/{nombre}", controllers.FileDowload).Methods("GEt")
}
