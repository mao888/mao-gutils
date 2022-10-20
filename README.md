# Go Utils README
基于日常学习及项目需要所积累的go语言常用工具库

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

## slice

- ### 切片去重
  ```go
  func removeRepByMap(slc []*SeriesRes) []*SeriesRes
  ```
- ### 结构体切片去重
    ``` go
  func removeStructRepByMap(slc []*SeriesRes) []*SeriesRes
    ```



## struct
- ### 拷贝结构体
    ``` go
    func StructCopy(src, dst interface{}) (err error)
    ```


持续更新