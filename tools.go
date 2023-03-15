package utils

import (
	"io"
	"mime/multipart"
	"os"
	"regexp"
)

// 校验手机号
func VerifyMobileFormat(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

//bool转[]bytez
func BoolIntoByte(b bool) []byte {
	if b {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

//byte转bool
func ByteIntoBool(b []byte) bool {
	if b[0] == 1 {
		return true
	} else {
		return false
	}
}

//byte转int
func ByteIntoInt(b []byte) uint8 {
	if b[0] == 1 {
		return 1
	} else {
		return 0
	}
}

//bool转int
func BoolIntoInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

//用于转化前端字符串布尔值为[]byte
func StrBoolIntoByte(s string) []byte {
	if s == "true" {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

//用于转化前端字符串布尔值为[]byte
func StrGenderIntoByte(s string) []byte {
	if s == "男" {
		return []byte{1}
	} else {
		return []byte{0}
	}
}

// 启用未启用状态布尔转字符串
func ByteEnabledToString(b []byte) string {
	if b[0] == 0 {
		return `禁用`
	} else {
		return `启用`
	}
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	err := os.MkdirAll("./static/uploadfile/", os.ModePerm)
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}
