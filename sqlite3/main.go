package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func createTable() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS "BOOKS" ("ID" INTEGER PRIMARY KEY, "TITLE" VARCHAR(255))`,
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Book struct {
	ID    int64
	Title string
}

func insertData(db *sql.DB, data Book) error {
	_, err := db.Exec(
		`INSERT INTO BOOKS (ID, TITLE) VALUES (?, ?)`,
		data.ID,
		data.Title,
	)
	if err != nil {
		return err
	}

	return nil
}

func selectAll(db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query(
		`SELECT * FROM BOOKS`,
	)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func selectByID(db *sql.DB, id int64) *sql.Row {
	row := db.QueryRow(
		`SELECT * FROM BOOKS WHERE ID=?`,
		id,
	)

	return row
}

func update(db *sql.DB, data Book) error {
	_, err := db.Exec(
		`UPDATE BOOKS SET TITLE=? WHERE ID=?`,
		data.Title,
		data.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func delete(db *sql.DB, id int64) error {
	_, err := db.Exec(
		`DELETE FROM BOOKS WHERE ID=?`,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// create books table
	db, err := createTable()
	if err != nil {
		log.Fatal("createTable()", err)
		return
	}

	// insert book data
	id := time.Now().Unix()
	var book = Book{
		id,
		"test",
	}
	err = insertData(db, book)
	if err != nil {
		log.Fatal("insertData()", err)
		return
	}

	// select all items
	rows, err := selectAll(db)
	if err != nil {
		log.Fatal("selectAll()", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}

		fmt.Printf("id: %d, title: %s\n", b.ID, b.Title)
	}

	// select item by id
	row := selectByID(db, id)
	var b Book
	err = row.Scan(&b.ID, &b.Title)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("Not found")
	case err != nil:
		log.Fatal("selectById()", err)
	default:
		fmt.Printf("id: %d, title: %s\n", book.ID, book.Title)
	}

	// update item by id
	dataForUpdate := Book{
		id,
		"updated",
	}
	if err := update(db, dataForUpdate); err != nil {
		log.Fatal("update()", err)
		return
	}

	row = selectByID(db, id)
	var checkData Book
	err = row.Scan(&checkData.ID, &checkData.Title)
	if err != nil {
		log.Fatal("selectById()", err)
		return
	}
	fmt.Printf("id: %d, title: %s\n", checkData.ID, checkData.Title)

	// delete item by id
	if err := delete(db, id); err != nil {
		log.Fatal("delete()", err)
		return
	}

	rows, err = selectAll(db)
	if err != nil {
		log.Fatal("selectAll()", err)
		return
	}

	defer rows.Close()
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title); err != nil {
			log.Fatal("rows.Scan()", err)
			return
		}

		fmt.Printf("id: %d, title: %s\n", b.ID, b.Title)
	}

}
