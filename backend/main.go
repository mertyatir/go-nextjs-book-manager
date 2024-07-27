package main

import (
	"book-manager/handlers"
	"log"
	"net/http"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	_ "book-manager/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
    r.HandleFunc("/books", handlers.AddBook).Methods("POST")
    r.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
    r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
    r.HandleFunc("/process-url", handlers.UrlHandler).Methods("POST")

   
    r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    log.Println("Server running on port 8000")


    corsHandler := gorillaHandlers.CORS(
        gorillaHandlers.AllowedOrigins([]string{"*"}),
        gorillaHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        gorillaHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )(r)


     log.Fatal(http.ListenAndServe(":8000", corsHandler))
}
