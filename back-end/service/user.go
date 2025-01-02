package service

import "back-end/model"

func CreateUser(user model.User) error {
	return model.DB.Create(&user).Error
}
func CreateWxUserDefault(openId string) error {
	return CreateUser(
		model.User{
			WxOpenId: openId,
			Type:     model.UserTypeUserUnsigned,
		},
	)
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

func UpdateUser(id uint, user model.User) error {
	return model.DB.Model(&model.User{}).Where("id = ?", id).Updates(&user).Error
}

func ListUsers() ([]model.User, error) {
	var users []model.User
	err := model.DB.Not("type =?", model.UserTypeAdmin).Find(&users).Error
	return users, err
}

func DeleteUser(id uint) error {
	return model.DB.Delete(&model.User{}, id).Error
}

func GetUserShopCart(user model.User) (model.ShopCart, error) {
	var err error
	shopCartId := user.ShopCartId
	if shopCartId == 0 {
		user.ShopCartId, err = DefaultShopCart(user.ID)
		if err != nil {
			return model.ShopCart{}, err
		}
	}
	err = UpdateUser(user.ID, user)
	if err != nil {
		return model.ShopCart{}, err
	}
	return GetShopCart(user.ShopCartId)
}

func UpdateUserShopCart(user model.User, shopCart model.ShopCart) error {
	return UpdateShopCart(user.ShopCartId, shopCart)
}

func UserShopCartInsertItems(user model.User, items []model.OrderItem) error {
	if user.ShopCartId == 0 {
		id, err := DefaultShopCart(user.ID)
		if err != nil {
			return err
		}
		user.ShopCartId = id
	}
	return ShopCartInsertItems(user.ShopCartId, items)
}

func UserShopCartDeleteItem(user model.User, idx uint) error {
	return ShopCartDeleteItems(user.ShopCartId, idx)
}

func UserShopCartChangeCount(user model.User, idx uint, count uint) error {
	return ShopCartChangeCount(user.ShopCartId, idx, count)
}
