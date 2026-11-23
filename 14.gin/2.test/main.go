package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("login.html", "index.tmpl", "push.html")

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

	r.GET("/push", func(c *gin.Context) {
		c.HTML(http.StatusOK, "push.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "文件上传失败",
			})
			return
		}
		c.SaveUploadedFile(file, "./upload/"+file.Filename)
		c.JSON(http.StatusOK, gin.H{
			"message":  "文件上传成功",
			"filename": file.Filename,
		})
	})

	r.POST("/uploadFiles", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "文件上传失败",
			})
			return
		}
		files := form.File["files"]
		for index, file := range files {
			c.SaveUploadedFile(file, "./upload/"+strconv.Itoa(index)+"_"+file.Filename)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "文件上传成功",
		})
	})

	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/login")
	})

	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "奶龙 404 Not Found",
		})
	})

	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/list", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "视频列表",
			})
		})
		videoGroup.GET("/detail", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "视频详情",
			})
		})
		videoGroup.GET("/upload", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "视频上传",
			})
		})
		uservideoGroup := videoGroup.Group("/user")
		{
			uservideoGroup.GET("/list", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "用户列表",
				})
			})
		}
	}
	r.Any("/any", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK, gin.H{
				"message": "any get",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{
				"message": "any post",
			})
		}
	})

	handleDesign := func(c *gin.Context) {
		fmt.Println("handleDesign")
		c.JSON(http.StatusOK, gin.H{
			"message": "design",
		})
	}
	m := func(c *gin.Context) {
		fmt.Println("m")
		start := time.Now()
		c.Next()	// 调用下一个中间件
		fmt.Println(time.Since(start))
		fmt.Println("m end")
	}
	//r.Use(m)// 全局中间件
	r.GET("/design", m, handleDesign)

	r.Run(":9000")
}
