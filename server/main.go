package main

import (
	"flag"

	"github.com/204iot/gpserver/server/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	listenAddr := flag.String("l", "0.0.0.0:8080", "监听地址")
	flag.Parse()

	r := gin.New()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	r.Use(cors.New(corsConfig))
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api/v1")
	api.POST("/device", service.UploadGpsData)
	api.GET("/devices", service.GetAllDevices)

	if err := r.Run(*listenAddr); err != nil {
		panic(err)
	}
}
