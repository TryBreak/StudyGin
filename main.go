package main

import (
	"fmt"
	"html/template"
	http "net/http"
	"time"

	gin "github.com/gin-gonic/gin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// 设置filter 方法
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	// 设置模板路径
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", func(c *gin.Context) {
		fmt.Println("mo77,index")
		c.HTML(http.StatusOK, "index/index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	r.GET("/post", func(c *gin.Context) {
		c.HTML(http.StatusOK, "post/index.tmpl", gin.H{
			"title": "post",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.tmpl", gin.H{
			"title": "user",
			"now":   time.Now(),
		})
	})

	r.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br />",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":9000")
}
