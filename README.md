# GoSeven

 Go Lang Learning : https://go.dev/doc/tutorial 

## Gin 安装

    安装 Gin 提示  Get "https://proxy.golang.org/github.com/gin-gonic/gin/@v/list": dial tcp 142.250.72.177:443: i/o timeout
    需要切换代理 
        默认代理GOPROXY='https://proxy.golang.org,direct'
        换成 go env -w GOPROXY=https://goproxy.io,direct  ｜｜ GOPROXY='https://goproxy.io,direct' 

## 测试 Gin 安装成功

    go mod init gin-demo 
    go mod tidy
    go get -u github.com/gin-gonic/gin
    go work use .
    go run .

## Running Gin

    package main

    import (
      "net/http"

      "github.com/gin-gonic/gin"
    )

    func main() {
      r := gin.Default()
      r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
          "message": "pong",
        })
      })
      r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
    }

    // visit 0.0.0.0:8080/ping on browser
