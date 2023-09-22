package main

import "github.com/skip2/go-qrcode"

func main() {

	//qrcode.WriteFile("230621VN1BUPHA", qrcode.High, -1, "./blog_qrcode.png")
	var q *qrcode.QRCode

	q, err := qrcode.New("230621VN1BUPHA", qrcode.High)

	if err != nil {
		//return err
	}
	q.DisableBorder = true
	q.WriteFile(100, "./blog_qrcode.png")
}
