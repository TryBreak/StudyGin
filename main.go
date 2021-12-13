package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	gin "github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// 模拟一些私人数据
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 设置filter 方法
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	// 设置模板路径
	router.LoadHTMLGlob("templates/**/*")
	// 静态目录
	router.Static("/assets", "./assets")
	router.StaticFS("/more_static", http.Dir("my_file_system"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	router.GET("/", func(c *gin.Context) {
		pusher := c.Writer.Pusher()
		fmt.Println("mo7-pusher", pusher)
		if pusher != nil {
			// 使用 pusher.Push() 做服务器推送
			err := pusher.Push("/assets/app.js", nil)
			fmt.Println("mo7-err", err)
			if err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}

		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"status": "success",
		})
	})

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
			"title": "post",
		})
	})
	router.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "user",
			"now":   time.Now(),
		})
	})

	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br />",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	return router
}

func main() {
	router := setupRouter()

	router.Run(":9000")
}
