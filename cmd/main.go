package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"encoding/json"
)

func main() {
	port := ":8080"

	var router *chi.Mux

	router = chi.NewRouter() //2

	router.Route("/product", func(r chi.Router){ //3
		r.Get("/", GetProductHandler)//4
	})

	http.ListenAndServe(port, router) //1
}

func GetProductHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Accept-Control-Allow-Origin", "*")

	w.WriteHeader(http.StatusOK)

	var data, _ = json.Marshal("qwe")

	w.Write(data)
}

type response struct {
	Status int `json:"status"`
	Result interface{} `json:"result"`
}