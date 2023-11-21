package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func Connection() {
	db, err1 := sql.Open("sqlite3", "my_db.db")
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	defer db.Close()

	// Ticket
	_, err2 := db.Exec("CREATE TABLE IF NOT EXISTS ticket (id INTEGER PRIMARY KEY, user_id INTEGER, title TEXT, description TEXT, category_id INTEGER, status TEXT, created_at TEXT, updated_at TEXT)")
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	// Ticket Comments
	_, err3 := db.Exec("CREATE TABLE IF NOT EXISTS ticket_comment (id INTEGER PRIMARY KEY, user_id INTEGER, supporter_id INTEGER, message TEXT, created_at TEXT, updated_at TEXT)")
	if err3 != nil {
		fmt.Println(err2.Error())
	}

	// Ticket Category
	_, err4 := db.Exec("CREATE TABLE IF NOT EXISTS ticket_category (id INTEGER PRIMARY KEY, name TEXT, content TEXT)")
	if err4 != nil {
		fmt.Println(err2.Error())
	}

	// Question
	_, err5 := db.Exec("CREATE TABLE IF NOT EXISTS question (id INTEGER PRIMARY KEY, title TEXT, content TEXT)")
	if err5 != nil {
		log.Fatal(err5)
	}

	// Question Category
	_, err6 := db.Exec("CREATE TABLE IF NOT EXISTS question_category (id INTEGER PRIMARY KEY, name TEXT, content TEXT)")
	if err6 != nil {
		log.Fatal(err6)
	}

	fmt.Println("Successfully connected!")
	return
}
