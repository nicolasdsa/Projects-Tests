package models

import (
	"fmt"
	"project/database"
	"strings"
)

type Author struct {
	Id				 int `json:"Id"`
	Age        int `json:"Age"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Country string `json:"Country"`
	Description string `json:"Description"`
}

func (a *Author) Create(firstName, lastName, country, description string, age int) (*Author, error) {
	db := database.ConnectDatabase()

	insertStmt, err := db.Prepare("INSERT INTO authors(FirstName, LastName, Country, Description, Age) VALUES($1, $2, $3, $4, $5) RETURNING Id")
	if err != nil {
		return nil, err
	}

	var id int
	err = insertStmt.QueryRow(firstName, lastName, country, description, age).Scan(&id)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return &Author{
		Id:          id,
		FirstName:   firstName,
		LastName:    lastName,
		Country:     country,
		Description: description,
		Age:         age,
	}, nil

}

func (a *Author) GetAll() ([]Author) {
	db := database.ConnectDatabase()
	result, err := db.Query("select Id, FirstName, LastName, Description, Country, Age from authors")

	if err != nil {
		panic(err)
	}

	p := Author{}
	authors := []Author{}

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

	return authors
}


func (a *Author) UpdateAuthor(authorID int, result map[string]interface{}) (*Author, error) {
	db := database.ConnectDatabase()
	updateStmt := "UPDATE authors SET"
	var params []interface{}
	i := 2 	
	for key, value := range result {
		updateStmt += fmt.Sprintf(" %s = $%d,", key, i)
		params = append(params, value)
		i++
	}
	updateStmt = strings.TrimSuffix(updateStmt, ",") + " WHERE Id = $1"
	params = append([]interface{}{authorID}, params...)
	stmt, err := db.Prepare(updateStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(params...)
	if err != nil {
		return nil, err
	}

	// Retrieve the updated author
	updatedAuthor := &Author{}
	err = db.QueryRow("SELECT Id, FirstName, LastName, Description, Country, Age FROM authors WHERE Id = $1", authorID).Scan(
		&updatedAuthor.Id,
		&updatedAuthor.FirstName,
		&updatedAuthor.LastName,
		&updatedAuthor.Description,
		&updatedAuthor.Country,
		&updatedAuthor.Age,
	)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return updatedAuthor, nil
}

