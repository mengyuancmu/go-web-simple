package main

import "fmt"
import "github.com/gin-gonic/gin"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/BurntSushi/toml"

type tomlConfig struct {
	DB    database `toml:"database"`
	Redis redis
}
type database struct {
	Server  string
	Port    int
	ConnMax int `toml:"connection_max"`
}
type redis struct {
	Server  string
	Port    int
	ConnMax int
}

func main() {
	configBlob, _ := ioutil.ReadFile(".config")
	config
	db, err := sql.Open("mysql", "user:password@localhost/mydb")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/post/:id", func(c *gin.Context) {
		postId := c.Param("id")
	})
	router.Run(":8080")
}
