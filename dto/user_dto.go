package dto

import "zzy/go-learn/module"

type UserDto struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserDto(user module.User) UserDto {
	userDto := UserDto{
		Name:  user.Name,
		Phone: user.Phone,
	}
	return userDto

}
