package routes

import (
	//user "CRUD-Operation/controllers/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

//StartGin function
func StartGin() {
	router := gin.Default()
	//TODO Enable CORS
	//TODO Enable Size
	//TODO Enable Authorization

	api := router.Group("/api")
	{
		api.GET("/", hello)
		// api.GET("/users", user.GetAllUser)
		// api.POST("/users", user.CreateUser)
		// api.GET("/users/:id", user.GetUser)
		// api.PUT("/users/:id", user.UpdateUser)
		// api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}
