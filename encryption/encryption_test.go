package encryption

import (
	"fmt"
	"testing"
)

var (
	aesKey    = "e10adc3949ba59abbe56e057f20f883e"
	publicPem = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3zWKtpGquMeswY51kVpe
kQmPihtgvVRE3o5hyR9V9663EEexNMCq9pYlUOkpyTKQWiS+R9EcBVAzQKKjfns7
TqRA6vrjkybyQjizuAULBgJZkMDmMikfDGkHuVbRaDiyr5xDNfMnqb1QmaFK6aqu
WUuXopjxZQUDMaGq41GyIGJfIDdTiPvs9V5t/aTFZZ9VBEsTLtbbgemV+Dzagtq6
uMUmfm272RFfTw0M7Q6jPwNF1PurCqOaja17+2oH671BdvupmtSWj/ysuvo5i220
kOeHpe2PW5fOwhjEBoARz9qRlPyGADNqhmupUJgsVwavwT45/ns3/4ECwy2Z/Oqp
UQIDAQAB
-----END PUBLIC KEY-----`
	privatePem = `-----BEGIN PRIVATE KEY-----
MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDfNYq2kaq4x6zB
jnWRWl6RCY+KG2C9VETejmHJH1X3rrcQR7E0wKr2liVQ6SnJMpBaJL5H0RwFUDNA
oqN+eztOpEDq+uOTJvJCOLO4BQsGAlmQwOYyKR8MaQe5VtFoOLKvnEM18yepvVCZ
oUrpqq5ZS5eimPFlBQMxoarjUbIgYl8gN1OI++z1Xm39pMVln1UESxMu1tuB6ZX4
PNqC2rq4xSZ+bbvZEV9PDQztDqM/A0XU+6sKo5qNrXv7agfrvUF2+6ma1JaP/Ky6
+jmLbbSQ54el7Y9bl87CGMQGgBHP2pGU/IYAM2qGa6lQmCxXBq/BPjn+ezf/gQLD
LZn86qlRAgMBAAECggEAXddjzyfkz7TbE7EyCn8k3PNZDqY+ALtKDQttXyI2gWBK
BKUvgk8yUWOEosS0VatVdyCVgxoOHsm+EoGOHSHwKDr0NBZw0TABkAfIvygXn1oj
j7EZ1qLlM9GAdtJiTd/wd9ZJh4gH99vo4/4kwaKJlG8sfCay7layJg6jCtkQ+323
AuZR5Ey7shAQdlV8hd2QlPIkgZ1Qxl7zTE9kbu603xgy7UUc0sdIumGok0U3IDIW
sRu5FyQvfkOkX3act4+HKAfpIHu04z009z53IBysP4bPPMyR6OUAwlwDv9p/KYyZ
yrACF6s/UusbduU75pDoXL+slcY+Xw3h5zdrGP5yNQKBgQD3mOWgspVLjZNd/7m9
oqMplpNrjII+ll44cu34WZl3/N8JBFCpY7GL/MS7rNkBHtxCiXt+iIk3rIcfdxB+
Ep8Um19SHzNWDc9lztSDnxBEIDwyLgq8ZVyeGwM3/sUM5coqes4vR91FllyWowVl
Z1SHd0NpOBpHKBae2lOkYnjtUwKBgQDmyMNebj2szF9dDJuhEn0nrfMt4np/sWS3
kYdncxATMV3eeBjVnu7LI5TXCQWVvJOn3ireoXoR/Hu2rXtic68JZcK9oGWLG6ES
N/pawxrPDOZmbaiBifPbBZF9iYRNos/rM59EDpx857YdmFrm3NiMmYySkPcagnkl
Iy2wSYgWSwKBgQCUCzXx+hXOkR52a/uRafeJXGFVQzR1W8+GjSHBB0H+PpLVf2ED
h3qXq1GurdErHFqy5iWTKOxy/5PjunuABn5cAfLi1YkKgUNrKlpjQO6m7WdQ0brV
nv3cA2bPSmRkK9nPwJIcLyoc+6Ts5RjK4xflz+fcuBDuTwgHf5/XgJGxxQKBgEaG
FUW3FHvInF+36ZsIJlzIyGY+Lkb1M5zxCNqxC7mBvXitZJFrFvfWsLYhmKZQUSLc
Pmd4GvCCImAJlYnZZiTOBvuhl0YyND94AbZYpsmc3I9ydSrYuD827WXabya2WSk7
xjCq8mEtAbb7htPDqx0UEyP6zrg4JMMXg++6GK8PAoGBAMmKfqvZEzuO14rLGeMo
MJzt4S440CnvcbxX8KTHB7/vvNjWYBoc8Lu96zFv/G9nZ8l/YM8vm03anxqMAvL5
iaGV6sUs251Q5dDncCYJ4oLTMatqONlHvdRRTtuZwV6fOUigiPRardj8AnQSzcS1
Qr9kuTwDg7yJcrBx20KTupob
-----END PRIVATE KEY-----`

	public  = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3zWKtpGquMeswY51kVpekQmPihtgvVRE3o5hyR9V9663EEexNMCq9pYlUOkpyTKQWiS+R9EcBVAzQKKjfns7TqRA6vrjkybyQjizuAULBgJZkMDmMikfDGkHuVbRaDiyr5xDNfMnqb1QmaFK6aquWUuXopjxZQUDMaGq41GyIGJfIDdTiPvs9V5t/aTFZZ9VBEsTLtbbgemV+Dzagtq6uMUmfm272RFfTw0M7Q6jPwNF1PurCqOaja17+2oH671BdvupmtSWj/ysuvo5i220kOeHpe2PW5fOwhjEBoARz9qRlPyGADNqhmupUJgsVwavwT45/ns3/4ECwy2Z/OqpUQIDAQAB`
	private = `MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDfNYq2kaq4x6zBjnWRWl6RCY+KG2C9VETejmHJH1X3rrcQR7E0wKr2liVQ6SnJMpBaJL5H0RwFUDNAoqN+eztOpEDq+uOTJvJCOLO4BQsGAlmQwOYyKR8MaQe5VtFoOLKvnEM18yepvVCZoUrpqq5ZS5eimPFlBQMxoarjUbIgYl8gN1OI++z1Xm39pMVln1UESxMu1tuB6ZX4PNqC2rq4xSZ+bbvZEV9PDQztDqM/A0XU+6sKo5qNrXv7agfrvUF2+6ma1JaP/Ky6+jmLbbSQ54el7Y9bl87CGMQGgBHP2pGU/IYAM2qGa6lQmCxXBq/BPjn+ezf/gQLDLZn86qlRAgMBAAECggEAXddjzyfkz7TbE7EyCn8k3PNZDqY+ALtKDQttXyI2gWBKBKUvgk8yUWOEosS0VatVdyCVgxoOHsm+EoGOHSHwKDr0NBZw0TABkAfIvygXn1ojj7EZ1qLlM9GAdtJiTd/wd9ZJh4gH99vo4/4kwaKJlG8sfCay7layJg6jCtkQ+323AuZR5Ey7shAQdlV8hd2QlPIkgZ1Qxl7zTE9kbu603xgy7UUc0sdIumGok0U3IDIWsRu5FyQvfkOkX3act4+HKAfpIHu04z009z53IBysP4bPPMyR6OUAwlwDv9p/KYyZyrACF6s/UusbduU75pDoXL+slcY+Xw3h5zdrGP5yNQKBgQD3mOWgspVLjZNd/7m9oqMplpNrjII+ll44cu34WZl3/N8JBFCpY7GL/MS7rNkBHtxCiXt+iIk3rIcfdxB+Ep8Um19SHzNWDc9lztSDnxBEIDwyLgq8ZVyeGwM3/sUM5coqes4vR91FllyWowVlZ1SHd0NpOBpHKBae2lOkYnjtUwKBgQDmyMNebj2szF9dDJuhEn0nrfMt4np/sWS3kYdncxATMV3eeBjVnu7LI5TXCQWVvJOn3ireoXoR/Hu2rXtic68JZcK9oGWLG6ESN/pawxrPDOZmbaiBifPbBZF9iYRNos/rM59EDpx857YdmFrm3NiMmYySkPcagnklIy2wSYgWSwKBgQCUCzXx+hXOkR52a/uRafeJXGFVQzR1W8+GjSHBB0H+PpLVf2EDh3qXq1GurdErHFqy5iWTKOxy/5PjunuABn5cAfLi1YkKgUNrKlpjQO6m7WdQ0brVnv3cA2bPSmRkK9nPwJIcLyoc+6Ts5RjK4xflz+fcuBDuTwgHf5/XgJGxxQKBgEaGFUW3FHvInF+36ZsIJlzIyGY+Lkb1M5zxCNqxC7mBvXitZJFrFvfWsLYhmKZQUSLcPmd4GvCCImAJlYnZZiTOBvuhl0YyND94AbZYpsmc3I9ydSrYuD827WXabya2WSk7xjCq8mEtAbb7htPDqx0UEyP6zrg4JMMXg++6GK8PAoGBAMmKfqvZEzuO14rLGeMoMJzt4S440CnvcbxX8KTHB7/vvNjWYBoc8Lu96zFv/G9nZ8l/YM8vm03anxqMAvL5iaGV6sUs251Q5dDncCYJ4oLTMatqONlHvdRRTtuZwV6fOUigiPRardj8AnQSzcS1Qr9kuTwDg7yJcrBx20KTupob`

	publicPKCS1  = `MEgCQQCvt6NgpvX1cTlT3qnhmlq8Sz0xOSYXlzFr3C6Ub8ezyGHSRSGUB3URrxiBIGYTYwAS6eY79YH72+wFPEexRTWDAgMBAAE=`
	privatePKCS1 = `MIIBOgIBAAJBAK+3o2Cm9fVxOVPeqeGaWrxLPTE5JheXMWvcLpRvx7PIYdJFIZQHdRGvGIEgZhNjABLp5jv1gfvb7AU8R7FFNYMCAwEAAQJAD2eyH+i2hFLtpVGBsBDjbSgx45bdcapq35Uchrby2YTNoR8EOCBTL8f3eJUP78boGBVBo3O70rQTQEuNE4rQwQIhAOqWX4Df7jFhzhgYT2lI6a0qJERzeWJZrB5Orp85HgJfAiEAv8GkhDnZUBYuYkPmfmSnKyzB62S5v/Yri8oj1tkKR10CIDpfuNwk3OzE3bf5NKmKTzub5PJzmZzzm3TfN2y/lcwZAiEAvLUlYRbI/J6HkR6/Q11sgfmu8SeUAQySeCQ6tuGq0uUCIGXZE1N82/JpcBGdq9Mg2M4XxwLs/vDrKNmoOEgaiklP
`
)

