package controllers

import (
	"encoding/json"
	"log"
	"sub/src/db"
	"sub/src/models"
	"sub/src/repositories"
)

func HandlePerson(data []byte) error {
	var person models.Person

	err := json.Unmarshal(data, &person)
	if err != nil {
		return err
	}

	db, err := db.Connect()
	if err != nil {
		return err
	}
	defer db.Close()

	repo := repositories.NewPersonRepository(db)
	personId, err := repo.AddPerson(person)
	if err != nil {
		return err
	}

	log.Printf("[*] Received a message. Person registered: #%d - %s", personId, person.FullName)
	return nil
}
