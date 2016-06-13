package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Acronym is struct which represents  structure of Acronym table
type Acronym struct {
	AcronymId  int
	Acronym    string
	Definition string
	Language   string
	Url        string
	CreatedAt  int64
	ModifiedAt int64
}

// NewAcronym returns new Acronym
func NewAcronym() *Acronym {
	acronym := new(Acronym)
	return acronym
}

var acronymCreatedAtIndex = 6
var emptyAcronymArray = make([]Acronym, 0)

// AcronymRepository is an interface for accessing Acronym table
type AcronymRepository interface {
	Open() error
	Find(acronym string) (found []Acronym, err error)
	Random(count int) (found []Acronym, err error)
	Close() error
}

// SqliteAcronymRepository is a concreate implementation of AcronymRepository using SQLite as store
type SqliteAcronymRepository struct {
	Driver   string
	Database string

	db *sql.DB
}

// NewSqliteAcronymRepository returns new SqliteAcronymRepository
func NewSqliteAcronymRepository(databasePath string) *SqliteAcronymRepository {
	repository := new(SqliteAcronymRepository)
	repository.Database = databasePath
	repository.Driver = "sqlite3"

	return repository
}

// Open opens database for transactions
func (repository *SqliteAcronymRepository) Open() error {
	db, err := sql.Open(repository.Driver, repository.Database)
	if err != nil {
		return err
	}

	repository.db = db

	return nil
}

// Close closes database
func (repository *SqliteAcronymRepository) Close() error {
	if repository.db != nil {
		repository.db.Close()
		repository.db = nil
	}

	return nil
}

func (repository *SqliteAcronymRepository) Find(acronym string) (found []Acronym, err error) {

	db := repository.db

	selectQuery :=
		`SELECT * FROM Acronyms WHERE Acronym LIKE ?`

	statement, err := db.Prepare(selectQuery)
	if err != nil {
		return emptyAcronymArray, err
	}
	defer statement.Close()

	rows, err := statement.Query(
		acronym + "%",
	)
	if err != nil {
		return emptyAcronymArray, err
	}
	defer rows.Close()

	found = make([]Acronym, 0)
	for rows.Next() {
		dbFound := *NewAcronym()
		rows.Scan(
			&dbFound.AcronymId,
			&dbFound.Acronym,
			&dbFound.Definition,
			&dbFound.Language,
			&dbFound.Url,
			&dbFound.CreatedAt,
			&dbFound.ModifiedAt,
		)
		found = append(found, dbFound)
	}

	return found, err
}

func (repository *SqliteAcronymRepository) Random(count int) (found []Acronym, err error) {

	db := repository.db

	selectQuery :=
		`SELECT * FROM Acronyms ORDER BY RANDOM() LIMIT ?`

	statement, err := db.Prepare(selectQuery)
	if err != nil {
		return emptyAcronymArray, err
	}
	defer statement.Close()

	rows, err := statement.Query(
		count,
	)
	if err != nil {
		return emptyAcronymArray, err
	}
	defer rows.Close()

	found = make([]Acronym, 0)
	for rows.Next() {
		dbFound := *NewAcronym()
		rows.Scan(
			&dbFound.AcronymId,
			&dbFound.Acronym,
			&dbFound.Definition,
			&dbFound.Language,
			&dbFound.Url,
			&dbFound.CreatedAt,
			&dbFound.ModifiedAt,
		)
		found = append(found, dbFound)
	}

	return found, err
}
