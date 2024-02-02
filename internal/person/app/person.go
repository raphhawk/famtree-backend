package app

import (
	"log"
	"time"

	ae "github.com/raphhawk/famtree/cmd/common/error"
	"github.com/raphhawk/famtree/internal/person/db"
	"github.com/raphhawk/famtree/internal/person/ports"
)

type Person struct {
	Server ports.PersonServer
	DB     ports.PersonDatabase
}

func New() (*Person, error) {
	var personDB ports.PersonDatabase
	personDB = db.New(
		// take env vars
		db.POSTGRES_DRIVER,
		db.POSTGRES_PASSWORD,
		db.POSTGRES_DB,
		db.POSTGRES_USER,
		db.POSTGRES_PORT,
	)
	personDB.InitDB()

	// Initialize Person table
	res, err := personDB.GetDB().Exec(db.InitQuery)
	if err != nil {
		return nil, err
	}
	x, _ := res.LastInsertId()
	log.Println(x)
	return &Person{
		DB: personDB,
	}, nil
}

func (p Person) String() string {
	s := ``
	if p.DB.GetDB() != nil {
		s += "DB Active\n"
	} else {
		s += "DB Inactive\n"
	}
	return s
}

func handleResult(
	result map[string]int64,
	err error,
	successInfo ae.ArtificialErrors,
) ports.PersonDTO {
	if err != nil {
		info := ae.InvalidDBId
		info.Error = err
		log.Println("RH")
		return ports.PersonDTO{
			Info: info,
		}
	}

	info := successInfo
	info.Others = result
	return ports.PersonDTO{
		Info: info,
	}
}

func (p *Person) CreatePerson(person ports.PersonDTO) ports.PersonDTO {
	dao := dtoTodao(person)
	var lastid int64
	err := p.DB.GetDB().QueryRow(
		db.CreateQuery,
		dao.Name,
		dao.Dob,
		dao.Email,
		time.Now(),
		time.Now(),
	).Scan(&lastid)
	log.Println("lastid", lastid)
	return handleResult(
		map[string]int64{"lastid": lastid}, err, ae.InsertSuccess,
	)
}

func (p *Person) GetPersonById(id int) ports.PersonDTO {
	var dao ports.PersonDAO
	err := p.DB.GetDB().QueryRow(
		db.GetByIdQuery, id).
		Scan(
			&dao.PersonId,
			&dao.Name,
			&dao.Dob,
			&dao.Email,
		)
	if err != nil {
		info := ae.InvalidDBId
		info.Error = err
		return ports.PersonDTO{
			Info: info,
		}
	}
	dto := daoTodto(dao)
	info := ae.ReadSuccess
	dto.Info = info
	return dto
}

func (p *Person) GetPersonByEmail(email string) (ports.PersonDTO, error) {
	return ports.PersonDTO{}, nil
}

func (p *Person) UpdatePersonName(person ports.PersonDTO) ports.PersonDTO {
	dao := dtoTodao(person)
	var lastid int64
	err := p.DB.GetDB().QueryRow(
		db.UpdateName,
		dao.Name,
		time.Now(),
		dao.PersonId,
	).Scan(&lastid)
	return handleResult(
		map[string]int64{"lastid": lastid}, err, ae.AlterSuccess,
	)
}

func (p *Person) UpdatePersonEmail(person ports.PersonDTO) ports.PersonDTO {
	dao := dtoTodao(person)
	var lastid int64
	err := p.DB.GetDB().QueryRow(
		db.UpdateEmail,
		dao.Email,
		time.Now(),
		dao.PersonId,
	).Scan(&lastid)
	return handleResult(
		map[string]int64{"lastid": lastid}, err, ae.AlterSuccess,
	)
}

func (p *Person) UpdatePersonDob(person ports.PersonDTO) ports.PersonDTO {
	dao := dtoTodao(person)
	var lastid int64
	err := p.DB.GetDB().QueryRow(
		db.UpdateDob,
		dao.Dob,
		time.Now(),
		dao.PersonId,
	).Scan(&lastid)
	return handleResult(
		map[string]int64{"lastid": lastid}, err, ae.AlterSuccess,
	)
}

// UpdatePersonGender: FirstName and LastName required
func (p *Person) UpdatePersonGender(person ports.PersonDTO) ports.PersonDTO {
	return p.UpdatePersonName(person)
}

func (p *Person) DeletePerson(person ports.PersonDTO) ports.PersonDTO {
	dao := dtoTodao(person)
	result, err := p.DB.GetDB().Exec(
		db.DeleteById,
		dao.PersonId,
	)
	rowsaffected, _ := result.RowsAffected()
	return handleResult(
		map[string]int64{"rowsaffected": rowsaffected}, err, ae.AlterSuccess,
	)
}
