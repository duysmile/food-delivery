package main

import (
	"200lab/food-delivery/component"
	"200lab/food-delivery/component/uploadprovider"
	"200lab/food-delivery/middleware"
	"200lab/food-delivery/modules/cart/carttransport/gincart"
	"200lab/food-delivery/modules/food/foodtransport/ginfood"
	"200lab/food-delivery/modules/foodlike/foodtransport/ginfoodlike"
	"200lab/food-delivery/modules/restaurant/restauranttransport/ginrestaurant"
	"200lab/food-delivery/modules/restaurantlike/restaurantliketransport/ginrestaurantlike"
	"200lab/food-delivery/modules/upload/uploadtransport/ginupload"
	"200lab/food-delivery/modules/user/usertransport/ginuser"
	"200lab/food-delivery/modules/useraddress/useraddresstransport/ginuseraddress"
	"200lab/food-delivery/pubsub/pblocal"
	"200lab/food-delivery/socketio"
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

	r := gin.Default()

	realtimeEngine := socketio.NewEngine()

	if err := realtimeEngine.Run(appCtx, r); err != nil {
		log.Fatalln(err)
	}

	if err := subscriber.NewEngine(appCtx).Start(); err != nil {
		log.Fatalln(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// CRUD
	r.StaticFile("/demo/", "./demo.html")

	r.Use(middleware.Recover(appCtx))

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.UploadImage(appCtx))

	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequireAuth(appCtx), ginuser.GetProfileUser(appCtx))

	userAddresses := v1.Group("/user-addresses")
	{
		userAddresses.POST("", middleware.RequireAuth(appCtx), ginuseraddress.CreateUserAddress(appCtx))
		userAddresses.GET("", middleware.RequireAuth(appCtx), ginuseraddress.ListUserAddresses(appCtx))
		userAddresses.GET("/:id", middleware.RequireAuth(appCtx), ginuseraddress.GetUserAddress(appCtx))
		userAddresses.PATCH("/:id", middleware.RequireAuth(appCtx), ginuseraddress.UpdateUserAddress(appCtx))
	}

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
		foods.DELETE("/:id", middleware.RequireAuth(appCtx), ginfood.DeleteFood(appCtx))
		foods.PATCH("/:id", middleware.RequireAuth(appCtx), ginfood.UpdateFood(appCtx))
		foods.POST("/:id/like", middleware.RequireAuth(appCtx), ginfoodlike.LikeFood(appCtx))
		foods.POST("/:id/unlike", middleware.RequireAuth(appCtx), ginfoodlike.UnlikeFood(appCtx))
	}

	carts := v1.Group("/carts")
	{
		carts.POST("", middleware.RequireAuth(appCtx), gincart.CreateCart(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//func startSocketIOServer(engine *gin.Engine, appCtx component.AppContext) {
//	server, _ := socketio.NewServer(&engineio.Options{
//		Transports: []transport.Transport{websocket.Default},
//	})
//
//	server.OnConnect("/", func(s socketio.Conn) error {
//		//s.SetContext("")
//		fmt.Println("connected:", s.ID(), " IP:", s.RemoteAddr())
//
//		//s.Join("Shipper")
//		//server.BroadcastToRoom("/", "Shipper", "test", "Hello 200lab")
//
//		return nil
//	})
//
//	server.OnError("/", func(s socketio.Conn, e error) {
//		fmt.Println("meet error:", e)
//	})
//
//	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
//		fmt.Println("closed", reason)
//		// Remove socket from socket engine (from app context)
//	})
//
//	server.OnEvent("/", "authenticate", func(s socketio.Conn, token string) {
//
//		// Validate token
//		// If false: s.Close(), and return
//
//		// If true
//		// => UserId
//		// Fetch db find user by Id
//		// Here: s belongs to who? (user_id)
//		// We need a map[user_id][]socketio.Conn
//
//		db := appCtx.GetMainDBConnection()
//		store := userstorage.NewSQLStore(db)
//		//
//		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
//		//
//		payload, err := tokenProvider.Validate(token)
//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//		//
//		user, err := store.FindUser(context.Background(), map[string]interface{}{"id": payload.UserId})
//		//
//		if err != nil {
//			s.Emit("authentication_failed", err.Error())
//			s.Close()
//			return
//		}
//
//		if user.Status == 0 {
//			s.Emit("authentication_failed", errors.New("you has been banned/deleted"))
//			s.Close()
//			return
//		}
//
//		user.Mask(false)
//
//		s.Emit("your_profile", user)
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		log.Println(msg)
//	})
//
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//	}
//
//	server.OnEvent("/", "notice", func(s socketio.Conn, p Person) {
//		fmt.Println("server receive notice:", p.Name, p.Age)
//
//		p.Age = 33
//		s.Emit("notice", p)
//
//	})
//
//	server.OnEvent("/", "test", func(s socketio.Conn, msg string) {
//		fmt.Println("server receive test:", msg)
//	})
//	//
//	//server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
//	//	s.SetContext(msg)
//	//	return "recv " + msg
//	//})
//	//
//	//server.OnEvent("/", "bye", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//	//
//	//server.OnEvent("/", "noteSumit", func(s socketio.Conn) string {
//	//	last := s.Context().(string)
//	//	s.Emit("bye", last)
//	//	s.Close()
//	//	return last
//	//})
//
//	go server.Serve()
//
//	engine.GET("/socket.io/*any", gin.WrapH(server))
//	engine.POST("/socket.io/*any", gin.WrapH(server))
//}
