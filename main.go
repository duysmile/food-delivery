package main

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/component/uploadprovider"
	"200lab/food-delivery/middleware"
	"200lab/food-delivery/modules/food/foodtransport/ginfood"
	"200lab/food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"200lab/food-delivery/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"200lab/food-delivery/modules/upload/uploadtransport/ginupload"
	"200lab/food-delivery/modules/user/usertransport/ginuser"
	"200lab/food-delivery/pubsub/pblocal"
	"200lab/food-delivery/subscriber"
	"log"

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

	provider := uploadprovider.NewS3Provider(
		env.S3BucketName,
		env.S3Region,
		env.S3APIKey,
		env.S3Secret,
		env.S3Domain,
	)

	if err := runService(db, provider, env.SecretKeyJWT); err != nil {
		log.Fatalln(err)
	}
}

func runService(
	db *gorm.DB,
	provider uploadprovider.UploadProvider,
	secretKey string,
) error {
	appCtx := component.NewAppContext(db, provider, secretKey, pblocal.NewPubSub())

	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatalln(err)
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.UploadImage(appCtx))

	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfileUser(appCtx))

	restaurants := v1.Group("/restaurants")
	{
		restaurants.GET("/:id/foods", ginfood.ListFoodByRestaurantId(appCtx))

		restaurants.POST("", middleware.RequireAuth(appCtx), ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurantByCondition(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurantById(appCtx))
		restaurants.PATCH("/:id", middleware.RequireAuth(appCtx), ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.POST("/:id/like", middleware.RequireAuth(appCtx), ginrestaurantlike.LikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike", middleware.RequireAuth(appCtx), ginrestaurantlike.UnlikeRestaurant(appCtx))
		restaurants.DELETE("/:id", middleware.RequireAuth(appCtx), ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginrestaurantlike.ListUsersLikeRestaurant(appCtx))
	}

	foods := v1.Group("/foods")
	{
		foods.POST("", ginfood.CreateFood(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
