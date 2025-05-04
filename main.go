package main

import(
	"os"
	"github.com/gin-gonic/gin"
	routes "github.com/12sub/go-jwt/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == ""{
		port="8000"
	}

	routr := gin.New()
	routr.Use(gin.Logger())

	routes.AuthRoutes(routr)
	routes.UserRoutes(routr)

	routr.GET("/api-1", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api 1"})
	})
	
	routr.GET("/api-2", func(c *gin.Context){
		c.JSON(200, gin.H{"success":"Access granted for api 2"})
	})

	routr.Run(":" + port)

}