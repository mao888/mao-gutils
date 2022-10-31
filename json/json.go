package gutil

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io"

	jsoniter "github.com/json-iterator/go"
)

const EmptyString = ""

var useNumber = jsoniter.Config{
	EscapeHTML:             true,
	UseNumber:              true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
}.Froze()

var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary

func JSON2Object(data []byte, obj interface{}) {
	_ = jsonIterator.Unmarshal(data, obj)
}

func JSON2ObjectE(data []byte, obj interface{}) (err error) {
	return jsonIterator.Unmarshal(data, obj)
}

func JSON2ObjectUseNumberE(data []byte, obj interface{}) (err error) {
	return useNumber.Unmarshal(data, obj)

}

func Object2JSON(obj interface{}) string {
	resp, err := Object2JSONE(obj)
	if err != nil {
		return EmptyString
	}
	return resp
}

func Object2JSONByte(obj interface{}) []byte {
	resp, _ := Object2JSONByteE(obj)
	return resp
}

func Object2JSONByteE(obj interface{}) ([]byte, error) {
	return jsonIterator.Marshal(obj)
}

func Object2JSONE(obj interface{}) (string, error) {
	resp, err := Object2JSONByteE(obj)
	if err != nil {
		return EmptyString, err
	}
	return string(resp), nil
}

func JSON2Map(json []byte) map[string]interface{} {
	row := make(map[string]interface{})
	if err := jsonIterator.Unmarshal(json, &row); err != nil {
		return nil
	}
	return row
}

func JSON2MapUseNumber(json []byte) map[string]interface{} {
	row := make(map[string]interface{})
	if err := useNumber.Unmarshal(json, &row); err != nil {
		return nil
	}
	return row
}

//Valid 验证JSON字符串是否合法。此方法只验证标准格式的，开头和结尾为{}。
//jsoniter.Valid方法“abc”也可以验证通过
func Valid(json []byte) bool {
	if len(json) == 0 {
		return false
	}
	//是否已{开头或者}结尾
	if json[0] != 123 && json[len(json)-1] != 125 {
		return false
	}
	api := jsoniter.ConfigCompatibleWithStandardLibrary
	row := make(map[string]interface{})
	if err := api.Unmarshal(json, &row); err != nil {
		return false
	}
	return true
}

func GzipEncode(body []byte) (result []byte) {
	if len(body) == 0 {
		return
	}
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	defer writer.Close() //nolint
	if _, err := writer.Write(body); err != nil {
		return
	}
	if err := writer.Flush(); err != nil {
		return
	}
	return buf.Bytes()
}

func GzipDecode(body []byte) (result []byte) {
	if len(body) == 0 {
		return
	}
	reader, err := gzip.NewReader(bytes.NewReader(body))
	if err != nil {
		return nil
	}
	defer reader.Close() //nolint
	result, _ = io.ReadAll(reader)
	return result
}

func HuffmanEncode(body []byte) (result []byte) {
	if len(body) == 0 {
		return
	}
	var buf bytes.Buffer
	writer, err := flate.NewWriter(&buf, flate.HuffmanOnly)
	if err != nil {
		return
	}
	defer writer.Close() //nolint
	if _, err = writer.Write(body); err != nil {
		return
	}
	writer.Flush() //nolint
	return buf.Bytes()
}

func HuffmanDecode(body []byte) (result []byte) {
	if len(body) == 0 {
		return
	}
	reader := flate.NewReader(bytes.NewReader(body))
	defer reader.Close() //nolint
	result, _ = io.ReadAll(reader)
	return
}
