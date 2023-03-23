package encryption

import (
	"bytes"
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

var (
	iv       = []byte("A05B16C27D38E49F")
	pemStart = []byte("\n-----BEGIN ")
)

type PKCSType string

const (
	RsaPKCS1 PKCSType = "PKCS#1"
	RsaPKCS8 PKCSType = "PKCS#8"
)

//AESEncrypt Aes加密
func AESEncrypt(msg, key []byte) []byte {
	encrypted, _ := AESEncryptE(msg, key)
	return encrypted
}

//AESEncryptE Aes加密
func AESEncryptE(msg, key []byte) ([]byte, error) {
	//获取block块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//补码
	msg = PKCS7Padding(msg, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//创建明文长度的数组
	encrypted := make([]byte, len(msg))
	//加密明文
	blockMode.CryptBlocks(encrypted, msg)
	return encrypted, nil
}

//AESEncryptIv Aes加密
func AESEncryptIv(msg, key, iv []byte) ([]byte, error) {
	//获取block块
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//补码
	msg = PKCS7Padding(msg, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//创建明文长度的数组
	encrypted := make([]byte, len(msg))
	//加密明文
	blockMode.CryptBlocks(encrypted, msg)
	return encrypted, nil
}

func PKCS7Padding(origData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(origData)%blockSize
	patent := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, patent...)
}

//AESDecrypt Aes解密
func AESDecrypt(msg, key []byte) []byte {
	decrypt, _ := AESDecryptE(msg, key)
	return decrypt
}

//AESDecryptE Aes解密
func AESDecryptE(msg, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(msg))
	blockMode.CryptBlocks(origData, msg)
	return PKCS7UnPadding(origData), nil
}

//AESDecryptIv Aes解密
func AESDecryptIv(msg, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(msg))
	blockMode.CryptBlocks(origData, msg)
	return PKCS7UnPadding(origData), nil
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	return origData[:length-int(origData[length-1])]
}

//RSAEncrypt 使用公钥进行加密
func RSAEncrypt(public []byte, msg []byte) []byte {
	cipherText, _ := RSAEncryptE(public, msg)
	return cipherText
}

//RSAEncryptPKCS1 使用公钥进行加密
func RSAEncryptPKCS1(public []byte, msg []byte) []byte {
	cipherText, _ := RSAEncryptPKCS1E(public, msg)
	return cipherText
}

//RSAEncryptE 使用公钥进行加密
func RSAEncryptE(public []byte, msg []byte) ([]byte, error) {
	if bytes.HasPrefix(public, pemStart[1:]) {
		return RSAEncryptPemE(public, msg)
	} else {
		return RSAEncryptStringE(public, msg)
	}
}

//RSAEncryptPKCS1E 使用公钥进行加密
func RSAEncryptPKCS1E(public []byte, msg []byte) ([]byte, error) {
	if bytes.HasPrefix(public, pemStart[1:]) {
		return RSAEncryptPemE(public, msg)
	} else {
		return RSAEncryptStringE(public, msg)
	}
}

//RSAEncryptPem 使用公钥（pem格式）进行加密
func RSAEncryptPem(public []byte, msg []byte) []byte {
	cipherText, _ := RSAEncryptPemE(public, msg)
	return cipherText
}

//RSAEncryptPemE 使用公钥（pem格式）进行加密
func RSAEncryptPemE(public []byte, msg []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode(public)
	return rsaEncrypt(block.Bytes, msg, RsaPKCS8)
}

//RSAEncryptPKCS1Pem 使用公钥（pem格式）进行加密
func RSAEncryptPKCS1Pem(public []byte, msg []byte) []byte {
	cipherText, _ := RSAEncryptPKCS1PemE(public, msg)
	return cipherText
}

//RSAEncryptPKCS1PemE 使用公钥（pem格式）进行加密
func RSAEncryptPKCS1PemE(public []byte, msg []byte) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode(public)
	return rsaEncrypt(block.Bytes, msg, RsaPKCS1)
}

//RSAEncryptString 使用公钥（String格式）进行加密
func RSAEncryptString(public, msg []byte) []byte {
	cipherText, _ := RSAEncryptStringE(public, msg)
	return cipherText
}

