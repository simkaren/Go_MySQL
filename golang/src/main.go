package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func openDb(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return openDb(path, count)
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))
	return openDb(path, 100)
}

type Data struct {
	id   int
	text string
}

func main() {
	fmt.Printf("Go launched.\n")
	db := connectDB()
	defer db.Close()
	printAllRows(db)
	d := Data{}
	d.text = "new inserted"
	insertElement(d, db)
	printAllRows(db)
}

func printAllRows(db *sql.DB) {
	fmt.Printf("Print all rows.\n")
	rows, err := db.Query("SELECT * FROM TEST;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		data := Data{}
		err = rows.Scan(&data.id, &data.text)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", data)
	}
}

func insertElement(d Data, db *sql.DB) {
	fmt.Printf("Insert %v into DB.\n", d)
	_, err := db.Exec("INSERT INTO TEST VALUES(?,?)", d.id, d.text)
	if err != nil {
		panic(err)
	}
}
