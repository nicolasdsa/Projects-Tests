package controllers

import (
	"encoding/json"
	"net/http"
	"project/database"
	"project/models"
)

// GetAllAuthors retrieves all authors

func getAll(w http.ResponseWriter, r *http.Request) {	
	db := database.ConnectDatabase()
	result, err := db.Query("select * from authors");

	p := models.Author{}
	authors := []models.Author{}

	for result.Next(){
		var Id, Age        int
		var FirstName, LastName, Country, Description string
		err = result.Scan(&Id, &FirstName, &LastName, &Description, &Country, &Age)
		if err != nil {
			panic(err.Error())
		}
		p.Id = Id
		p.FirstName = FirstName
		p.LastName = LastName
		p.Country = Country
		p.Description = Description
		p.Age = Age
		authors = append(authors, p)
	}

	json.NewEncoder(w).Encode(p)
}