package routes

import (
	"net/http"
	"project/controllers"
	"project/middlewares"
)

func CarregaRotas() {
	http.HandleFunc("/insert", func(w http.ResponseWriter, r *http.Request) {
    middlewares.MethodNotAllowedHandler(http.HandlerFunc(controllers.CreateAuthor)).ServeHTTP(w, r)
})
	http.HandleFunc("/getAll", controllers.GetAllAuthors)
}