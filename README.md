# Go Utils README
基于日常学习及项目需要所积累的go语言常用工具库

## byte
- ### int64 to byte
  ```go
  func Int64ToBytes(i int64) []byte
  ```
- ### byte to int64
  ```go
  func BytesToInt64(buf []byte) int64
  ```
## dingtalk
- ### 钉钉报警
  ```go
  func DingTalkAlarm(serverName, message string) bool
  ```
  ```go
  func DingTalkAlarmUrl(url, serverName, message string) bool
  ```
## encryption
- ### AES加密
  ```go
  func AESEncrypt(msg, key []byte) []byte
  
  func AESEncryptE(msg, key []byte) ([]byte, error)
  
  func AESEncryptIv(msg, key, iv []byte) ([]byte, error)
  ```
- ### AES解密
  ```go
  func AESDecrypt(msg, key []byte) []byte
  
  func AESDecryptE(msg, key []byte) ([]byte, error)
  
  func AESDecryptIv(msg, key, iv []byte) ([]byte, error)
  ```
- ### RSA加密
  ```go
  func RSAEncrypt(public []byte, msg []byte) []byte
  
  func RSAEncryptE(public []byte, msg []byte) ([]byte, error)
  
  func RSAEncryptPKCS1(public []byte, msg []byte) []byte
  
  func RSAEncryptPKCS1E(public []byte, msg []byte) ([]byte, error)
  
  func RSAEncryptPem(public []byte, msg []byte) []byte
  
  func RSAEncryptPemE(public []byte, msg []byte) ([]byte, error)
  
  func RSAEncryptPKCS1Pem(public []byte, msg []byte) []byte
  
  func RSAEncryptPKCS1PemE(public []byte, msg []byte) ([]byte, error)
  
  func RSAEncryptString(public, msg []byte) []byte
  
  func RSAEncryptStringE(public, msg []byte) ([]byte, error)
  
  func RSAEncryptPKCS1String(public, msg []byte) []byte
  
  func RSAEncryptPKCS1StringE(public, msg []byte) ([]byte, error)
  ```
- ### RSA解密
  ```go
  // RSADecrypt 使用私钥进行解密
  func RSADecrypt(private []byte, cipherText []byte) []byte
  
  // RSADecryptPKCS1 使用私钥进行解密
  func RSADecryptPKCS1(private []byte, cipherText []byte) []byte
  
  // RSADecryptE 使用私钥进行解密
  func RSADecryptE(private []byte, cipherText []byte) ([]byte, error)
  
  // RSADecryptPKCS1E 使用私钥进行解密
  func RSADecryptPKCS1E(private []byte, cipherText []byte) ([]byte, error)
  
  // RSADecryptPem 使用私钥（pem格式）进行解密
  func RSADecryptPem(private []byte, cipherText []byte) []byte
  
  // RSADecryptPKCS1Pem 使用私钥（pem格式）进行解密
  func RSADecryptPKCS1Pem(private []byte, cipherText []byte) []byte
  
  // RSADecryptString 使用私钥（String格式）进行解密
  func RSADecryptString(private, cipherText []byte) []byte
  
  // RSADecryptPKCS1String 使用私钥（String格式）进行解密
  func RSADecryptPKCS1String(private, cipherText []byte) []byte
  
  // RSADecryptPemE 使用私钥（pem格式）进行解密
  func RSADecryptPemE(private []byte, cipherText []byte) ([]byte, error)
  
  // RSADecryptStringE 使用私钥（String格式）进行解密
  ```
## hash
- ### hash算法
  ```go
  func Hash64Byte(data []byte) uint64
  
  func Hash64(data string) uint64
  
  func Hash32Byte(data []byte) uint32
  
  func Hash32(data string) uint32
  ```

## html
- ### 删除html标签
  ```go
  func TrimHtml(src string) string
  ```

## http
- ### http请求
  ```go
  func HttpDO(method string, url string, body io.Reader,
	header map[string][]string) (httpStatus int, resp []byte, err error)
  
  func HttpDOTimeOut(method string, url string, body io.Reader,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error)
  ```
- ### http POST 请求
  ```go
  func HttpPost(url string, body io.Reader,
	header map[string][]string) (httpStatus int, resp []byte, err error)
  ```
- ### http GET 请求
  ```go
  func HttpGet(url string, body io.Reader,
      header map[string][]string) (httpStatus int, resp []byte, err error)
  ```
