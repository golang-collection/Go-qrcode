package tools

/**
* @Author: super
* @Date: 2020-08-17 10:58
* @Description:
**/

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/skip2/go-qrcode"
)

func GenerateQRCode(str string) (barcode.Barcode, error) {
	qrCode, _ := qr.Encode(str, qr.M, qr.Auto)

	qrCode, err := barcode.Scale(qrCode, 256, 256)
	return qrCode, err
}

func GenerateQRCodeByte(str string)([]byte, error){
	return qrcode.Encode(str, qrcode.Highest, 256)
}