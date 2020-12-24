package main

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

const dataFile = "dataFile.gob"

var data = make(map[string]record)

type record struct {
	Name    string
	Surname string
	ID      string
}

func save() error {
	fmt.Printf("Saving data to %s\n", dataFile)
	err := os.Remove(dataFile)
	if err != nil {
		log.Println("Could not remove existing data file:", err)
	}

	f, err := os.Create(dataFile)
	if err != nil {
		log.Println("Could not create data file:", err)
		return err
	}
	defer f.Close()

	encoder := gob.NewEncoder(f)
	if err := encoder.Encode(data); err != nil {
		log.Println("Could not save data to", dataFile)
		return err
	}

	return nil
}

func load() error {
	fmt.Printf("Loading data from %s\n", dataFile)
	f, err := os.Open(dataFile)
	if err != nil {
		log.Println("Could not open data file:", err)
		return err
	}
	defer f.Close()

	decoder := gob.NewDecoder(f)
	if err := decoder.Decode(&data); err != nil {
		log.Println("Could not load data from", dataFile)
		return err
	}

	return nil
}

func insert(key string, r record) bool {
	if key == "" {
		log.Println("Empty key given; ignored")
		return false
	}

	if get(key) == nil {
		log.Printf("Added new item with key %s, value %s\n", key, r)
		data[key] = r
		return true
	}

	log.Printf("Key %s already occupied", key)
	return false
}

func remove(key string) bool {
	if get(key) != nil {
		delete(data, key)
		log.Printf("Removed item with key %s\n", key)
		return true
	}

	log.Printf("Tried to remove item with key %s but has no data to delete\n", key)
	return false
}

func get(key string) *record {
	log.Printf("Someone looked up key %s\n", key)
	if _, ok := data[key]; ok {
		r := data[key]
		return &r
	}

	return nil
}

func update(key string, r record) bool {
	data[key] = r
	log.Printf("Item with key %s updated to %s\n", key, r)
	return true
}

func show() {
	for k, v := range data {
		fmt.Printf("%s: %v\n", k, v)
	}
}

func init() {
	fmt.Println("Initializing")
	f, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could not open log file; %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving:", r.Host, "for", r.URL.Path)
	myT := template.Must(template.ParseGlob("home.gotpl"))
	myT.ExecuteTemplate(w, "home.gotpl", nil)
}

func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store!")
	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">Home sweet home!</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\">List all elements</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right: 20px;\">Change an element</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right: 20px;\">Insert new element</a>")
	fmt.Fprintf(w, "<h1>The contents of the KV store are:</h1>")
	fmt.Fprintf(w, "<ul>")
	for k, v := range data {
		fmt.Fprintf(w, "<li>")
		fmt.Fprintf(w, "<strong>%s</strong> with value: %v\n", k, v)
		fmt.Fprintf(w, "</li>")
	}
	fmt.Fprintf(w, "</ul>")
}

func changeElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Changing an element of KV store")
	tmpl := template.Must(template.ParseFiles("update.gotpl"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	n := record{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}

	if !update(key, n) {
		fmt.Println("Update operation failed!")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func insertElement(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inserting an element to the kv store")
	tmpl := template.Must(template.ParseFiles("insert.gotpl"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	key := r.FormValue("key")
	n := record{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      r.FormValue("id"),
	}

	if !insert(key, n) {
		fmt.Println("Insert operation failed!")
	} else {
		err := save()
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

func main() {
	if err := load(); err != nil {
		log.Println(err)
	}
	defer func() {
		if err := save(); err != nil {
			log.Println(err)
		}
	}()

	port := ":8080"
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = args[1]
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeElement)
	http.HandleFunc("/list", listAll)
	http.HandleFunc("/insert", insertElement)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
	}
}
