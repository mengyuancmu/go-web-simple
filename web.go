package main

import "fmt"
import "github.com/gin-gonic/gin"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
	db, err := sql.Open("mysql", "user:password@localhost/mydb")
	fmt.Println("vim-go")
}
