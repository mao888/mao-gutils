package bif

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"

	"github.com/ZZMarquis/gm/sm2"
	"github.com/caict-4iot-dev/BIF-Core-SDK-Go/module/encryption/key"
)

// 加密类型
const (
	ED25519 = iota + 1
	SM2

	ED25519Value = 'e'
	SM2Value     = 'z'

	BASE58Value = 'f'
)

// GetRawPrivateKey 星火私钥转原生私钥
func GetRawPrivateKey(encPrivateKey []byte) (int, []byte, error) {
	priKeyTmp := Base58Decode(encPrivateKey)
	if len(priKeyTmp) <= 5 {
		return 0, nil, errors.New("private key (" + string(encPrivateKey) + ") is invalid 1")
	}

	if priKeyTmp[3] != ED25519Value && priKeyTmp[3] != SM2Value {
		return 0, nil, errors.New("private key (" + string(encPrivateKey) + ") is invalid 2")
	}
	var keyType int
	switch priKeyTmp[3] {
	case ED25519Value:
		{
			keyType = ED25519
		}
	case SM2Value:
		{
			keyType = SM2
		}
	default:
		return 0, nil, errors.New("Private key (" + string(encPrivateKey) + ") is invalid")
	}
	if priKeyTmp[4] != BASE58Value {
		return 0, nil, errors.New("private key (" + string(encPrivateKey) + ") is invalid 3")
	}

	var rawPrivateKey []byte
	switch keyType {
	case ED25519:
		rawPrivateKey = ed25519.NewKeyFromSeed(priKeyTmp[5:])
	case SM2:
		rawPrivateKey = priKeyTmp[5:]
	}

	return keyType, rawPrivateKey, nil
}

// EncPublicKey 原生公钥转星火公钥
func EncPublicKey(rawPublicKey []byte, keyType int) string {
	buff := make([]byte, len(rawPublicKey)+3)
	buff[0] = 0xB0
	switch keyType {
	case ED25519:
		buff[1] = ED25519Value
	case SM2:
		buff[1] = SM2Value
	default:
		return ""
	}

	buff[2] = BASE58Value
	buff = append(buff[:3], rawPublicKey...)

	return hex.EncodeToString(buff)
}

// GetEncPublicKeyByEncPrivateKey 星火私钥获取星火公钥
func GetEncPublicKeyByEncPrivateKey(encPrivateKey []byte) (string, error) {
	keyType, rawPrivateKey, err := GetRawPrivateKey(encPrivateKey)
	if err != nil {
		return "", err
	}
	var rawPublicKey []byte

	switch keyType {
	case ED25519:
		rawPublicKey = rawPrivateKey[32:]
	case SM2:
		priKey, err := sm2.RawBytesToPrivateKey(rawPrivateKey)
		if err != nil {
			return "", err
		}
		pubKey := sm2.CalculatePubKey(priKey)
		rawPublicKey, err = hex.DecodeString("04" + hex.EncodeToString(pubKey.GetRawBytes()))
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("type does not exist")
	}

	return EncPublicKey(rawPublicKey, keyType), nil
}

// Sign 签名
func Sign(encPrivate []byte, message []byte) ([]byte, error) {
	keyType, rawPrivateKey, err := GetRawPrivateKey(encPrivate)
	if err != nil {
		return nil, err
	}

	var sign []byte
	switch keyType {
	case ED25519:
		sign25519 := ed25519.Sign(rawPrivateKey, message)
		sign = sign25519
	case SM2:
		priKey, err := sm2.RawBytesToPrivateKey(rawPrivateKey)
		if err != nil {
			return nil, err
		}
		r, s, err := sm2.SignToRS(priKey, []byte("1234567812345678"), message)
		if err != nil {
			return nil, err
		}
		rBytes := r.Bytes()
		sBytes := s.Bytes()
		sig := make([]byte, 64)
		if len(rBytes) == 33 {
			copy(sig[:32], rBytes[1:])
		} else {
			copy(sig[:32], rBytes[:])
		}
		if len(sBytes) == 33 {
			copy(sig[32:], sBytes[1:])
		} else {
			copy(sig[32:], sBytes[:])
		}

		sign = sig
	}

	return sign, nil
}

func Verify(signData, publicKey, msg string) (bool, error) {
	signDataStr, err := hex.DecodeString(signData)
	if err != nil {
		return false, err
	}
	msgB, err := hex.DecodeString(msg)
	if err != nil {
		return false, err
	}
	// 先验证 ED25519 算法
	bPublicKey := []byte(publicKey)
	var keyType int
	rawPublicKey := key.GetRawPublicKey(bPublicKey)
	if l := len(rawPublicKey); l == ed25519.PublicKeySize {
		keyType = key.ED25519
	} else {
		keyType = key.SM2
	}
	passed := key.Verify(
		bPublicKey,
		msgB,
		signDataStr,
		keyType)
	return passed, nil
}