//RSAEncryptStringE 使用公钥（string格式）进行加密
func RSAEncryptStringE(public, msg []byte) ([]byte, error) {
	//base64解码
	publicBytes := make([]byte, base64.StdEncoding.DecodedLen(len(public)))
	n, err := base64.StdEncoding.Decode(publicBytes, public)
	if err != nil {
		return nil, err
	}
	return rsaEncrypt(publicBytes[:n], msg, RsaPKCS8)
}

//RSAEncryptPKCS1String 使用公钥（String格式）进行加密
func RSAEncryptPKCS1String(public, msg []byte) []byte {
	cipherText, _ := RSAEncryptPKCS1StringE(public, msg)
	return cipherText
}

//RSAEncryptPKCS1StringE 使用公钥（string格式）进行加密
func RSAEncryptPKCS1StringE(public, msg []byte) ([]byte, error) {
	//base64解码
	publicBytes := make([]byte, base64.StdEncoding.DecodedLen(len(public)))
	n, err := base64.StdEncoding.Decode(publicBytes, public)
	if err != nil {
		return nil, err
	}
	return rsaEncrypt(publicBytes[:n], msg, RsaPKCS1)
}

func rsaEncrypt(public []byte, msg []byte, pkcs PKCSType) ([]byte, error) {
	var pubKey *rsa.PublicKey
	//x509解码,得到一个interface类型的pub
	switch pkcs {
	case RsaPKCS1:
		pub, err := x509.ParsePKCS1PublicKey(public)
		if err != nil {
			return nil, err
		}
		pubKey = pub
	case RsaPKCS8:
		pub, err := x509.ParsePKIXPublicKey(public)
		if err != nil {
			return nil, err
		}
		pubKey = pub.(*rsa.PublicKey)
	default:
		return nil, errors.New("encrypt pkcs error")
	}

	//加密操作,需要将接口类型的pub进行类型断言得到公钥类型
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, msg)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

//RSADecrypt 使用私钥进行解密
func RSADecrypt(private []byte, cipherText []byte) []byte {
	afterDecrypted, _ := RSADecryptE(private, cipherText)
	return afterDecrypted
}

//RSADecryptPKCS1 使用私钥进行解密
func RSADecryptPKCS1(private []byte, cipherText []byte) []byte {
	afterDecrypted, _ := RSADecryptPKCS1E(private, cipherText)
	return afterDecrypted
}

//RSADecryptE 使用私钥进行解密
func RSADecryptE(private []byte, cipherText []byte) ([]byte, error) {
	if bytes.HasPrefix(private, pemStart[1:]) {
		return RSADecryptPemE(private, cipherText)
	} else {
		return RSADecryptStringE(private, cipherText)
	}
}

//RSADecryptPKCS1E 使用私钥进行解密
func RSADecryptPKCS1E(private []byte, cipherText []byte) ([]byte, error) {
	if bytes.HasPrefix(private, pemStart[1:]) {
		return RSADecryptPKCS1PemE(private, cipherText)
	} else {
		return RSADecryptPKCS1StringE(private, cipherText)
	}
}

//RSADecryptPem 使用私钥（pem格式）进行解密
func RSADecryptPem(private []byte, cipherText []byte) []byte {
	afterDecrypted, _ := RSADecryptPemE(private, cipherText)
	return afterDecrypted
}

//RSADecryptPKCS1Pem 使用私钥（pem格式）进行解密
func RSADecryptPKCS1Pem(private []byte, cipherText []byte) []byte {
	afterDecrypted, _ := RSADecryptPKCS1PemE(private, cipherText)
	return afterDecrypted
}

//RSADecryptString 使用私钥（String格式）进行解密
func RSADecryptString(private, cipherText []byte) []byte {
	afterDecrypted, _ := RSADecryptStringE(private, cipherText)
	return afterDecrypted
}

//RSADecryptPKCS1String 使用私钥（String格式）进行解密
func RSADecryptPKCS1String(private, cipherText []byte) []byte {
	afterDecrypted, _ := RSADecryptPKCS1StringE(private, cipherText)
	return afterDecrypted
}

//RSADecryptPemE 使用私钥（pem格式）进行解密
func RSADecryptPemE(private []byte, cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(private)
	return rsaDecrypt(block.Bytes, cipherText, RsaPKCS8)
}

