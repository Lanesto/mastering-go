package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var data []Entry
var tpl string

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("host: %s path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tpl))
	myT.ExecuteTemplate(w, tpl, data)
}

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Printf("Usage: %s <sqlite_db_file> <template>\n", filepath.Base(args[0]))
		return
	}

	fmt.Println("accessing database")
	database := args[1]
	tpl = args[2]
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("emptying database table.")
	_, err = db.Exec("DELETE FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("populating", database)
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) VALUES(?, ?, ?)")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
	}

	rows, err := db.Query("SELECT * FROM data")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		temp := Entry{}
		err = rows.Scan(&temp.Number, &temp.Double, &temp.Square)
		data = append(data, temp)
	}

	fmt.Println("serving")
	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
