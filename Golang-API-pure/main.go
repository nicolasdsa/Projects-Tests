package main

import (
	"net/http"
	"project/routes"
)

func main(){
	routes.CarregaRotas();
	http.ListenAndServe(":8000", nil);
}