//RSADecryptStringE 使用私钥（String格式）进行解密
func RSADecryptStringE(private, cipherText []byte) ([]byte, error) {
	privateBytes := make([]byte, base64.StdEncoding.DecodedLen(len(private)))
	n, err := base64.StdEncoding.Decode(privateBytes, private)
	if err != nil {
		return nil, err
	}
	return rsaDecrypt(privateBytes[:n], cipherText, RsaPKCS8)
}

//RSADecryptPemE 使用私钥（pem格式）进行解密
func RSADecryptPKCS1PemE(private []byte, cipherText []byte) ([]byte, error) {
	block, _ := pem.Decode(private)
	return rsaDecrypt(block.Bytes, cipherText, RsaPKCS1)
}

//RSADecryptStringE 使用私钥（String格式）进行解密
func RSADecryptPKCS1StringE(private, cipherText []byte) ([]byte, error) {
	privateBytes := make([]byte, base64.StdEncoding.DecodedLen(len(private)))
	n, err := base64.StdEncoding.Decode(privateBytes, private)
	if err != nil {
		return nil, err
	}
	return rsaDecrypt(privateBytes[:n], cipherText, RsaPKCS1)
}

func rsaDecrypt(private []byte, cipherText []byte, pkcs PKCSType) ([]byte, error) {
	var privateKey *rsa.PrivateKey
	//x509解码,得到一个interface类型的pub
	switch pkcs {
	case RsaPKCS1:
		priv, err := x509.ParsePKCS1PrivateKey(private)
		if err != nil {
			return nil, err
		}
		privateKey = priv
	case RsaPKCS8:
		priv, err := x509.ParsePKCS8PrivateKey(private)
		if err != nil {
			return nil, err
		}
		privateKey = priv.(*rsa.PrivateKey)
	default:
		return nil, errors.New("decrypt pkcs error")
	}
	//二次解码完毕，调用解密函数
	afterDecrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return nil, err
	}
	return afterDecrypted, nil
}

//RSAVerifySHA256Pem 验证数据是否被修改过
func RSAVerifySHA256Pem(pubKey, msg, sign []byte) bool {
	block, _ := pem.Decode(pubKey)
	return rsaVerify(block.Bytes, msg, sign, RsaPKCS8)
}

//RSAVerifyPKCS1SHA256Pem 验证数据是否被修改过
func RSAVerifyPKCS1SHA256Pem(pubKey, msg, sign []byte) bool {
	block, _ := pem.Decode(pubKey)
	return rsaVerify(block.Bytes, msg, sign, RsaPKCS1)
}

//RSAVerify 验证数据是否被修改过
func RSAVerifySHA256String(pubKey, msg, sign []byte) bool {
	publicBytes := make([]byte, base64.StdEncoding.DecodedLen(len(pubKey)))
	n, err := base64.StdEncoding.Decode(publicBytes, pubKey)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return rsaVerify(publicBytes[:n], msg, sign, RsaPKCS8)
}

//RSAVerifyPKCS1SHA256String 验证数据是否被修改过
func RSAVerifyPKCS1SHA256String(pubKey, msg, sign []byte) bool {
	publicBytes := make([]byte, base64.StdEncoding.DecodedLen(len(pubKey)))
	n, err := base64.StdEncoding.Decode(publicBytes, pubKey)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return rsaVerify(publicBytes[:n], msg, sign, RsaPKCS1)
}

//RSAVerify 验证数据是否被修改过
func RSAVerifySHA256(pubKey, msg, sign []byte) bool {
	if bytes.HasPrefix(pubKey, pemStart[1:]) {
		return RSAVerifySHA256Pem(pubKey, msg, sign)
	} else {
		return RSAVerifySHA256String(pubKey, msg, sign)
	}
}

//RSAVerifyPKCS1SHA256 验证数据是否被修改过
func RSAVerifyPKCS1SHA256(pubKey, msg, sign []byte) bool {
	if bytes.HasPrefix(pubKey, pemStart[1:]) {
		return RSAVerifyPKCS1SHA256Pem(pubKey, msg, sign)
	} else {
		return RSAVerifyPKCS1SHA256String(pubKey, msg, sign)
	}
}

