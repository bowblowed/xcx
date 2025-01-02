package service

import (
	"back-end/model"
	"mime/multipart"
)

func CreateProduct(Name string, CategoryId uint, DefaultPrice uint,
	Count uint32, Description string, picFile *multipart.FileHeader,
	DescriptionFiles []*multipart.FileHeader,
	Props map[string][]string) error {
	product := model.Product{
		Name:         Name,
		CategoryId:   CategoryId,
		DefaultPrice: DefaultPrice,
		Count:        Count,
		Description:  Description,
		Props:        Props,
	}
	if picFile != nil {
		picId, err := CreateFile(picFile, "product")
		if err != nil {
			return err
		}
		product.PicId = picId
	}
	for i := 0; i < len(DescriptionFiles); i++ {
		picId, err := CreateFile(DescriptionFiles[i], "product_description")
		if err != nil {
			return err
		}
		product.DescriptionPicIds = append(product.DescriptionPicIds, picId)
	}
	return model.DB.Create(&product).Error

}

func UpdateProduct(Id uint, Name string, CategoryId uint, DefaultPrice uint,
	Count uint32, Description string, picFile *multipart.FileHeader,
	DescriptionFiles []*multipart.FileHeader,
	Props map[string][]string) error {
	var product model.Product
	product, err := GetProduct(Id)
	if err != nil {
		return err
	}
	product.Name = Name
	product.CategoryId = CategoryId
	product.DefaultPrice = DefaultPrice
	product.Count = Count
	product.Description = Description
	product.Props = Props
	if picFile != nil {
		if product.PicId != 0 {
			DeleteFile(product.PicId)
		}
		product.PicId, err = CreateFile(picFile, "product")
	}
	if err != nil {
		return err
	}
	for i := 0; i < len(product.DescriptionPicIds); i++ {
		DeleteFile(product.DescriptionPicIds[i])
	}
	product.DescriptionPicIds = []uint{}
	for i := 0; i < len(DescriptionFiles); i++ {
		picId, _ := CreateFile(DescriptionFiles[i], "product_description")
		product.DescriptionPicIds = append(product.DescriptionPicIds, picId)
	}
	return model.DB.Save(&product).Error
}

func DeleteProduct(ID uint) error {
	product, err := GetProduct(ID)
	if err != nil {
		return err
	}
	if product.PicId != 0 {
		DeleteFile(product.PicId)
	}
	for _, v := range product.DescriptionPicIds {
		DeleteFile(v)
	}
	return model.DB.Delete(&product).Error
}

func GetProduct(id uint) (model.Product, error) {
	var product2 model.Product
	err := model.DB.Where("id = ?", id).First(&product2).Error
	return product2, err
}

func ListProducts() ([]model.Product, error) {
	var products []model.Product
	err := model.DB.Find(&products).Error
	return products, err
}
func ListProductsByCategory(categoryId uint) ([]model.Product, error) {
	var products []model.Product
	err := model.DB.Where("category_id = ?", categoryId).Find(&products).Error
	return products, err
}

func ProductCaculatePrice(productId uint, propPair map[string]string, count uint) (uint, error) {
	product, err := GetProduct(productId)
	if err != nil {
		return 0, err
	}
	priceTaglist, err := ListPriceTagByProductId(productId)
	if err != nil {
		return 0, err
	}
	isSubset := func(subset, set map[string]string) bool {
		for k, v := range subset {
			if v2, ok := set[k]; !ok || v != v2 {
				return false
			}
		}
		return true
	}
	for _, v := range priceTaglist {
		if isSubset(propPair, v.PropPairs) {
			return v.Price * count, nil
		}
	}
	return product.DefaultPrice * count, nil
}

func ListProductsByIds(IDs []uint) ([]model.Product, error) {
	var products []model.Product
	err := model.DB.Find(&products, IDs).Error
	return products, err
}