- ### http Post Json 请求
  ```go
  func HttpPostJson(url string, body []byte,
      header map[string][]string) (httpStatus int, resp []byte, err error)
  
  func HttpPostTimeOut(url string, body io.Reader,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error)
  
  func HttpPostJsonTimeOut(url string, body []byte,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error)
  ```
- ### http GET Json 请求
  ```go
  func HttpGetJson(url string, body []byte,
      header map[string][]string) (httpStatus int, resp []byte, err error)
  
  func HttpGetTimeOut(url string, body io.Reader,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error)
  
  func HttpGetJsonTimeOut(url string, body []byte,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error)
  ```
- ### http Form 请求
  ```go
  func HttpPostForm(postUrl string, body map[string][]string,
      header map[string][]string) (httpStatus int, resp []byte, err error)
  
  func HttpPostFormTimeOut(postUrl string, body map[string][]string,
	header map[string][]string, millisecond int) (httpStatus int, resp []byte, err error)
  ```
## interface
- ### 判断接口具体类型
  ```go
  func JudgeType(v interface{}) string
  ```
- ### 使用反射判断接口具体类型
  ```go
  func JudgeTypeByReflect(v interface{}) string
  ```

## json
- ###
  ```go
  func JSON2Object(data []byte, obj interface{})
  
  func JSON2ObjectE(data []byte, obj interface{})
  
  func JSON2ObjectUseNumberE(data []byte, obj interface{}) (err error)
  
  func Object2JSON(obj interface{}) string
  
  func Object2JSONByte(obj interface{}) []byte
  
  func Object2JSONByteE(obj interface{}) ([]byte, error)
  
  func Object2JSONE(obj interface{}) (string, error)
  
  func JSON2Map(json []byte) map[string]interface{}
  
  func JSON2MapUseNumber(json []byte) map[string]interface{}
  
  // Valid 验证JSON字符串是否合法。此方法只验证标准格式的，开头和结尾为{}
  // jsoniter.Valid方法“abc”也可以验证通过
  func Valid(json []byte) bool
  
  func GzipEncode(body []byte) (result []byte)
  
  func GzipDecode(body []byte) (result []byte)
  
  func HuffmanEncode(body []byte) (result []byte)
  
  func HuffmanDecode(body []byte) (result []byte)
  ```

## map
- ### 获取map所有key的方式
  ```go
  // 方法1（效率很高）：
  func getMapKeys1(m map[int]int) []int
  ```
  ```go
  // 方法2（效率很高）：
  func getMapKeys2(m map[int]int) []int
  ```
  ```go
  // 方法3（效率较低）：
  func getMapKeys3(m map[int]int) []int
  ```
  ```go
  // 方法4（效率极低）：
  func getMapKeys4(m map[int]int) int
  ```

- ### 使用map过滤数组
  ```go
  func FilterSliceByMap(filter map[string]struct{}, pram []string, isExist bool) (array []string)
  ```
- ### 使用map过滤map
  ```go
  func FilterMapByMap(filter map[string]struct{}, pram map[string]string, isExist bool) (array map[string]string)
  ```
- ### 合并map
  ```go
  func MergeMap(m1, m2 map[string]interface{}) map[string]interface{}
  ```
- ### 数组去重
  ```go
  func UniqueArray(m []string) []string 
  ```
- ### 根据 "," 分割map为map数组
  ```go
  func MapSplitByComma(fields map[string]string) []map[string]string
  ```

## net
- ### 验证IPv4的合法性
  ```go
  func IsIP4(ip string) bool
  ```

- ### 验证IPv6的合法性
  ```go
  func IsIP6(ip string) bool
  ```

- ### 验证ID是否合法性
  ```go
  // 0: invalid ip
  // 4: IPv4
  // 6: IPv6
  func ParseIP(s string) int
  ```

- ### 判断IP是否为内网IP
  ```go
  func IsPrivateIP(ipStr string) bool
  ```

- ### 获取内网IP
  ```go
  func InternalIp() string
  ```

## os
- ### 使用map替换template模版中 $ 符号后的字符串
  ```go
  func ExpandByMap(template string, fields map[string]string) string
  ```
- ### 自定义函数规则替换字符串模版中 $ 符号后的字符串
  ```go
  func ExpandByFun(str string, f func(string) string) string
  ```
- ### 使用map替换template模版中 $ 符号后的字符串，返回字符串数组
  ```go
  /**
  *	输入示例：
  *	"command_fields": {
  *      "user_id": "1,2,3",
  *      "prop": "huChao"
  *   }
  *
  *  template: "chartid=${user_id}&prop=${prop}"
  *
  *  输出示例：
  *  [
  *       "chartid=1&prop=hudaoju",
  *       "chartid=2&prop=hudaoju",
  *       "chartid=3&prop=hudaoju"
  *   ]
  *
  */
  // isMultiple: 是否根据 "," 分割
  func GetComposedTemplateListExpandByMap(template string, isMultiple bool, fields map[string]string) []string
  ```