func rsaVerify(public, msg, sign []byte, pkcs PKCSType) bool {
	var pubKey *rsa.PublicKey
	//x509解码,得到一个interface类型的pub
	switch pkcs {
	case RsaPKCS1:
		pub, err := x509.ParsePKCS1PublicKey(public)
		if err != nil {
			return false
		}
		pubKey = pub
	case RsaPKCS8:
		pub, err := x509.ParsePKIXPublicKey(public)
		if err != nil {
			return false
		}
		pubKey = pub.(*rsa.PublicKey)
	default:
		return false
	}

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return false
	}
	msgHashSum := msgHash.Sum(nil)
	if err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, msgHashSum, sign); err != nil {
		return false
	}
	return true
}

//RSASignSHA256Pem 使用私钥签名
func RSASignSHA256Pem(privateKey, msg []byte) []byte {
	block, _ := pem.Decode(privateKey)
	return rsaSign(block.Bytes, msg, RsaPKCS8)
}

//RSASignSHA256String 使用私钥签名
func RSASignSHA256String(privateKey, msg []byte) []byte {
	privateBytes := make([]byte, base64.StdEncoding.DecodedLen(len(privateKey)))
	n, err := base64.StdEncoding.Decode(privateBytes, privateKey)
	if err != nil {
		return nil
	}
	return rsaSign(privateBytes[:n], msg, RsaPKCS8)
}

//RSASignSHA256 使用私钥签名
func RSASignSHA256(privateKey, msg []byte) []byte {
	if bytes.HasPrefix(privateKey, pemStart[1:]) {
		return RSASignSHA256Pem(privateKey, msg)
	} else {
		return RSASignSHA256String(privateKey, msg)
	}
}

//RSASignPKCS1SHA256Pem 使用私钥签名
func RSASignPKCS1SHA256Pem(privateKey, msg []byte) []byte {
	block, _ := pem.Decode(privateKey)
	return rsaSign(block.Bytes, msg, RsaPKCS1)
}

//RSASignPKCS1SHA256String 使用私钥签名
func RSASignPKCS1SHA256String(privateKey, msg []byte) []byte {
	privateBytes := make([]byte, base64.StdEncoding.DecodedLen(len(privateKey)))
	n, err := base64.StdEncoding.Decode(privateBytes, privateKey)
	if err != nil {
		return nil
	}
	return rsaSign(privateBytes[:n], msg, RsaPKCS1)
}

//RSASignPKCS1SHA256 使用私钥签名
func RSASignPKCS1SHA256(privateKey, msg []byte) []byte {
	if bytes.HasPrefix(privateKey, pemStart[1:]) {
		return RSASignPKCS1SHA256Pem(privateKey, msg)
	} else {
		return RSASignPKCS1SHA256String(privateKey, msg)
	}
}

func rsaSign(private, msg []byte, pkcs PKCSType) []byte {
	var privateKey *rsa.PrivateKey
	//x509解码,得到一个interface类型的pub
	switch pkcs {
	case RsaPKCS1:
		priv, err := x509.ParsePKCS1PrivateKey(private)
		if err != nil {
			return nil
		}
		privateKey = priv
	case RsaPKCS8:
		priv, err := x509.ParsePKCS8PrivateKey(private)
		if err != nil {
			return nil
		}
		privateKey = priv.(*rsa.PrivateKey)
	default:
		return nil
	}

	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return nil
	}
	msgHashSum := msgHash.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, msgHashSum)
	if err != nil {
		return nil
	}
	return sign
}

func GenRSAKeys(bits int) (private, public []byte, err error) {
	var privateBuf, publicBuf bytes.Buffer
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return
	}
	x509PrivateKey, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return
	}
	//将私钥字符串设置到pem格式块中
	privateBlock := pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509PrivateKey,
	}
	err = pem.Encode(&privateBuf, &privateBlock)
	if err != nil {
		return
	}
	publicKey := privateKey.PublicKey
	x509PublicKey, _ := x509.MarshalPKIXPublicKey(&publicKey)
	publicBlock := pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509PublicKey,
	}
	err = pem.Encode(&publicBuf, &publicBlock)
	if err != nil {
		return
	}
	return privateBuf.Bytes(), publicBuf.Bytes(), nil
}
