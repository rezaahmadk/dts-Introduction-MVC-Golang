package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rezaahmadk/dts-Introduction-MVC-Golang/app/controller"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.POST("/api/v1/antrian", controller.AddAntrianHandler)
	router.GET("/api/v1/antrian/status", controller.GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", controller.UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrian/delete", controller.DeleteAntrianHandler)
	router.Run(":8080")
}
