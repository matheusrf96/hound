package models

type Person struct {
	FullName  string `json:"name"`
	Email     string `json:"email"`
	BirthDate string `json:"birth_date"`
}
