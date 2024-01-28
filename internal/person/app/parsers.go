package app

import (
	"strings"
	"time"

	"github.com/raphhawk/famtree/internal/person/ports"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func decodeName(str string) (
	name []string, gender string,
) {
	name = strings.Split(str, " ")
	name = name[:len(name)-1]
	for i := range name {
		name[i] = cases.Title(
			language.English,
			cases.Compact,
		).String(name[i])
	}
	gender = string(str[len(str)-1])
	return name, gender
}

func encodeName(name []string, gender string) string {
	n := ""
	for i, s := range name {
		if i != 0 {
			n += " "
		}
		n += s
	}
	n += " " + gender
	return n
}

func decodeTime(t string) (time.Time, error) {
	parsedDate, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return time.Time{}, err
	}
	return parsedDate, nil
}

func getAge(dob string) int {
	dobTime, _ := decodeTime(dob)
	return int(
		time.Now().Sub(dobTime).Hours() / 24 / 365,
	)
}

func dtoTodao(dto ports.PersonDTO) ports.PersonDAO {
	var dao ports.PersonDAO
	dao.PersonId = dto.ID
	dao.Name = encodeName(
		[]string{dto.FirstName, dto.LastName},
		dto.Gender,
	)
	dao.Dob, dao.Email = dto.Dob, dto.Email
	return dao
}

func daoTodto(dao ports.PersonDAO) ports.PersonDTO {
	var dto ports.PersonDTO
	dto.ID = dao.PersonId
	name, gender := decodeName(dao.Name)
	dto.Gender = gender
	dto.FirstName, dto.LastName = name[0], name[len(name)-1]
	dto.Age = getAge(dao.Dob)
	return dto
}
