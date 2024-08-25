package routes

import (
	"github.com/gin-gonic/gin"
	"imagu/db/repo"
	"imagu/middlewares"
	"imagu/util"
)

type statsResponse struct {
	Images      int              `json:"images"`
	SubImages   int              `json:"sub_images"`
	TotalImages int              `json:"total_images"`
	ImageSize   string           `json:"image_size"`
	AccessStats util.AccessStats `json:"access_stats"`
}

func InitStatsRoutes(r *gin.Engine) {
	r.GET("/api/v1/stats", middlewares.AuthPermission("viewStats", false), getStats)
}

func getStats(c *gin.Context) {
	images, _ := repo.GetAllImages()
	imageCount := len(images)
	subImageCount := 0
	var imagesSize int64 = 0
	for _, value := range images {
		imagesSize += value.Size
		subImageCount += value.SubImages
	}
	stats := statsResponse{
		Images:      imageCount,
		SubImages:   subImageCount,
		TotalImages: imageCount + subImageCount,
		ImageSize:   util.FormatBytesToString(imagesSize),
		AccessStats: *util.CurrentAccessStats,
	}
	c.JSON(200, stats)
}
