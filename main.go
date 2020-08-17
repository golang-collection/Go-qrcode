package main

import (
	"GO-qrcode/tools"
	"fmt"
	"image/png"
	"os"
)

/**
* @Author: super
* @Date: 2020-08-17 11:06
* @Description:
**/

func main() {
	str := "http://superssssss.cn/"

	bytes, _ := tools.GenerateQRCodeByte(str)
	fmt.Println(bytes)

	barcode, _ := tools.GenerateQRCode("http://superssssss.cn/")
	file, _ := os.Create("qr.png")
	defer file.Close()

	png.Encode(file, barcode)
}