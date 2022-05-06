package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func Read(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var uid int
		var name string
		var lastname string
		var created time.Time
		err = rows.Scan(&uid, &name, &lastname, &created)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(uid)
		fmt.Println(name)
		fmt.Println(lastname)
		fmt.Println(created)
	}
}

func Update(db *sql.DB) {
	stmt, err := db.Prepare("UPDATE person set lastname=? where uid=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("smith", 1)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	log.Printf("Affected rows %d", affected)
}

func Create(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO person(name, lastname, created) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Mrs", "Smith", "2022-01-01")
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	log.Printf("Affected rows %d", affected)
}

func Delete(db *sql.DB) {
	stmt, err := db.Prepare("delete from person where uid=?")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(1)
	if err != nil {
		log.Fatal(err)
	}
	affected, _ := res.RowsAffected()
	log.Printf("Affected rows %d", affected)
}

func main() {
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("database open")
	Create(db)
	Read(db)
	// Update(db)

	fmt.Println("bye")

	fmt.Println("closing db")
	db.Close()

}
