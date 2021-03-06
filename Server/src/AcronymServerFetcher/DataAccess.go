package main

import (
	"database/sql"
	"time"

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

// AcronymRepository is an interface for accessing Acronym table
type AcronymRepository interface {
	Open() error
	Insert(acronym *Acronym) error
	DeleteAll() error
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
	err = createTables(db)
	if err != nil {
		return err
	}

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

// DeleteAll deletes all entries in database
func (repository *SqliteAcronymRepository) DeleteAll() error {
	db := repository.db

	deleteCommand := "DELETE FROM Acronyms"
	_, err := db.Exec(deleteCommand)

	return err
}

// Insert inserts and acronym to the store
func (repository *SqliteAcronymRepository) Insert(acronym *Acronym) error {
	db := repository.db

	now := time.Now().Unix()

	// Insert command
	insertCommand :=
		`INSERT INTO
	            Acronyms (
	                Acronym, Definition, Language,
	                Url, CreatedAt, ModifiedAt
	            )
	            VALUES (
	                ?, ?, ?,
	                ?, ?, ?
	            )`

	statement, err := db.Prepare(insertCommand)
	if err != nil {
		return err
	}

	_, err = statement.Exec(
		acronym.Acronym,
		acronym.Definition,
		acronym.Language,
		acronym.Url,
		now,
		now,
	)
	statement.Close()

	if err != nil {
		return err
	}

	return err
}

func createTables(db *sql.DB) error {
	createAcronymTableCommand :=
		`CREATE TABLE IF NOT EXISTS Acronyms (
	        AcronymId INTEGER PRIMARY KEY NOT NULL,
	        Acronym TEXT NOT NULL,
	        Definition TEXT NOT NULL,
	        Language TEXT NOT NULL,
	        Url TEXT NULL,
	        CreatedAt INTEGER NOT NULL,
	        ModifiedAt INTEGER NOT NULL
	    )
	    `
	_, err := db.Exec(createAcronymTableCommand)
	return err
}
