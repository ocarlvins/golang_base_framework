package db

import (
	"database/sql"
	"fmt"
	"log"

	// Import the PostgreSQL driver
	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	// Define the PostgreSQL connection parameters
	dbURL := "postgres://postgres:postgres2019@localhost:5432/testdb"

	// Open a connection to the PostgreSQL database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()

	// Verify the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the PostgreSQL database")
	return db
}

func runQuery(query string) sql.Result {
	// Insert data into the database
	db := connect()
	defer db.Close()
	// Prepare the SQL statement for insertion
	//sqlStatement := ` INSERT INTO your_table (column1, column2) VALUES ($1, $2)`

	// Execute the SQL statement
	fmt.Println("Data inserted successfully")
	Result, err := db.Exec(query)
	if err != nil {
		fmt.Println("There was an error in creating the table")
		fmt.Println(err)
	}
	return Result
}
