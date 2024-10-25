package service

import "os"

func init() {
	// 如果不存在，生成./images目录
	_, err := os.Stat("./images")
	if os.IsNotExist(err) {
		os.Mkdir("./images", os.ModePerm)
	}
}
