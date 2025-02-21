package base64Captcha

import (
	"github.com/mojocn/base64Captcha"
)

// configJsonBody 配置
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成验证码
func GenerateCaptcha(captchaType string) (idKey, b64s string, err error) {
	var param = configJsonBody{
		Id:          "",
		CaptchaType: captchaType,
		VerifyValue: "",
		DriverAudio: &base64Captcha.DriverAudio{},
		DriverString: &base64Captcha.DriverString{
			Length:          3,
			Height:          47,
			Width:           148,
			ShowLineOptions: 0,
			NoiseCount:      0,
			Source:          "1234567890", // 1234567890qwertyuioplkjhgfdsazxcvbnm
		},
		DriverChinese: &base64Captcha.DriverChinese{},
		DriverMath:    &base64Captcha.DriverMath{},
		DriverDigit:   &base64Captcha.DriverDigit{},
	}
	var driver base64Captcha.Driver

	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	idKey, b64s, _, err = c.Generate()
	return
}

// CaptchaVerify 验证码验证
func CaptchaVerify(idKey string, captcha string) bool {
	return store.Verify(idKey, captcha, true)
}
