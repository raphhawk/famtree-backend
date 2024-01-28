package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	POSTGRES_DRIVER   = "pgx"
	POSTGRES_PASSWORD = "famtree_password"
	POSTGRES_USER     = "famtree_user"
	POSTGRES_DB       = "famtree_db"
	POSTGRES_PORT     = "5432"
)

type Db struct {
	dbtype     string
	dbpassword string
	dbname     string
	dbuser     string
	dbport     string
	dbinstance *sql.DB
}

func New(
	dbtype, password, name, user, port string,
) *Db {
	return &Db{
		dbtype:     dbtype,
		dbpassword: password,
		dbname:     name,
		dbuser:     user,
		dbport:     port,
	}
}

func (db *Db) OpenDB(dsn string) (*sql.DB, error) {
	dbConnection, err := sql.Open(db.dbtype, dsn)
	if err != nil {
		return nil, err
	}

	if err := dbConnection.Ping(); err != nil {
		return nil, err
	}
	return dbConnection, nil
}

func (db *Db) ConnectToDB() *sql.DB {
	counts := 0
	dsn := fmt.Sprintf(
		"postgres://%v:%v@localhost:%v/%v",
		db.dbuser,
		db.dbpassword,
		db.dbport,
		db.dbname,
	)
	for {
		counts++
		connection, err := db.OpenDB(dsn)
		if err != nil {
			log.Println(db.dbtype, "not yet ready...")
		} else {
			log.Println(db.dbtype, "connected to DB")
			return connection
		}
		if counts > 10 {
			return nil
		}
		log.Println("Backing off for 1s")
		time.Sleep(1 * time.Second)
		continue
	}
}

func (db *Db) InitDB() error {
	conn := db.ConnectToDB()
	if conn == nil {
		return errors.New("[ERROR] cant connect to database")
	}
	db.dbinstance = conn
	return nil
}

func (db *Db) GetDB() *sql.DB {
	return db.dbinstance
}
