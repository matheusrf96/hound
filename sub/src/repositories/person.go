package repositories

import (
	"database/sql"
	"sub/src/models"
)

type Person struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *Person {
	return &Person{db}
}

func (repo Person) AddPerson(person models.Person) (uint32, error) {
	var personId uint32

	statement, err := repo.db.Prepare(`
		INSERT INTO person (
			full_name
			, email
			, birth_date
		)
		VALUES ($1, $2, $3)
		RETURNING id
	`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	err = statement.QueryRow(
		person.FullName,
		person.Email,
		person.BirthDate,
	).Scan(&personId)
	if err != nil {
		return 0, err
	}

	return personId, nil
}
