package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(make(map[string]interface{}))
		}

		next.ServeHTTP(w, r)
	})
}

func HandleReq() {
	log.Println("Start development server localhost:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", HomePage)
	myRouter.Handle("/article", MiddlewareAuth(http.HandlerFunc(CreateArticle))).Methods("OPTIONS", "POST")
	myRouter.Handle("/article/{limit}/{offset}", MiddlewareAuth(http.HandlerFunc(GetArticles))).Methods("OPTIONS", "GET")
	myRouter.Handle("/article/{id}", MiddlewareAuth(http.HandlerFunc(GetArticle))).Methods("OPTIONS", "GET")
	myRouter.Handle("/article/{id}", MiddlewareAuth(http.HandlerFunc(UpdateArticle))).Methods("OPTIONS", "PUT")
	myRouter.Handle("/article/{id}", MiddlewareAuth(http.HandlerFunc(DeleteArticle))).Methods("OPTIONS", "Delete")

	handler := cors.AllowAll().Handler(myRouter)
	
	log.Fatal(http.ListenAndServe(":9999", handler))
}