func TestAES(t *testing.T) {
	str := "hello aes string"
	encrypt := AESEncrypt([]byte(str), []byte(aesKey))
	decrypt := AESDecrypt(encrypt, []byte(aesKey))
	fmt.Printf("原始：%s 解密后：%s  是否相等：%v \n", str, string(decrypt), str == string(decrypt))
}

func TestGenRSAKeys(t *testing.T) {
	private, public, err := GenRSAKeys(2048)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(private))
	fmt.Println(string(public))
}

func TestRSAByPem(t *testing.T) {
	str := "hello rsa string"
	encrypt := RSAEncrypt([]byte(publicPem), []byte(str))
	decrypt := RSADecrypt([]byte(privatePem), encrypt)
	fmt.Printf("原始：%s 解密后：%s  是否相等：%v \n", str, string(decrypt), str == string(decrypt))

	sign := RSASignSHA256([]byte(privatePem), []byte(str))
	fmt.Println("验签结果：", RSAVerifySHA256([]byte(publicPem), []byte(str), sign))
}

func TestRSAByString(t *testing.T) {
	str := "hello rsa string"
	encrypt := RSAEncryptString([]byte(public), []byte(str))
	decrypt := RSADecryptString([]byte(private), encrypt)
	fmt.Printf("原始：%s 解密后：%s  是否相等：%v \n", str, string(decrypt), str == string(decrypt))

	sign := RSASignSHA256String([]byte(private), []byte(str))
	fmt.Println("验签结果：", RSAVerifySHA256String([]byte(public), []byte(str), sign))
}

