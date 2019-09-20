package main

import (
	"fmt"
	"os"
	"whgo/product/library/utils"
)

// Getwd
//

func main() {
	dir := "export"
	filename := utils.RandomString(8)
	fmt.Println(os.Getwd()) // 获取当前目录
	_, err := os.Stat(dir)  // 查看当前目录下是否存在dir文件or文件夹
	if err != nil {
		fmt.Println(err)
		if os.IsNotExist(err) {
			os.Mkdir(dir, os.ModePerm) // 当前目录创建dir
		}
	}

	createFile, _ := os.Create(filename + ".cvs")
	defer createFile.Close()

}
