package sql

import (
	"database/sql"
	"log"
)


type exampleModelWithStmt struct {
	DB         *sql.DB
	insertStmt *sql.Stmt
}

func NewExampleModel(db *sql.DB) (*exampleModelWithStmt, error) {
	
	inStmt, err := db.Prepare("INSERT INTO...")
	if err != nil {
		return nil, err
	}

	return &exampleModelWithStmt{db, inStmt}, nil
}


func (m *exampleModelWithStmt) Insert(args ...) error {
	
	if _, err := m.insertStmt.Exec(args...); err != nil {
		log.Printf("prepared Insert statement failed: %v", err)
		return err
	}

	return nil
}

// In the web application's main function we will need to initialize a new
// ExampleModel struct using the constructor function.

func examplMain() {
	db, err := sql.Open(...)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a new exampleModelWithStmt object, which includes the prepared statement.
	exampleModelWithStmt, err := NewExampleModel(db)
	if err != nil {
		log.Fatal(err)
	}

	// Defer a call to Close on the prepared statment to ensure that it is properly
	// closed before our main function terminates.
	defer func() {
		err := exampleModelWithStmt.insertStmt.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
