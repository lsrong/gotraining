// Sample program to show how write a custom Unmarshal and Marshal functions.
// 演示如何编写自定义 Unmarshal 和 Marshal 函数的示例程序。
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// NullableTime 定义一个可以支持null的时间对象
type NullableTime struct {
	time.Time
}

// MarshalJSON 定义可以编码null的时间
func (t NullableTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}

	b := make([]byte, 0, len(time.RFC3339)*2)
	b = append(b, '"')
	b = t.AppendFormat(b, time.RFC3339)
	b = append(b, '"')

	return b, nil
}

// UnmarshalJSON 将null解码成time.Time{}
func (t NullableTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		t.Time = time.Time{}
		return nil
	}
	tt, err := time.Parse(time.RFC3339, string(b))
	if err != nil {
		return err
	}

	t.Time = tt
	return nil
}

func main() {
	now := NullableTime{time.Now()}
	b, err := json.Marshal(&now)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Nullable: ", string(b))

	var n NullableTime
	if err := json.Unmarshal([]byte("null"), &n); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Nullable: ", n)

}
