package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("login.html", "index.tmpl")

	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
			"name":    "LJY",
			"age":     18,
			"sex":     "female",
		})
	})

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Sex  string `json:"sex"`
	}
	r.GET("/otherjson", func(c *gin.Context) {
		c.JSON(http.StatusOK, Person{
			Name: "LJY",
			Age:  18,
			Sex:  "female",
		})
	})

	r.GET("/web", func(c *gin.Context) {
		name := c.Query("name")
		age := c.Query("age")
		// name, ok := c.GetQuery("name")
		// if !ok {
		// 	name = "奶龙"
		// }
		//name := c.DefaultQuery("name", "奶龙")

		c.JSON(http.StatusOK, gin.H{
			"message": "hello " + name,
			"age":     age,
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		//username := c.DefaultPostForm("username", "默认用户名")
		// username, ok := c.GetPostForm("username")
		// if !ok {
		// 	username = "默认用户名"
		// }
		c.HTML(http.StatusOK, "index", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.GET("/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello " + name,
			"age":     age,
		})
	})

	type User struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}

	r.POST("/user", func(c *gin.Context) {
		var user User
		//err := c.ShouldBindQuery(&user)
		//err := c.ShouldBindJSON(&user)
		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "参数错误",
			})
			return
		} else {
			fmt.Println("参数正确")
			fmt.Println(user)
		}

	})

	r.Run(":9000")
}
