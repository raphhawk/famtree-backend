package ports

import (
	"time"

	artErr "github.com/raphhawk/famtree/cmd/common/error"
)

// controller contracts
var (
	Port = ":8080"

	//GET
	GetPerson = "/api/person/:id"
	GetPeople = "/api/people"

	//POST
	AddPerson    = "/api/person"
	AddPeople    = "/api/people"
	UpdatePerson = "/api/person/update/:id"
	RemovePerson = "/api/person/:id"
)

// PersonDTO acts as data transfer object for Person
type PersonDTO struct {
	ID        int                     `json:"id"`
	FirstName string                  `json:"f_name"`
	LastName  string                  `json:"l_name"`
	Gender    string                  `json:"gender"`
	Age       int                     `json:"age"`
	Dob       string                  `json:"dob"`
	Email     string                  `json:"email"`
	Info      artErr.ArtificialErrors `json:"info"`
}

// PersonDAO acts as data access object for Person
type PersonDAO struct {
	PersonId  int       `json:"p_id"`
	Name      string    `json:"name"`
	Dob       string    `json:"dob"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"c_at"`
	UpdatedAt time.Time `json:"u_at"`
}
