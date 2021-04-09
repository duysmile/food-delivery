package main

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"log"
	"net/http"
	"strconv"

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

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurantByCondition(appCtx))
		restaurants.GET("/:id", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			if err != nil {

			}
			var data Restaurant

			if err := db.Where("id = ?", id).First(&data).Error; err != nil {
				c.JSON(401, map[string]interface{}{
					"error": err.Error(),
				})
				return
			}

			c.JSON(http.StatusOK, data)
		})
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
