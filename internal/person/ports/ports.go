package ports

import "database/sql"

type PersonService interface {
	CreatePerson(person PersonDTO) PersonDTO
	GetPersonById(id int) PersonDTO
	GetPersonByEmail(email string) (PersonDTO, error)
	UpdatePersonName(person PersonDTO) PersonDTO
	UpdatePersonEmail(person PersonDTO) PersonDTO
	UpdatePersonDob(person PersonDTO) PersonDTO
	UpdatePersonGender(person PersonDTO) PersonDTO
	DeletePerson(person PersonDTO) PersonDTO
	SetFamily()
}

type PersonDatabase interface {
	OpenDB(dsn string) (*sql.DB, error)
	ConnectToDB() *sql.DB
	InitDB() error
	GetDB() *sql.DB
}

type PersonServer interface {
	MapUrls() error
	ListenAndServe() error
}
