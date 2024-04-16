package controllers

import (
	"encoding/json"
	"fmt"
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

	rules := map[string]string{
		"Age":  "required|int",
		"FirstName":     "required|string",
		"LastName":     "required|string",
		"Country":     "required|string",
		"Description":     "required|string",
	}

	err = validators.Validate(result, rules)

	if err != nil {
		fmt.Println("Validation failed:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	} else {
		fmt.Println("Validation successful!")
	}

	
	author := &models.Author{}

	firstName := result["FirstName"].(string)
	lastName := result["LastName"].(string)
	country := result["Country"].(string)
	description := result["Description"].(string)
	age := result["Age"].(int)

	authorReturn, _ := author.Create(firstName, lastName, country, description, age)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(authorReturn)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	//TODO
	// var result map[string]interface{}
	var result map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		panic(err)
	}

	/*authorID := validators.CheckIntField(result, "AuthorID")
	author := &models.Author{}

	//authorReturn, _ := author.UpdateAuthor(authorID, result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authorReturn)*/

}