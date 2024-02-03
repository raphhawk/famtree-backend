package main

/*
	action plan
	create person service, store and get person data
	create family service, store and get a set of people
	create tree service, store and get trees of family
*/

import (
	"log"

	"github.com/raphhawk/famtree/internal/person/ports"
	"github.com/raphhawk/famtree/internal/person/server"
)

func main() {
	var ps ports.PersonServer
	ps = server.NewServer(ports.Port)
	log.Println(
		ps.
			Personel().
			GetPersonById(1),
	)

	dto := ports.PersonDTO{
		FirstName: "Django Master",
		LastName:  "Unchainer",
		Dob:       "18401220",
		Gender:    "M",
		Email:     "django@hitoribushi.com",
	}
	log.Println(ps.Personel().CreatePerson(dto))
	dto2 := ports.PersonDTO{
		ID:        2,
		FirstName: "Django Master",
		LastName:  "Unchainer",
		Gender:    "F",
	}
	log.Println(ps.Personel().UpdatePersonGender(dto2))
	log.Println(ps.Personel().DeletePerson(dto2))

	ps.ListenAndServe()
}
