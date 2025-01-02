package service

import (
	"back-end/config"
	"back-end/model"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

var baseUrl = fmt.Sprintf("https://%s", config.Server.Domain)

func CreateFile(file *multipart.FileHeader, _type string) (uint, error) {
	if file == nil {
		return 0, errors.New("文件为空")
	}
	name := fmt.Sprintf("%s_%d", _type, time.Now().UnixNano())
	url := baseUrl + "/file/" + name
	path := fmt.Sprintf("./file/%s", name)
	src, err := file.Open()
	if err != nil {
		return 0, err
	}
	defer src.Close()

	if err = os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return 0, err
	}
	out, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return 0, err
	}
	var pic model.File = model.File{
		Name: name, Path: path, Url: url, Type: _type,
	}
	err = model.DB.Create(&pic).Error
	return pic.ID, err
}

func DeleteFile(id uint) error {
	var pic model.File
	err := model.DB.First(&pic, id).Error
	if err != nil {
		return err
	}
	err = os.Remove(pic.Path)
	if err != nil {
		return err
	}
	return model.DB.Delete(&pic, id).Error
}

func GetFile(id uint) (model.File, error) {
	if id == 0 {
		return model.File{}, errors.New("pic 不存在")
	}
	var pic model.File
	err := model.DB.First(&pic, id).Error
	return pic, err
}

func UpdateFile(id uint, file *multipart.FileHeader) error {
	var pic model.File
	err := model.DB.First(&pic, id).Error
	if err != nil {
		return err
	}
	path := pic.Path
	err = os.Remove(path)
	if err != nil {
		return err
	}
	name := fmt.Sprintf("%s_%d", pic.Type, time.Now().Unix())
	url := baseUrl + "/file/" + name
	path = fmt.Sprintf("./file/%s", name)
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	if err = os.MkdirAll(filepath.Dir(path), 0750); err != nil {
		return err
	}
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	if err != nil {
		return err
	}
	return model.DB.Model(&pic).Updates(model.File{Name: name, Path: path, Url: url}).Error
}
