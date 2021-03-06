package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

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

	var (
		ID        int
		Username  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		BirthDate string
		IsActive  bool
	)

	//get data
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	// rows.Columns()
	for rows.Next() {
		err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	}
	//alternative
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	rows.Close()

	//get specific data
	rows, errQ := db.Query("SELECT * from users WHERE ID = ?", 2)
	if errQ != nil {
		log.Fatal(errQ)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	}

	errQ = rows.Err()
	if errQ != nil {
		log.Fatal(errQ)
	}

	//query row
	err = db.QueryRow("SELECT * FROM users limit 1").Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			//kayıt yoksa
		} else {
			log.Fatal(err)
		}
	}
	log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))

	// multiple active result set
	// _, err := db.Exec("DELETE FROM xTable1; DELETE FROM xTable2")

	//preparing queries
	stmt, errQ := db.Prepare("SELECT * FROM users WHERE ID=?")
	if errQ != nil {
		log.Fatal(errQ)
	}
	defer stmt.Close()
	rows, errX := stmt.Query(5)
	if errX != nil {
		log.Fatal(errX)
	}
	defer rows.Close()
	for rows.Next() {
		scanErr := rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		log.Printf("Bulunan satır içeriği: %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	}

	//preparing query - single rows
	stmt, errQ2 := db.Prepare("SELECT * FROM users WHERE ID = ?")
	if errQ2 != nil {
		log.Fatal(errQ2)
	}
	errX2 := stmt.QueryRow(4).Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
	if errX2 != nil {
		log.Fatal(errX2)
	}
	fmt.Println(FirstName + " " + LastName)

	//modifying data - delete
	res2, _ := db.Exec("DELETE FROM users LIMIT 1")
	rowCount2, _ := res2.RowsAffected()
	fmt.Println(rowCount2)
	lastID2, _ := res2.LastInsertId()
	fmt.Println(lastID2)

	//insert - prepare,exec
	stmtV, _ := db.Prepare("INSERT INTO users(Username,Email,Password,FirstName,LastName,BirthDate,IsActive) VALUES(?,?,?,?,?,?,?)")
	Username = "hjkl"
	Email = "hjkl@example.com"
	Password = "12341"
	FirstName = "abc"
	LastName = "xyz"
	BirthDate = "1999.1.1"
	IsActive = true
	resV, errStmt := stmtV.Exec(Username, Email, Password, FirstName, LastName, BirthDate, IsActive)
	if errStmt != nil {
		log.Fatal(errStmt)
	}
	fmt.Println(resV.LastInsertId())

	//transaction
	tx, errTx := db.Begin()
	//...
	_, errTx = db.Exec("Update ...", args)
	//...
	if errTx != nil {
		log.Fatal(errTx)
		tx.Rollback()
	}
}
