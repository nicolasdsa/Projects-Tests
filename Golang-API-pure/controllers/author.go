package controllers

import (
	"encoding/json"
	"net/http"
	"project/models"
	"project/validators"
)

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {	

	author := &models.Author{}

	authorsReturn  := author.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authorsReturn)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {

	var result map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	firstName :=  validators.CheckStringField(result, "FirstName");
	lastName :=  validators.CheckStringField(result, "LastName");
	country :=  validators.CheckStringField(result, "Country");
	description := validators.CheckStringField(result, "Description");
	age := validators.CheckIntField(result, "Age")

	author := &models.Author{}

	authorReturn, _  := author.Create(firstName, lastName, country, description, age)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(authorReturn)
}