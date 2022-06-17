package handlers

import (
	"encoding/json"
	"fmt"
	"go-crud-article/connection"
	"go-crud-article/helpers"
	"go-crud-article/structs"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func CreateArticle(w http.ResponseWriter, r *http.Request)  {
	var article structs.Posts
	article.Created_date = time.Now().UTC()
	article.Updated_date = time.Now().UTC()

	errval := helpers.ValidatePayloadsArticle(&article, r)

	if len(errval) != 0 {
		json.NewEncoder(w).Encode(structs.Result{Code: 400, Data: errval, Message: "Bad Request"})
		return
	} 
	
	if err := connection.DB.Create(&article).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Data: article, Message: "Success Create Article"})
}

func GetArticles(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]
	
	articles := []structs.Posts{}

	if err := connection.DB.Limit(limit).Offset(offset).Order("updated_date").Find(&articles).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 
	
	json.NewEncoder(w).Encode(structs.Result{Code: 200, Data: articles, Message: "Success Get Articles"})
}

func GetArticle(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.Posts

	if err := connection.DB.First(&article, articleID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 
	
	json.NewEncoder(w).Encode(structs.Result{Code: 200, Data: article, Message: "Success Get Article"})
}

func UpdateArticle(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var articleUpdates structs.Posts
	articleUpdates.Updated_date = time.Now().UTC()

	var article structs.Posts

	errval := helpers.ValidatePayloadsArticle(&articleUpdates, r)

	if len(errval) != 0 {
		json.NewEncoder(w).Encode(structs.Result{Code: 400, Data: errval, Message: "Bad Request"})
		return
	} 

	if err := connection.DB.First(&article, articleID).Error; err != nil {		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	if err := connection.DB.Model(&article).Update(&articleUpdates).Error; err != nil {		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 

	json.NewEncoder(w).Encode(structs.Result{Code: 200, Data: articleUpdates, Message: "Success Update Articles"})
}

func DeleteArticle(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	articleID := vars["id"]

	var article structs.Posts

	if err := connection.DB.First(&article, articleID).Error; err != nil {		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 
	
	if err := connection.DB.Delete(&article).Error; err != nil {		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} 
	
	json.NewEncoder(w).Encode(structs.Result{Code: 200, Message: "Success Delete Article"})
}