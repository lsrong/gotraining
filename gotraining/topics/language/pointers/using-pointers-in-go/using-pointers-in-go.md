# 在 Go 中使用指针
- William Kennedy  

## 引言
我经常被问到什么时候什么时候不应该在 Go 中使用指针。大多数人面临的问题是，他们试图根据他们认为的性能权衡来做出这个决定。因此，问题是，不要根据您对性能可能有的毫无根据的想法做出编码决定。根据惯用、简单、可读和合理的代码做出编码决策。

我对指针的使用是基于我在查看标准库中的代码时的发现。这些规则总是有例外的，但我将向您展示的是常见做法。它首先对需要共享的值的类型进行分类。这些类型分类是内置、结构和引用类型。让我们分别看一看。

## 内置类型(Built-in-types)
Go 的 [内置类型](http://golang.org/ref/spec#Types) 表示原始数据值，它们是管理和处理数据的构建块。我将这些类型统称为布尔、数字和字符串类型的集合。在声明接受这些类型值的函数和方法时，标准库很少与指针共享它们。

让我们先看看isShellSpecialVar来自 [env包](http://golang.org/src/os/env.go) 的函数：

**清单 1**
```
38 func isShellSpecialVar(c uint8) bool {
39  switch c {
40      case '*', '#', '$', '@', '!', '?', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
41      return true
42  }
43  return false
44 }
```

清单 1 中的isShellSpecialVar函数被声明为接受uint8类型的值并返回bool类型的值。要使调用者使用此函数，他们必须将其uint8类型值的副本传递到该函数中。这对于返回值也是一样的。函数的bool类型值的副本正在返回给调用者。

接下来，让我们看看同一个env包中的getShellName函数：

**清单 2**
```
54 func getShellName(s string) (string, int) {
55 switch {
56 case s[0] == '{':
      ...
66 return "", 1 // Bad syntax; just eat the brace.
67 case isShellSpecialVar(s[0]):
68 return s[0:1], 1
69 }
      ...
74 return s[:i], i
75 }
```


清单 2 中的getShellName函数被声明为接受string类型的值并返回两个值，一个是string类型，另一个是int类型。一个字符串是内置型转到一个特殊的表示字节的不变片。由于此切片无法增长，因此容量值与其 [slice header](https://www.ardanlabs.com/blog/2013/08/understanding-slices-in-go-programming.html) 无关。最好将string类型的值与将 boolean 和 numeric 类型值相同的方式视为原始数据值。

当调用getShellName 时，调用者会将其字符串值的副本传递给函数。该函数生成一个新的字符串值并将该值的副本返回给调用者。传入和传出此函数的所有值都是原始值的副本。

这种共享字符串值副本的做法在 [strings](http://golang.org/src/strings/strings.go) 包中非常普遍：

**清单 3**

```
620 func Trim(s string, cutset string) string {
621 if s == "" || cutset == "" {
622     return s
623 }
624     return TrimFunc(s, makeCutsetFunc(cutset))
625 }
```


strings包中的所有函数都接受调用者字符串值的副本并返回它们创建的字符串值的副本。清单 3 显示了Trim函数的实现。该函数接受两个字符串值的副本，并返回传入的第一个字符串值的副本或已修剪出割集的新字符串值的副本。

如果您查看标准库中共享内置类型值的更多代码，您将看到这些值很少与指针共享。如果函数或方法需要更改内置类型的值，通常会将反映该更改的新值返回给调用者。

**通常，不要与指针共享内置类型值**。

## 结构类型(struct type)
[结构类型](http://golang.org/ref/spec#Struct_types) 类型允许通过将不同类型组合在一起来创建复杂的数据类型。这是通过组合一系列字段来实现的，每个字段都有一个名称和一个类型。它们还支持[嵌入](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) ，这增加了结构类型的组合方式。

可以实现结构类型以使其行为类似于内置类型。当它们是时，你应该这样对待它们。要查看表现为原始数据值的结构类型，我们可以查看[time](http://golang.org/src/time/time.go) 包：

**清单 4**

```
39 type Time struct {
40     // sec gives the number of seconds elapsed since
41     // January 1, year 1 00:00:00 UTC.
42     sec int64
43
44     // nsec specifies a non-negative nanosecond
45     // offset within the second named by Seconds.
46     // It must be in the range [0, 999999999].
47     nsec int32
48
49     // loc specifies the Location that should be used to
50     // determine the minute, hour, month, day, and year
51     // that correspond to this Time.
52     // Only the zero Time has a nil Location.
53     // In that case it is interpreted to mean UTC.
54     loc *Location
55 }
```

清单 4 显示了Time结构体类型。此类型表示时间并已实现为原始数据值。如果查看工厂函数Now，您将看到它返回一个Time类型的值，而不是一个指针：

**清单 5**
```
781 func Now() Time {
782     sec, nsec := now()
783     return Time{sec + unixToInternal, nsec, Local}
784 }
```


清单 5 显示了Now函数如何返回Time类型的值。这表明Time类型的值可以安全复制，并且是共享它们的首选方式。

接下来，让我们看一个用于更改Time值的方法：

**清单 6**
```
610 func (t Time) Add(d Duration) Time {
611     t.sec += int64(d / 1e9)
612     nsec := int32(t.nsec) + int32(d%1e9)
613     if nsec >= 1e9 {
614         t.sec++
615         nsec -= 1e9
616     } else if nsec < 0 {
617         t.sec–
618         nsec += 1e9
619     }
620     t.nsec = nsec
621     return t
622 }
```

就像我们在使用内置类型时看到的那样，清单 6 显示了如何针对调用者的Time值的副本调用Add方法。该方法更改接收者值的本地副本并将该更改的副本返回给调用者。接受时间值的函数也接受这些值的副本：

**清单 7**

```
1118 func div(t Time, d Duration) (qmod2 int, r Duration) {
```

清单 7 显示了div函数的声明，该函数接受Time和Duration类型的值。同样，Time类型的值被视为原始数据类型，并在共享时进行复制。

**大多数情况下，struct 类型不是为了表现得像原始数据类型而创建的。在这些情况下，使用指针共享值是更好的方法**。让我们看一个来自os包的示例：

**清单 8**
```
238 func Open(name string) (file *File, err error) {
239 return OpenFile(name, O_RDONLY, 0)
240 }
```

在清单 8 中，我们看到了 os 包中的Open函数。它打开一个文件进行读取，并返回一个指向File类型值的指针。接下来，让我们看看UNIX 平台的File结构类型的声明：  
**清单 9**
```
15 // File represents an open file descriptor.
16 type File struct {
17     *file
18 }
19
20 // file is the real representation of *File.
21 // The extra level of indirection ensures that no clients of os
22 // can overwrite this data, which could cause the finalizer
23 // to close the wrong file descriptor.
24 type file struct {
25     fd int
26     name string
27     dirinfo *dirInfo // nil unless directory being read
28     nepipe int32 // number of consecutive EPIPE in Write
29 }
```

我在清单 9 中留下了这些类型声明的注释，因为它们确实带回了我想要表达的观点。当你有一个像Open这样的工厂函数为你提供一个指针时，这是一个好兆头，你不应该复制返回的引用值。Open正在返回一个指针，因为复制所返回的引用File值是不安全的。该值应始终通过指针使用和共享。

即使函数或方法没有改变File结构类型值的状态，它仍然需要与指针共享。让我们看一下UNIX 平台os包中的epipecheck函数：

**清单 10**
```
58 func epipecheck(file *File, e error) {
59     if e == syscall.EPIPE {
60         if atomic.AddInt32(&file.nepipe, 1) >= 10 {
61             sigpipe()
62         }
63     } else {
64         atomic.StoreInt32(&file.nepipe, 0)
65     }
66 }
```


在清单 10 中，epipecheck函数接受File类型的指针。因此，调用者通过指针与函数共享其文件类型值。请注意，epipecheck函数不会更改File值的状态，而是使用它来执行其操作。

这也适用于为File类型声明的方法：  

**清单 11**
```
224 func (f *File) Chdir() error {
225 if f == nil {
226 return ErrInvalid
227 }
228 if e := syscall.Fchdir(f. fd); e != nil {
229 return &PathError{"chdir", f.name, e}
230 }
231 return nil
232 }
```

清单 11 中的Chdir方法使用指针接收器来实现该方法，并且不会更改接收器值的状态。在所有这些情况下，要共享File类型的值，必须使用指针来完成。一个文件值不是原始数据值。

如果您查看标准库中的更多代码，您将看到 struct 类型如何实现为与内置类型一样的原始数据值，或者实现为需要与指针共享且永不复制的值。给定结构类型的工厂函数将为您提供有关如何实现该类型的重要线索。

**通常，使用指针共享结构类型值，除非结构类型已被实现为表现得像原始数据值**

*如果您仍然不确定，这是另一种思考方式。将每个结构视为具有性质。如果结构的性质是不应该改变的，比如时间、颜色或坐标，那么将结构实现为原始数据值。如果结构的性质是可以改变的，即使它从来不在你的程序中，它也不是原始数据值，应该实现为与指针共享。不要创建具有性质二元性的结构* 。

## 引用类型
引用类型是切片、映射、通道、接口和函数值。这些值包含通过指针和其他元数据引用底层数据结构的标头值。我们很少与指针共享引用类型值，因为标头值旨在复制。标头值已经包含一个指针，该指针默认为我们共享底层数据结构。

让我们看一个来自net包的例子：  
**清单 12**
```
32 type IP []byte
```

清单 12 显示了来自名为IP的net包的命名类型，其基本类型是字节片。当您需要围绕内置或引用类型声明行为时，使用命名类型是有价值的。让我们看看IP命名类型的MarshalText方法：

**清单 13**
```
329 func (ip IP) MarshalText() ([]byte, error) {
330     if len(ip) == 0 {
331         return []byte(""), nil
332     }
333     if len(ip) != IPv4len && len(ip) != IPv6len {
334         return nil, errors.New("invalid IP address")
335     }
336     return []byte(ip.String()), nil
337 }
```

在清单 13 中，我们可以看到MarshalText方法如何使用值接收器。这正是我希望看到的，因为我们不与指针共享引用类型。如果您查看net包中为IP命名类型声明的其余方法，您将看到使用了更多值接收器。这适用于将引用类型值共享为函数和方法的参数：

**清单 14**

```
318 // ipEmptyString 与 ip.String 类似，不同之处在于它返回
319 // 未设置 ip 时的空字符串。
320 func ipEmptyString(ip IP) string {
321 if len(ip) == 0 {
322 return ""
323 }
324 return ip.String()
325 }
```


清单 14 中的ipEmptyString函数接受命名类型IP的值。没有指针用于共享此值，因为IP的基本类型是字节片，因此是引用类型。

不与指针共享引用类型的规则有一个常见的例外：

**清单 15**

```
341 func (ip *IP) UnmarshalText(text []byte) error {
342 if len(text) == 0 {
343 *ip = nil
344 return nil
345 }
346 s := string(text)
347 x := ParseIP(s)
348 if x == nil {
349 return &ParseError{"IP address", s}
350 }
351 *ip = x
352 return nil
353 }
```
任何时候将数据解组(unmarshal)为引用类型时，都需要使用指针共享该引用类型值。清单 15 显示了UnmarshalText方法，该方法执行解组操作并使用指针接收器声明。编码包中的Decode和Unmarshal函数也期望接收指向引用类型的指针。如果您查看标准库中的更多代码，您将看到在大多数情况下引用类型的值是如何不与指针共享的。由于引用类型包含一个标题值，其目的是共享底层数据结构，因此不需要用指针共享这些值。已经有一个指针在使用。

**通常，除非您要实现解组(unmarshal)类型的功能，否则不要与指针共享引用类型值。**

## 值的切片
我可以避免的一件事是使用指针切片存储数据。当我从数据库、Web 甚至文件中检索数据时，我会将这些数据存储在一个值片段中：

**清单 16**

```
10 func FindRegion(s *Service, region string) ([]BuoyStation, error) {
11 var bs []BuoyStation
12 f := func(c *mgo.Collection) error {
13 queryMap := bson.M{"region" : region}
14 return c.Find(queryMap).All(&bs)
15 }
16
17 if err := s.DBAction(cfg.Database, "buoy_stations", f); err != nil {
18 return nil, err
19 }
20
21 return bs, nil
22 }
```


下面是我的一个项目中清单 16 中的一些代码，它通过mgo包调用 MongoDB 数据库。在第 14 行，我将bs切片的地址传递给All方法。在所有的方法进行解组的呼叫下，为片上创建的值。然后通过将切片标头值的副本传递回调用者来返回数据值切片。

使用值切片允许将程序的数据存储在连续的内存块中。这意味着我使用的更多核心数据可以一次被 CPU 缓存，并希望在缓存中停留更长时间。如果我创建一个指针切片，则无法保证这些核心数据值的内存是连续的，只有指向这些值的指针会连续存储。尽管我在考虑这种情况下的性能，但我认为它更符合习惯。

有时这是不可能的。想象一下，如果我需要一部分文件类型值。由于我无法复制文件类型值，因此我需要创建一个文件切片类型指针。当使用标准库中的结构类型而不是您自己的结构类型时，这种情况经常发生。

**通常，尽可能创建值的切片和映射。**

## 结论
标准库在如何根据所使用的值类型共享值方面是相当一致的。
- 除非有特殊需要，否则不要使用具有内置数据类型的指针。
- 结构类型具有二元性。如果 struct 类型实现为原始数据类型，则不要使用指针。如果没有，则使用指针共享这些值。
- 最后，除了极少数例外，不应与指针共享引用类型。

我想在结束时重申其他三个想法。首先，根据代码的惯用性、简单性、可读性和合理性来做出编码决策。其次，这不是对与错，想想你正在编写的代码以及你所做决定背后的原因。最后，将每种情况和场景视为个案，尽量不要应用一揽子模式或解决方案。
