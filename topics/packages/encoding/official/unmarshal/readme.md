# func Unmarshal
```
func Unmarshal(data []byte, v interface{}) error
```
Unmarshal 解析 JSON 编码的数据并将结果存储在 v 指向的值中。如果 v 为 nil 或不是指针，则 Unmarshal 返回 InvalidUnmarshalError。

Unmarshal 使用 Marshal 使用的编码的逆编码，根据需要分配映射、切片和指针，并具有以下附加规则：

要将 JSON 解组为指针，Unmarshal 首先处理 JSON 为 JSON 文字 null 的情况。在这种情况下，Unmarshal 将指针设置为 nil。否则，Unmarshal 将 JSON 解组为指针指向的值。如果指针为 nil，Unmarshal 为其分配一个新值来指向。

要将 JSON 解组为实现 Unmarshaler 接口的值，Unmarshal 会调用该值的 UnmarshalJSON 方法，包括当输入为 JSON null 时。否则，如果该值实现 encoding.TextUnmarshaler 并且输入是 JSON 引用的字符串，则 Unmarshal 使用该字符串的未引用形式调用该值的 UnmarshalText 方法。

要将 JSON 解组到结构中，Unmarshal 将传入的对象键与 Marshal 使用的键（结构字段名称或其标记）匹配，首选完全匹配但也接受不区分大小写的匹配。默认情况下，不具有相应结构字段的对象键将被忽略（请参阅 Decoder.DisallowUnknownFields 了解替代方案）。

为了将 JSON 解组为接口值，Unmarshal 将其中一项存储在接口值中：

bool，用于 JSON 布尔值
float64，用于 JSON 数字
字符串，用于 JSON 字符串
[]interface{}，用于 JSON 数组
map[string]interface{}，用于 JSON 对象
无 JSON null
要将 JSON 数组解组为切片，Unmarshal 会将切片长度重置为零，然后将每个元素附加到切片。作为一种特殊情况，为了将空 JSON 数组解组为切片，Unmarshal 将切片替换为新的空切片。

要将 JSON 数组解组为 Go 数组，Unmarshal 将 JSON 数组元素解码为相应的 Go 数组元素。如果 Go 数组小于 JSON 数组，则丢弃额外的 JSON 数组元素。如果 JSON 数组小于 Go 数组，则额外的 Go 数组元素设置为零值。

要将 JSON 对象解组到映射中，Unmarshal 首先建立要使用的映射。如果映射为 nil，Unmarshal 分配一个新映射。否则 Unmarshal 会重用现有映射，保留现有条目。然后 Unmarshal 将 JSON 对象中的键值对存储到映射中。映射的键类型必须是任何字符串类型、整数、实现 json.Unmarshaler 或实现 encoding.TextUnmarshaler。

如果 JSON 值不适合给定的目标类型，或者 JSON 编号溢出目标类型，Unmarshal 会跳过该字段并尽其所能完成解组。如果没有遇到更严重的错误，Unmarshal 返回一个 UnmarshalTypeError 描述最早的此类错误。在任何情况下，都不能保证有问题的字段后面的所有剩余字段都将被解组到目标对象中。

JSON null 值通过将 Go 值设置为 nil 来解组为接口、映射、指针或切片。因为 null 在 JSON 中经常用来表示“不存在”，所以将 JSON null 解组到任何其他 Go 类型对值没有影响并且不会产生错误。

解组带引号的字符串时，无效的 UTF-8 或无效的 UTF-16 代理对不会被视为错误。相反，它们被 Unicode 替换字符 U+FFFD 替换。