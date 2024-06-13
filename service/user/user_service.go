package user_service

import (
	"fmt"
	"hotel-booking-golang-gin/initializers"
	user_interface "hotel-booking-golang-gin/interfaces/user"
	crypto_util "hotel-booking-golang-gin/util/crypto"
	error_handler "hotel-booking-golang-gin/util/error"
)

func CreateUser(userData user_interface.User) (user_interface.User, *error_handler.ErrArg) {
	existingUser := initializers.DB.Where(&user_interface.User{Email: userData.Email}).First(&user_interface.User{})
	fmt.Println("======", existingUser.RowsAffected)
	if existingUser.RowsAffected > 0 {
		return user_interface.User{}, &error_handler.ErrArg{
			Code:        "USER_ALREADY_EXISTS",
			Description: "User already exists",
			Title:       "User already exists",
		}
	}
	userData.Password = crypto_util.HashPassword(userData.Password)
	initializers.DB.Create(&userData)
	return userData, nil
}
