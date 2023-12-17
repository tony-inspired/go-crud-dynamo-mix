package main

import (
	"net/http"
	"github.com/go-chi/chi"	
	HttpStatusx "go-crud-dynamo-mix/utils/http"
)

func main() {
	port := ":8080"

	var router *chi.Mux

	router = chi.NewRouter() //2 to handle requests on incoming connections

	router.Route("/product", func(r chi.Router){ //3
		r.Get("/", GetProductHandler)//4
		r.Get("/{ID}", GetProductHandler)
	})

	http.ListenAndServe(port, router) //1 
}

func GetProductHandler(w http.ResponseWriter, r *http.Request){
	myResponseData := "myResponseData"

	if chi.URLParam(r, "ID") != ""{
		GetOneProduct(w, r, myResponseData)
	} else {
		GetAllProducts(w, r, myResponseData)
	}
}

func GetOneProduct(w http.ResponseWriter, r *http.Request, myResponseData interface{}){
	HttpStatusx.StatusOk(w, r, myResponseData)
}

func GetAllProducts(w http.ResponseWriter, r *http.Request, myResponseData interface{}){
	HttpStatusx.StatusOk(w, r, myResponseData)
}