func TestRSAByPKCS1Pem(t *testing.T) {
	str := "hello rsa string"
	encrypt := RSAEncryptPKCS1([]byte(publicPem), []byte(str))
	decrypt := RSADecryptPKCS1([]byte(privatePem), encrypt)
	fmt.Printf("原始：%s 解密后：%s  是否相等：%v \n", str, string(decrypt), str == string(decrypt))

	sign := RSASignSHA256([]byte(privatePem), []byte(str))
	fmt.Println("验签结果：", RSAVerifyPKCS1SHA256([]byte(publicPem), []byte(str), sign))
}

func TestRSAByPKCS1String(t *testing.T) {
	str := "hello rsa string"
	encrypt := RSAEncryptPKCS1String([]byte(publicPKCS1), []byte(str))
	decrypt := RSADecryptPKCS1String([]byte(privatePKCS1), encrypt)
	fmt.Printf("原始：%s 解密后：%s  是否相等：%v \n", str, string(decrypt), str == string(decrypt))

	sign := RSASignPKCS1SHA256String([]byte(privatePKCS1), []byte(str))
	fmt.Println("验签结果：", RSAVerifyPKCS1SHA256String([]byte(publicPKCS1), []byte(str), sign))
}

func TestRSAByAuto(t *testing.T) {
	str := "hello rsa string"
	encrypt := RSAEncrypt([]byte(public), []byte(str))
	decrypt := RSADecrypt([]byte(private), encrypt)
	fmt.Printf("传入String-原始：%s 解密后：%s  是否相等：%v \n", str, string(decrypt), str == string(decrypt))

	encrypt = RSAEncrypt([]byte(publicPem), []byte(str))
	decrypt = RSADecrypt([]byte(privatePem), encrypt)
	fmt.Printf("传入Pem-原始：%s 解密后：%s  是否相等：%v \n", str, string(decrypt), str == string(decrypt))

	sign := RSASignSHA256([]byte(private), []byte(str))
	fmt.Println("传入Pem-验签结果：", RSAVerifySHA256([]byte(publicPem), []byte(str), sign))

	sign = RSASignSHA256([]byte(private), []byte(str))
	fmt.Println("传入String-验签结果：", RSAVerifySHA256([]byte(public), []byte(str), sign))
}
