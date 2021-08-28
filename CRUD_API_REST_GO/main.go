package main

import (
	"log"
	"net/http"

	"github.com/alejandrosubero/crud-api/commons"
	"github.com/alejandrosubero/crud-api/routes"
	"github.com/gorilla/mux"
)

func main() {

	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)
	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("ejecutansose servidor en el puerto 9000")
	log.Panicln(server.ListenAndServe())
}

/*
package main

import (
    "log"
    "fmt"
    "net/http"
    "path/filepath"
)

// temporary directory location
var tmpDir = filepath.FromSlash("/Users/Uday.Hiwarale/tmp/")

func main() {

    // default route
    http.HandleFunc( "/", func( res http.ResponseWriter, req *http.Request ) {
        fmt.Fprint( res, "Hello Golang!" )
    } )

    // return a `.html` file for `/index.html` route
    http.HandleFunc( "/index.html", func( res http.ResponseWriter, req *http.Request ) {
        http.ServeFile( res, req, filepath.Join( tmpDir, "/index.html" ) );
    } )

    // start HTTP server with `http.DefaultServeMux` handler
    log.Fatal(http.ListenAndServe( ":9000", nil ))

}

*/
