## Encoding - Standard Library

Encoding is the process or marshaling or unmarshaling data into different forms. Taking JSON string documents and convert them to values of our user defined types is a very common practice in many go programs today. Go's support for encoding is amazing and improves and gets faster with every release.  
编码是将数据编组或解组为不同形式的过程。获取 JSON 字符串文档并将它们转换为我们用户定义类型的值是当今许多 Go 程序中非常常见的做法。 Go 对编码的支持是惊人的，并且随着每个版本的发布而改进和变得更快。
## Notes

* Support for Decoding and Encoding JSON and XML are provided by the standard library.
* 标准库提供对 JSON 和 XML 解码和编码的支持。
* 
* This package gets better and better with every release.
* 这个包随着每个版本的发布而变得越来越好。

## Links

[Package Encoding JSON](https://pkg.go.dev/encoding/json@go1.17.6)  
[Decode JSON Documents In Go](https://www.ardanlabs.com/blog/2014/01/decode-json-documents-in-go.html) - William Kennedy  
[JSON and Go](https://go.dev/blog/json) - Andrew Gerrand

## Code Review

[Unmarshal JSON documents](code/example1/example1.go)    
[Unmarshal JSON files](code/example2/example2.go)    
[Marshal a user defined type](code/example3/example3.go)    
[Custom Marshaler and Unmarshler](code/example4/example4.go)  


## Official Code 
[Custom Marshal JSON](official/custom_marshal_json/main.go)  
[Text Marshal JSON](official/text_marshal_json/main.go)  
[Marshal JSON](official/marshal/main.go)  
[Unmarshal JSON](official/unmarshal/main.go)  

## Exercises

### Exercise 1

**Part A** Create a file with an array of JSON documents that contain a user name and email address. Declare a struct type that maps to the JSON document. Using the json package, read the file and create a slice of this struct type. Display the slice.

**Part B** Marshal the slice into pretty print strings and display each element.

[Template](exercises/template1/template1.go)  |
[Answer](exercises/exercise1/exercise1.go)
