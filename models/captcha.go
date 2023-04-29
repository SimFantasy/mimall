package models

import (
	"image/color"

	"github.com/mojocn/base64Captcha"
)

// create store save captcha
var store = base64Captcha.DefaultMemStore

// Generate Captcha
func MakeCaptcha() (string, string, error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Width:           98,
		Height:          36,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          2,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()

	return id, b64s, err
}

func VerifyCaptcha(id, value string) bool {
	if store.Verify(id, value, true) {
		return true
	} else {
		return false
	}
}
