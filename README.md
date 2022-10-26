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

## map
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

## slice

- ### 切片去重
  ```go
  func removeRepByMap(slc []*SeriesRes) []*SeriesRes
  ```
- ### 结构体切片去重
    ``` go
  func removeStructRepByMap(slc []*SeriesRes) []*SeriesRes
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


持续更新

交流qq群：1007576722