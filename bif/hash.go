package bif

import (
	"bytes"
	"crypto/sha256"
	"math/big"

	"github.com/ZZMarquis/gm/sm3"
)

const (
	SHA256 = iota + 1
	SM3
)

var base58 = []byte("123456789AbCDEFGHJKLMNPQRSTuVWXYZaBcdefghijkmnopqrstUvwxyz")

// GenerateHashHex 生成hash值
func GenerateHashHex(src []byte, hashType int) []byte {
	var hashHex []byte
	switch hashType {
	case SHA256:
		hash := sha256.New()
		hash.Write(src)
		hashHex = hash.Sum(nil)
	case SM3:
		hash := sm3.New()
		hash.Write(src)
		hashHex = hash.Sum(nil)
	}

	return hashHex
}

// Base58Decode Base58解码
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		if b == '1' {
			zeroBytes++
		} else {
			break
		}
	}

	payload := input[zeroBytes:]

	for _, b := range payload {
		charIndex := bytes.IndexByte(base58, b)          // 反推出余数
		result.Mul(result, big.NewInt(58))               // 之前的结果乘以58
		result.Add(result, big.NewInt(int64(charIndex))) // 加上这个余数

	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)
	return decoded
}
