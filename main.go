package main

import (
	"200lab/food-delivery/common"
	"200lab/food-delivery/component"
	"200lab/food-delivery/middleware"
	"200lab/food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
}

func main() {
	var env Env
	env = Init()

	db, err := gorm.Open(mysql.Open(env.DBConnectionStr), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	appCtx := component.NewAppContext(db)

	r.Use(middleware.Recover(appCtx))

	r.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "static/img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)

		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		fileName := fileHeader.Filename
		destination := fmt.Sprintf("./%s/%s", folder, fileName)
		log.Println(destination)
		errUpload := c.SaveUploadedFile(fileHeader, destination)
		if errUpload != nil {
			c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": errUpload,
			})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(destination))
	})
	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurantByCondition(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurantById(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
