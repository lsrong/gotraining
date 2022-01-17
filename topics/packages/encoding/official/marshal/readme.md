# func Marshal

```
func Marshal(v interface{}) ([]byte, error)
```

Marshal traverses the value v recursively. If an encountered value implements the Marshaler interface and is
not a nil pointer, Marshal calls its MarshalJSON method to produce JSON. If no MarshalJSON method is present
but the value implements encoding.TextMarshaler instead, Marshal calls its MarshalText method
and encodes the result as a JSON string. The nil pointer exception is not strictly necessary but mimics a similar,
necessary exception in the behavior of UnmarshalJSON.

Marshal 递归地遍历值 v。如果遇到的值实现了 Marshaler 接口并且不是 nil 指针，则 Marshal 调用其 MarshalJSON 方法来生成 JSON。
如果不存在 MarshalJSON 方法但值实现 encoding.TextMarshaler ，则 Marshal 调用其 MarshalText 方法并将结果编码为 JSON 字符串。
nil 指针异常不是绝对必要的，但在 UnmarshalJSON 的行为中模仿了类似的必要异常。

Otherwise, Marshal uses the following type-dependent default encodings:  
否则，Marshal 使用以下依赖于类型的默认编码：

Boolean values encode as JSON booleans.  
布尔值编码为 JSON 布尔值。

Floating point, integer, and Number values encode as JSON numbers.  
浮点、整数和数字值编码为 JSON 数字。

String values encode as JSON strings coerced to valid UTF-8, replacing invalid bytes with the Unicode replacement rune. So that the JSON will be safe to embed inside HTML `<script>` tags, the string is encoded using HTMLEscape, which replaces "<", ">", "&", U+2028, and U+2029 are escaped to "\u003c","\u003e", "\u0026", "\u2028", and "\u2029". This replacement can be disabled when using an Encoder, by calling SetEscapeHTML(false).  
字符串值编码为强制转换为有效 UTF-8 的 JSON 字符串，用 Unicode 替换符文替换无效字节。
为了让 JSON 可以安全地嵌入 HTML `<script>` 标记，字符串使用 HTMLEscape 进行编码，它替换了 "<"、">"、"&"、U+2028 和 U+2029
转义为 "\ u003c”、“\u003e”、“\u0026”、“\u2028”和“\u2029”。使用编码器时，可以通过调用 SetEscapeHTML(false) 禁用此替换。

Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string,
and a nil slice encodes as the null JSON value.  
数组和切片值编码为 JSON 数组，除了 []byte 编码为 base64 编码字符串，nil 切片编码为空 JSON 值。

Struct values encode as JSON objects. Each exported struct field becomes a member of the object,
using the field name as the object key, unless the field is omitted for one of the reasons given below.  
结构值编码为 JSON 对象。每个导出的结构字段都成为对象的成员，使用字段名称作为对象键，除非由于以下原因之一省略了该字段。

The encoding of each struct field can be customized by the format string stored under the "json" key in the struct field's tag.  
每个结构字段的编码可以通过存储在结构字段标签中“json”键下的格式字符串进行自定义。

The format string gives the name of the field, possibly followed by a comma-separated list of options.  
格式字符串给出了字段的名称，可能后跟以逗号分隔的选项列表。

The name may be empty in order to specify options without overriding the default field name.  
名称可以为空，以便在不覆盖默认字段名称的情况下指定选项。

The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value,
defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.  
“omitempty”选项指定如果字段具有空值（定义为 false、0、nil 指针、nil 接口值以及任何空数组、切片、映射或字符串），则应从编码中省略该字段。

As a special case, if the field tag is "-", the field is always omitted. Note that a field with name "-"
can still be generated using the tag "-,".  
作为一种特殊情况，如果字段标记为“-”，则始终省略该字段。请注意，名称为“-”的字段仍然可以使用标签“-”生成。

Anonymous struct fields are usually marshaled as if their inner exported fields were fields in the outer struct, subject to the usual Go visibility rules amended as described in the next paragraph. An anonymous struct field with a name given in its JSON tag is treated as having that name, rather than being anonymous. An anonymous struct field of interface type is treated the same as having that type as its name, rather than being anonymous.  
匿名结构字段通常被编组，就好像它们的内部导出字段是外部结构中的字段一样，受制于下一段所述修改的通常的 Go 可见性规则。在其 JSON 标记中给出名称的匿名结构字段被视为具有该名称，而不是匿名的。接口类型的匿名结构字段被视为以该类型作为其名称，而不是匿名的。

Map values encode as JSON objects. The map's key type must either be a string, an integer type, or implement encoding.TextMarshaler. The map keys are sorted and used as JSON object keys by applying the following rules, subject to the UTF-8 coercion described for string values above:  
映射值编码为 JSON 对象。映射的键类型必须是字符串、整数类型或实现 encoding.TextMarshaler。通过应用以下规则对映射键进行排序并用作 JSON 对象键，但要遵守上面为字符串值描述的 UTF-8 强制：

- keys of any string type are used directly. 直接使用任何字符串类型的键
- encoding.TextMarshalers are marshaled. encoding.TextMarshalers 会被编码
- integer keys are converted to strings. 整数键转换为字符串
- Pointer values encode as the value pointed to. A nil pointer encodes as the null JSON value. 指针值编码为指向的值。 nil 指针编码为空 JSON 值。

Interface values encode as the value contained in the interface. A nil interface value encodes as the null JSON value.  
接口值编码为接口中包含的值。一个 nil 接口值编码为空 JSON 值。

Channel, complex, and function values cannot be encoded in JSON. Attempting to encode such a value causes Marshal to return an UnsupportedTypeError.  
通道、复数和函数值不能用 JSON 编码。尝试对此类值进行编码会导致 Marshal 返回 UnsupportedTypeError。

JSON cannot represent cyclic data structures and Marshal does not handle them. Passing cyclic structures to Marshal will result in an error.  
JSON 不能表示循环数据结构并且 Marshal 不处理它们。将循环结构传递给 Marshal 将导致错误。