## slice

- ### 切片去重
  ```go
  func removeRepByMap(slc []*SeriesRes) []*SeriesRes
  ```
- ### 结构体切片去重
    ``` go
  func removeStructRepByMap(slc []*SeriesRes) []*SeriesRes
    ```
- ### 切片转链表
    ``` go
  func SliceToLinkList(nums []int, head *ListNode) *ListNode
    ```
## strings
- ### 去除空格
  ```go
  func TrimSpace(str string) string
  ```

- ### 返回可比较的字符串，当传入非法ASCII码时返回空字符串
  ```go
  func VersionOrdinal(version string) string
  ```
- ### 比较两个版本号（版本号只有数字和点组成）
  ```go
  func VersionGreater(versionA, versionB string) int
  ```

- ### 通过Builder拼接字符串
  ```go
  func StringJoin(strs ...string) string
  ```

- ### 精准的字符串匹配，区分大小写
  ```go
  func IsExactExist(array []string, row string) bool
  ```

- ### 将字符串切片转化成int32切片
  ```go
  func Str2Int32Array(strArray []string) ([]int32, error)
  ```

## struct
- ### 拷贝结构体
    ``` go
    func StructCopy(src, dst interface{}) (err error)
    ```
## time
- ### 获取当前时间
  ```go
  func NowTime() string
  ```
- ### 获取当前时间戳
  ```go
  func NowUnix() int64
  ```
- ### 获取当前时间
  ```go
  func UnixToFormatTime(timeStamp int64) string
  ```
- ### 根据开始日期和结束日期计算出时间段内所有日期
  ```go
  func GetBetweenMonths(startTime, endTime time.Time) ([]string, error)   // 月
  func GetBetweenDates(startTime, endTime time.Time) ([]string, error)    // 日
  ```
- ### FormatTimeIfNotZero
  ```go
  func FormatTimeIfNotZero(time time.Time, layout string) string
  ```
## time - 时区
- ### 北京时区
  ```go
  func DateCST(layout string) string
  ```
- ### 美国标准时区
  ```go
  func DatePST(layout string) string
  ```
- ### 日本时区
  ```go
  func DateJST(layout string) string
  ```
- ### 韩国时区
  ```go
  func DateKST(layout string) string
  ```
- ### UTC时间
  ```go
  func DateUTC(layout string) string
  ```  

## uuid
- ### 生成36位UUID
  ```go
  func UUID() string
  ```
- ### 生成32位UUID
  ```go
  func UUID32() string
  ```
- ### 通过内容生成MD5
  ```go
  func MD5(body []byte) string
  ```
- ### 通过传入的参数生成签名
  ```go
  func PramSign(pram []string) string 
  ```

## version
- ### 版本比较
  ```go 
  //VersionOrdinal 返回可比较的字符串，当传入非法ASCII码时返回空字符串
  //用于版本比较
  func VersionOrdinal(version string) string
  ```
- ### 比较两个版本号。版本号只有数字和点组成
  ```go
  //VersionGreater 比较两个版本号。版本号只有数字和点组成
  // 如：versionA == versionB  返回 0
  // 如：versionA > versionB  返回 1
  // 如：versionA < versionB  返回 -1
  func VersionGreater(versionA, versionB string) int
  ```
- ### 验证版本
  ```go
  //VersionCheck 验证版本，该方法只支持app市场版本的格式：主版本.此版本.修订版本.热更版本。（且可以使用0开头）
  //理论最大版本号：999.999.999.999
  func VersionCheck(v string) bool
  ```
- ### 通过传入的版本号获取app的市场版本
  ```go
  //VersionApp 通过传入的版本号获取app的市场版本。
  //app的市场版本格式为：x.x.x
  func VersionApp(v string) string
  ```
- ### 通过传入的版本号和count确定返回几位的版本号
  ```go
  //VersionAppByCount 通过传入的版本号和count确定返回几位的版本号
  //v：版本号
  //count：需要返回的版本号位数
  func VersionAppByCount(v string, count int) string
  ```
  
持续更新,欢迎pr👏
- 微信公众号：Gopher毛
- 交流QQ群3: 805360166(活跃) 
- 交流QQ群2: 579480724
- 交流QQ群1：1007576722
