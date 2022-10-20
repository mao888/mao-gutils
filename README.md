# Go Utils README
基于日常学习及项目需要所积累的go语言常用工具库

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