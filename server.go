package main

import (
	// "net/http"

	controller "gilab.com/pramaticreviews/golang-gin-poc/controllers"
	"gilab.com/pramaticreviews/golang-gin-poc/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	// server.GET("/test", func(ctx *gin.Context)  {
	// 	ctx.JSON(200, gin.H{
	// 		"message" : "OK!",
	// 	})
	// })

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})
	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}

// func main() {
// 	router := gin.Default()

// 	// Query string parameters are parsed using the existing underlying request object.
// 	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
// 	router.GET("/welcome", func(c *gin.Context) {
// 		firstname := c.DefaultQuery("firstname", "Guest")
// 		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

// 	// 	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// 	// })
// 	// // // This handler will match /user/john but will not match /user/ or /user
// 	// // router.GET("/user/:name", func(c *gin.Context) {
// 	// // 	name := c.Param("name")
// 	// // 	c.String(http.StatusOK, "Hello %s", name)
// 	// // })
// 	// // However, this one will match /user/john/ and also /user/john/send
// 	// // If no other routers match /user/john, it will redirect to /user/john/
// 	// router.GET("/user/:name/*action", func(c *gin.Context) {
// 	// 	name := c.Param("name")
// 	// 	action := c.Param("action")
// 	// 	message := name + " is " + action
// 	// 	c.String(http.StatusOK, message)
// 	// })

// 	// // For each matched request Context will hold the route definition
// 	// router.POST("/user/:name/*action", func(c *gin.Context) {
// 	// 	b := c.FullPath() == "/user/:name/*action" // true
// 	// 	c.String(http.StatusOK, "%t", b)
// 	// })

// 	router.Run(":8080")
// }
