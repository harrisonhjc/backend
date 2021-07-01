package router

import (
	"backend/api/v1"
	"net/http"
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	

	
)


func Init() {
	r := gin.Default()
	pprof.Register(r)

	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", 8000))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	lineRouter := r.Group("v1")
	{
		lineRouter.GET("/line", v1.Line)

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	r.Run(":8000")

}
