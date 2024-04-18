package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/models"
	"project/validators"
	"strconv"
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
	
	var result map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	idStr := r.URL.Path[len("/update/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Invalid ID",
		})
		return
	}

	resultRules := result
	resultRules["id"] = id

	rules := map[string]string{
		"Age":        "nullable|int",
		"FirstName":  "nullable|string",
		"LastName":   "nullable|string",
		"Country":    "nullable|string",
		"Description": "nullable|string",
		"id":         "required|int",
	}

	err = validators.Validate(resultRules, rules)

	if err != nil {
		fmt.Println("Validation failed:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	author := &models.Author{}

	authorReturn, err := author.UpdateAuthor(id, result)
	if err != nil {
		fmt.Println("Update failed:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authorReturn)
}