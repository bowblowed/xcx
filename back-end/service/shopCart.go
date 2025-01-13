package service

import "back-end/model"

func DefaultShopCart(userId uint) (uint, error) {
	var s model.ShopCart = model.ShopCart{
		UserId: userId,
	}
	err := model.DB.Create(&s).Error
	if err != nil {
		return 0, err
	}
	return s.ID, nil
}

func GetShopCart(id uint) (model.ShopCart, error) {
	var shopcar model.ShopCart
	err := model.DB.First(&shopcar, id).Error
	return shopcar, err
}

func UpdateShopCart(id uint, shopcart model.ShopCart) error {
	return model.DB.Model(&model.ShopCart{}).Where("id = ?", id).Updates(shopcart).Error
}

func ShopCartInsertItems(id uint, items model.OrderItems) error {
	shopCart, err := GetShopCart(id)
	if err != nil {
		return err
	}
	checkMapEqual := func(a, b map[string]string) bool {
		for i, v := range a {
			if v2, ok := b[i]; !ok || v != v2 {
				return false
			}
		}
		return true
	}
	for _, vi := range items {
		isInsert := false
		for _, vj := range shopCart.OrderItems {
			if vi.ProductId == vj.ProductId && checkMapEqual(vj.Props, vi.Props) {
				vj.Count += vi.Count
				isInsert = true
				break
			}
		}
		if !isInsert {
			shopCart.OrderItems = append(shopCart.OrderItems, vi)
		}
	}
	return UpdateShopCart(id, shopCart)
}

func ShopCartDeleteItems(id uint, idxs uint) error {
	shopCart, err := GetShopCart(id)
	if err != nil {
		return err
	}
	if idxs >= uint(len(shopCart.OrderItems)) {
		return nil
	}
	shopCart.OrderItems = append(shopCart.OrderItems[:idxs], shopCart.OrderItems[idxs+1:]...)
	return UpdateShopCart(id, shopCart)
}

func ShopCartChangeCount(id uint, idxs uint, count uint) error {
	shopCart, err := GetShopCart(id)
	if err != nil {
		return err
	}
	shopCart.OrderItems[idxs].Count = count
	return UpdateShopCart(id, shopCart)
}
