package routes

import (
	"net/http"
	"project/controllers"
	"project/middlewares"
)

func CarregaRotas() {
	http.HandleFunc("/update/", func(w http.ResponseWriter, r *http.Request) {
    middlewares.MethodNotAllowedHandler(http.HandlerFunc(controllers.UpdateAuthor)).ServeHTTP(w, r)
})
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
    middlewares.MethodNotAllowedHandler(http.HandlerFunc(controllers.GetAllAuthors)).ServeHTTP(w, r)
})
http.HandleFunc("/getAll", func(w http.ResponseWriter, r *http.Request) {
	middlewares.MethodNotAllowedHandler(http.HandlerFunc(controllers.GetAllAuthors)).ServeHTTP(w, r)
})
}