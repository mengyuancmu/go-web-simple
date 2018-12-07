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
	Server   string
	User     string
	Password string
	Dbname   string `toml:"db_name"`
	Port     int
	ConnMax  int `toml:"connection_max"`
}
type redis struct {
	Server  string
	Port    int
	ConnMax int `toml:"connection_max"`
}

func main() {
	_, err := toml.DecodeFile("config.toml", &tomlConfig)
	db, err := sql.Open("mysql", tomlConfig.DB.User+":"+tomlConfig.DB.Password+"@"+tomlConfig.DB.Server+"/"+tomlConfig.DB.Dbname)
	db.SetMaxOpenConns(tomlConfig.DB.ConnMax)
	db.SetMaxIdleConns(tomlConfig.DB.ConnMax)
	defer db.Close()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/post/:id", func(c *gin.Context) {
		postId := c.Param("id")
		row, err := 
	})
	router.Run(":8080")
}
