package uuid

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strings"

	"github.com/google/uuid"
)

// UUID 生成UUID
func UUID() string {
	return uuid.New().String()
}

func UUID32() string {
	return strings.ReplaceAll(UUID(), "-", "")
}

// MD5 通过内容生成MD5
func MD5(body []byte) string {
	h := md5.New()
	h.Write(body)
	return hex.EncodeToString(h.Sum(nil))
}

//PramSign 通过传入的参数生成签名
func PramSign(pram []string) string {
	sort.Strings(pram)
	return MD5([]byte(strings.Join(pram, "")))
}
