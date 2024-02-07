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
	FirstName string                  `json:"f_name," validate:"required,min=5,max=50,alpha"`
	LastName  string                  `json:"l_name" validate:"required,min=5,max=50,alpha"`
	Gender    string                  `json:"gender" validate:"required,alpha,uppercase,containsany=MFO,len=1"`
	Age       int                     `json:"age"`
	Dob       string                  `json:"dob" validate:"dateonly"`
	Email     string                  `json:"email" validate:"email"`
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
