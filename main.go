package main

import (
    "database/sql"
    //"fmt"
    "html/template"
    "net/http"
    "path"
    _ "github.com/mattn/go-sqlite3"
)

//Compile templates on start
var templates = template.Must(template.ParseFiles("templates/header.html", "templates/footer.html", "templates/main.html", "templates/about.html", "templates/boos2.html"))

type Boo struct {
    Title string
    Monster  string
    Says string
    Apology string
}

type Book struct {
    Title  string
    Author string
}


func ShowBook(w http.ResponseWriter, r *http.Request) {
    
    db := NewDB()
    var title, author string
    err := db.QueryRow("select title, author from books").Scan(&title, &author)
    if err != nil {
        panic(err)
    }

   // fmt.Fprintf(rw, "The first book is '%s' by '%s'", title, author)

   book := Book{title,author}

   fp := path.Join("templates", "index.html")
   tmpl, err := template.ParseFiles(fp)
       
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func NewDB() *sql.DB {
    db, err := sql.Open("sqlite3", "example.sqlite")
    if err != nil {
        panic(err)
    }

    _, err = db.Exec("create table if not exists books(title text, author text)")
    if err != nil {
        panic(err)
    }

    return db
}

func main() {
    //db := NewDB()
    http.HandleFunc("/boo", ScareBoos)
    http.HandleFunc("/", ShowStaticBook)
    http.HandleFunc("/books", ShowBook)
    http.ListenAndServe(":8080", nil)
    //http.ListenAndServe(":8080", ShowBooks(db))
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl, data)
}

func ScareBoos(w http.ResponseWriter, r *http.Request) {
    boo := Boo{"Boo Scare", "Boooogey Man", "bahhhhhhhh!", "oh, I'm sorry I scared you."}
    
    display(w, "boos2", boo)
    
    /*fp := path.Join("templates", "boos.html")
    
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, boo); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }*/
}

func ShowStaticBook(w http.ResponseWriter, r *http.Request) {
    //book := Book{"Building Web Apps with Go", "Jeremy Saenz"}
    book := Book{"Smooglie-Boo", "John Mosqueda"}

    fp := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, book); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

