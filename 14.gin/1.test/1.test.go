package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/xxx", "./statics")
	r.SetFuncMap(template.FuncMap{
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	})

	r.LoadHTMLGlob("templates/**/*")
	r.GET("/posts", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts", gin.H{
			"Name": "<script>alert('hello')</script>",
		})
	})
	r.GET("/users", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users", gin.H{
			"Name": "LJY",
		})
	})
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"Name": "<script>alert('hello')</script>",
		})
	})

	r.Run(":9000")
}
