package service

import (
	"net/http"

	"github.com/204iot/gpserver/server/cache"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Error string `json:"error"`
}

func UploadGpsData(c *gin.Context) {
	var req struct {
		ID        string `json:"id" binding:"required"`
		Longitude string `json:"longitude" binding:"required"`
		Latitude  string `json:"latitude" binding:"required"`
	}

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Error{Error: err.Error()})
		return
	}

	_ = cache.Set(&cache.IOTDevice{
		ID:        req.ID,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
	})

	c.JSON(http.StatusOK, gin.H{})
}

func GetAllDevices(c *gin.Context) {
	devices := cache.GetAll()
	if devices == nil {
		c.JSON(http.StatusOK, []int{})
		return
	}

	c.JSON(http.StatusOK, devices)
}
