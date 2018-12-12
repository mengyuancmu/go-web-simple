package main

import (
	"net/http"
)
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
	Port     string `toml:"port"`
	ConnMax  int    `toml:"connection_max"`
}
type redis struct {
	Server  string
	Port    int
	ConnMax int `toml:"connection_max"`
}
type post struct {
	Id      int
	Content string
}
type comment struct {
	Id      int
	Content string
}

func main() {
	var config tomlConfig
	_, err := toml.DecodeFile("config.toml", &config)
	dsn := config.DB.User + ":" + config.DB.Password + "@tcp(" + config.DB.Server + ":" + config.DB.Port + ")/" + config.DB.Dbname
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(config.DB.ConnMax)
	db.SetMaxIdleConns(config.DB.ConnMax)
	defer db.Close()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/post/:id", func(c *gin.Context) {
		postId := c.Param("id")
		postRs := make(chan post)
		go func() {
			var postObj post
			row := db.QueryRow("select * from tb_post where id = ?", postId)
			row.Scan(&postObj.Id, &postObj.Content)
			postRs <- postObj
		}()
		commentRs := make(chan []comment)
		go func() {
			var commentList []comment
			rows, err := db.Query("select id,content from tb_comment where post_id = ?", postId)
			if err != nil {
				panic(err)
			}
			defer rows.Close()
			for rows.Next() {
				var commentObj comment
				rows.Scan(&commentObj.Id, &commentObj.Content)
				commentList = append(commentList, commentObj)
			}
			commentRs <- commentList
		}()
		postInfo := <-postRs
		commentListInfo := <-commentRs

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"Post":     postInfo,
			"Comments": commentListInfo,
		})
	})
	router.Run(":8080")
}
