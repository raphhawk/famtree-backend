package ports

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type PersonService interface {
	CreatePerson(person PersonDTO) PersonDTO
	GetPersonById(id int) PersonDTO
	GetPersonByEmail(email string) (PersonDTO, error)
	UpdatePersonName(person PersonDTO) PersonDTO
	UpdatePersonEmail(person PersonDTO) PersonDTO
	UpdatePersonDob(person PersonDTO) PersonDTO
	UpdatePersonGender(person PersonDTO) PersonDTO
	DeletePerson(person PersonDTO) PersonDTO
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
	Personel() PersonService
	GetPerson(ctx echo.Context) error
	GetPeople(ctx echo.Context) error
	SendPeople(ctx echo.Context) error
	AddPerson(ctx echo.Context) error
	UpdatePerson(ctx echo.Context) error
	RemovePerson(ctx echo.Context) error
}
