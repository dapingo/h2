package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	//for template
	r.LoadHTMLGlob("templates/*")
	//for static file
	r.Static("/assets","./assets")
	//
	r.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{

		})
	})
	r.GET("/a", func(c *gin.Context){
		c.JSON(200,gin.H{
			"name":"daping",
		})
	})
	r.GET("/abc", func(c *gin.Context){
		c.JSON(200,gin.H{
			"name":"daping",
		})
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(200,"index.html",gin.H{

		})
	})
	r.Run(`127.0.0.1:8080`) // 监听并在 0.0.0.0:8080 上启动服务
}