package main
 

import (
    "database/sql"
    "fmt"
    "go_todo_api/controllers"
    "go_todo_api/models"
    "log"
    "net/http"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    // "os"
    "html/template"
)

var tpl *template.Template
 
func main() {
    db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
 
	// テンプレートを読み込む
	tpl = template.Must(template.ParseFiles("root/index.html"))

    todoModel := models.NewTodoModel(db)
    todoHandler := controlllers.NewTodoController(todoModel)
 
    router := mux.NewRouter()
    // router.PathPrefix("/").Handler(http.FileServer(http.Dir("root/"))).Methods("GET")
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        err := tpl.Execute(w, nil)
        if err != nil {
            fmt.Println("Error rendering template:", err)
            http.Error(w, "Error rendering template", http.StatusInternalServerError)
        }
    })
    router.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
    router.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")

    
    fmt.Println("Server starting at :8080")
    log.Fatal(http.ListenAndServe(":8080", router))


}
