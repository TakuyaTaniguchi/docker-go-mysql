package main
 
import (
      "database/sql"
      "fmt"
      "log"
 
      _ "github.com/go-sql-driver/mysql"
)
 
func main() {
    db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/todo_db")
      if err != nil {
            log.Fatal(err)
      }
      defer db.Close()
 
      stmt, err := db.Prepare("INSERT INTO todos (task) VALUES (?)")
      if err != nil {
            log.Fatal(err)
      }
      defer stmt.Close()
 
      tasks := []string{
            "Task 1",
            "Task 2",
            "Task 3",
      }
 
      for _, task := range tasks {
            _, err := stmt.Exec(task)
            if err != nil {
                  log.Fatal(err)
            }
      }
 
      fmt.Println("Seeder executed successfully")
}