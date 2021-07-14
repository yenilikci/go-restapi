package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/demodb")
	if err != nil {
		panic(err.Error())
	}
	db.Ping()
	defer db.Close()

	//insert data
	res, err := db.Exec("INSERT INTO users(Username,Email,Password,FirstName,LastName,BirthDate,IsActive) VALUES('Deneme1','deneme@deneme.com','Name1','password1','Firstname1','1900.1.1',1)")
	if err != nil {
		log.Fatal(err)
	}
	rowCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted %d rows", rowCount)

	lastID, err1 := res.LastInsertId()
	if err1 != nil {
		log.Fatal(err)
	}
	log.Printf("Last insert id: %d", lastID)
}
