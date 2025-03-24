package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Go lang ,and Gin")
	})

  // router.GET("/ping", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{
  //     "message": "pong",
  //   })
  // })

  // router.GET("/user/:name", func(c *gin.Context) {
  //   name := c.Param("name")
  //   c.String(http.StatusOK, "Hello %s", name)
  // })


// 匹配users?name=xxx&role=xxx，role可选
// router.GET("/users", func(c *gin.Context) {
// 	name := c.Query("name")
// 	role := c.DefaultQuery("role", "teacher")
// 	c.String(http.StatusOK, "%s is a %s", name, role)
// })


// POST
// router.POST("/form", func(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.DefaultPostForm("password", "000000") // 可设置默认值

// 	c.JSON(http.StatusOK, gin.H{
// 		"username": username,
// 		"password": password,
// 	})
// })


// GET 和 POST 混合
// router.POST("/posts", func(c *gin.Context) {
// 	id := c.Query("id")
// 	page := c.DefaultQuery("page", "0")
// 	username := c.PostForm("username")
// 	password := c.DefaultPostForm("username", "000000") // 可设置默认值

// 	c.JSON(http.StatusOK, gin.H{
// 		"id":       id,
// 		"page":     page,
// 		"username": username,
// 		"password": password,
// 	})
// })


// map 参数
// router.POST("/post", func(c *gin.Context) {
// 	ids := c.QueryMap("ids")
// 	names := c.PostFormMap("names")

// 	c.JSON(http.StatusOK, gin.H{
// 		"ids":   ids,
// 		"names": names,
// 	})
// })


// router.GET("/redirect", func(c *gin.Context) {
//     c.Redirect(http.StatusMovedPermanently, "/index")
// })

// router.GET("/goindex", func(c *gin.Context) {
// 	c.Request.URL.Path = "/"
// 	router.HandleContext(c)
// })



// group routes 分组路由
defaultHandler := func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"path": c.FullPath(),
	})
}
// group: c1
c1 := router.Group("/c1")
{
	c1.GET("/gets", defaultHandler)
	c1.GET("/list", defaultHandler)
}
// group: c2
c2 := router.Group("/c2")
{
	c2.GET("/gets", defaultHandler)
	c2.GET("/list", defaultHandler)
}

  router.Run() // listen and serve on 0.0.0.0:8080
}


