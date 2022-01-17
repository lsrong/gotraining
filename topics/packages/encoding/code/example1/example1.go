package main

import (
	"encoding/json"
	"fmt"
)

// Sample program to show how to unmarshal a JSON document into
// a user defined struct type.
// 演示如何将 JSON 文档转义为用户定义的结构类型的示例程序。

// document 演示的json字符串
var document = `{
"credentials": {
    "token": "06142010_1:75bf6a413327dd71ebe8f3f30c5a4210a9b11e93c028d6e11abfca7ff"
},
"valid": true,
"locale": "en_US",
"tnc_version": 2,
"preference_info": {
    "currency_code": "USD",
    "time_zone": "PST",
    "number_format": {
        "decimal_separator": ".",
        "grouping_separator": ",",
        "group_pattern": "###,##0.##"
    }
 }
}`

// PreferenceInfo UserContext 需要从document解析结构体字段信息
type (
	PreferenceInfo struct {
		CurrencyCode string `json:"currency_code"`
		TimeZone     string `json:"time_zone"`
		NumberFormat struct {
			DecimalSeparator  string `json:"decimal_separator"`
			GroupingSeparator string `json:"grouping_separator"`
			GroupPattern      string `json:"group_pattern"`
		} `json:"number_format"`
	}
	UserContext struct {
		Credentials struct {
			Token string `json:"token"`
		} `json:"credentials"`
		Valid          bool           `json:"valid"`
		Locate         string         `json:"locate"`
		TncVersion     int            `json:"tnc_version"`
		PreferenceInfo PreferenceInfo `json:"preference_info"`
	}
)

func main() {
	// 解码到结构体变量中
	var uc UserContext
	if err := json.Unmarshal([]byte(document), &uc); err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("userContext: %+v \n", uc)

	// 解码到一个nil指针, 取值: *pointer
	var ucp *UserContext
	if err := json.Unmarshal([]byte(document), &ucp); err != nil {
		fmt.Println("error: ", err)
		return
	}
	fmt.Printf("*userContext: %+v", *ucp)
}
