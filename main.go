package main

import (
	hotel_controller "hotel-booking-golang-gin/controllers/hotel"
	user_controller "hotel-booking-golang-gin/controllers/user"
	"hotel-booking-golang-gin/initializers"
	hotel_middleware "hotel-booking-golang-gin/middleware/hotel"
	user_middlewares "hotel-booking-golang-gin/middleware/user"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.InitDatabase()
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	{
		v1 := r.Group("/api/v1")

		v1_user := v1.Group("/user")

		v1_user.POST("/create", user_middlewares.UserValidator, user_controller.CreateUser)
		v1_user.POST("/login", user_middlewares.UserLoginValidator, user_controller.LoginUser)

		v1_hotel := v1.Group("/hotel")

		v1_hotel.POST("/create",
			user_middlewares.Authenticate,
			user_middlewares.ForRole("MERCHANT"),
			hotel_middleware.HotelValidator,
			hotel_controller.CreateHotel,
		)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}
