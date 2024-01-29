package main

/*
	action plan
	create person service, store and get person data
	create family service, store and get a set of people
	create tree service, store and get trees of family
*/

import (
	"log"

	"github.com/raphhawk/famtree/internal/person/app"
	"github.com/raphhawk/famtree/internal/person/ports"
)

func main() {
	var personServ ports.PersonService
	personServ, err := app.New()
	if err != nil {
		log.Println(err)
	}
	log.Println(personServ.GetPersonById(1))
	dto := ports.PersonDTO{
		FirstName: "Django Master",
		LastName:  "Unchainer",
		Dob:       "18401220",
		Gender:    "M",
		Email:     "django@hitoribushi.com",
	}
	log.Println(personServ.CreatePerson(dto))
	dto2 := ports.PersonDTO{
		ID:        2,
		FirstName: "Django Master",
		LastName:  "Unchainer",
		Gender:    "F",
	}
	log.Println(personServ.UpdatePersonGender(dto2))
	log.Println(personServ.DeletePerson(dto2))
}
