package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "localuser:pass123@tcp(localhost:3306)/local_test")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func selectExample(db *sql.DB) error {
	rows, err := db.Query("SELECT name, age FROM characters")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var age int
		if err := rows.Scan(&name, &age); err != nil {
			return err
		}
		fmt.Printf("Name: %s, Age: %d\n", name, age)
	}
	return nil
}

func insertExample(db *sql.DB) error {
	result, err := db.Exec("INSERT INTO characters (name, age) VALUES (?, ?)", "Joseph Joestar", 30)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("%d record inserted.\n", rowsAffected)
	return nil
}

func updateExample(db *sql.DB) error {
	result, err := db.Exec("UPDATE characters SET age = ? WHERE name = ?", 40, "Joseph Joestar")
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	fmt.Printf("%d record(s) affected\n", rowsAffected)
	return nil
}

func transactionExample(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("INSERT INTO characters (name, age) VALUES (?, ?)", "Somkey", 20)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println(db)

	if err := selectExample(db); err != nil {
		panic(err)
	}

	if err := insertExample(db); err != nil {
		panic(err)
	}

	if err := updateExample(db); err != nil {
		panic(err)
	}

	if err := transactionExample(db); err != nil {
		panic(err)
	}
}
