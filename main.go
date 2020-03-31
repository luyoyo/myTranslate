package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    //注册一个路由器
    r := gin.Default()

    //默认请求  http://localhost:8080/
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, fmt.Sprintln(gin.H {"data": "hello world"}))
    })

    //get请求 返回json格式
    r.GET("/get", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"action": "get", "format": "json"})
    })

    //post请求 返回json格式
    r.POST("/post", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"action": "post", "format": "json"})
    })

    //get请求html页面
    r.GET("/index", func(c *gin.Context) {
        r.LoadHTMLFiles("view/index.html")
        c.HTML(http.StatusOK, "index.html", nil)
    })

    //运行（默认是8080端口）
    r.Run()
}