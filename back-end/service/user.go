package service

import "back-end/model"

func CreateUser(user model.User) error {
	return model.DB.Create(&user).Error
}

func GetUser(id uint) (model.User, error) {
	var user model.User
	err := model.DB.Where("id = ?", id).First(&user).Error
	return user, err
}

func GetUserByOpenId(openId string) (model.User, error) {
	var user model.User
	err := model.DB.Where("wx_open_id = ?", openId).First(&user).Error
	return user, err
}

func UpdateUser(user model.User) error {
	return model.DB.Model(&model.User{}).Updates(user).Error
}

func ListUsers() ([]model.User, error) {
	var users []model.User
	err := model.DB.Find(&users).Error
	return users, err
}

func DeleteUser(id uint) error {
	return model.DB.Delete(&model.User{}, id).Error
}
