环境
- go 1.16.3

# 1

通过接口来针对面向对象编程，通过 goroutine 和 channel 来支持并发和并行编程。

Go 语言有一个被称之为 “没有废物” 的宗旨，就是将一切没有必要的东西都去掉，不能去掉的就无底线地简化，同时追求最大程度的自动化。

Go 以囊地鼠（Gopher）作为它的吉祥物。

Go 语言在这 3 个条件之间做到了最佳的平衡：快速编译，高效执行，易于开发。

Go 语言的主要目标是将静态语言的安全性和高效性与动态语言的易开发性进行有机结合，因此，Go 语言是一门类型安全和内存安全的编程语言。

Go 语言的另一个目标是对于网络通信、并发和并行编程的支持，从而更好地利用大量的分布式和多核的计算机。设计者通过 goroutine 这种轻量级线程的概念来实现这个目标，然后使用 channel 来实现各个 goroutine 之间的通信。

Go 语言中另一个非常重要的特性就是它的构建速度（编译和链接到机器代码的速度），一般情况下构建一个程序的时间只需要数百毫秒到几秒。

Go 语言采用包模型，这个模型通过严格的依赖关系检查机制来加快程序构建的速度，提供了非常好的可量测性。

Go语言通过减少关键字的数量（25 个）来简化编码过程中的混乱和复杂度。

Go 语言有一套完整的编码规范：[The Go Programming Language Specification](https://golang.org/ref/spec) 。

Go 语言使用静态类型，所以它是类型安全的一门语言；作为强类型语言，隐式的类型转换是不被允许的；Go 语言支持交叉编译。

Go 语言一个非常好的目标就是实现所谓的复杂事件处理（[CEP](http://en.wikipedia.org/wiki/Complex_event_processing)）。

因为垃圾回收和自动内存分配的原因，Go 语言不适合用来开发对实时性要求很高的软件。

# 2

Go 从 1.5 版本开始已经实现自举，Go 语言的编译器和链接器都是用 Go 语言编写的。

注意：在创建目录时，文件夹名称永远不应该包含空格，而应该使用下划线 "_" 或者其它一般符号代替。

- `GOROOT` 表示 Go 的安装路径
- `GOPATH` 可以包含多个 Go 语言源码文件、包文件和可执行文件的路径
  - 这些路径下又必须分别包含三个规定的目录：src、pkg 和 bin，这三个目录分别用于存放源码文件、包文件和可执行文件。

1. 下载可执行文件，安装
2. 添加环境变量 `GOROOT`
3. 配置 `GOPROXY`
4. 安装 Goland

Go 拥有简单却高效的标记-清除回收器。Go 的可执行文件都比相对应的源代码文件要大很多，这说明了 Go 的 runtime 嵌入到了每一个可执行文件当中。

# 3

使用 Go 自带的工具来构建应用程序：
- `go build` 编译自身包和依赖包
- `go install` 编译并安装自身包和依赖包

`go fmt`（`gofmt`）这个工具可以将源代码格式化成符合官方统一标准的风格，属于语法风格层面上的小型重构。
- `gofmt -w source.go` 格式化并重写 source.go
- `gofmt -w` 格式化并重写所有 Go 源文件
- `gofmt path` 格式化并重写目录及其子目录下的所有 Go 源文件
- [Command gofmt](https://golang.org/cmd/gofmt/)

`go doc` 工具会从 Go 程序和包文件中提取顶级声明的首行注释以及每个对象的相关注释，并生成相关文档。
- `go doc package` 获取包的文档注释
- `go doc package/subpackage` 获取子包的文档注释
- `go doc package function` 获取某个函数在某个包中的文档注释
- [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc)

这个工具只能获取在 Go 安装目录下 `../go/src` 中的注释内容。此外，它还可以作为一个本地文档浏览 web 服务器。在命令行输入 `godoc -http=:6060`，然后使用浏览器打开 http://localhost:6060 就可以看到本地文档浏览服务器提供的页面。

Go 自带的工具集主要使用脚本和 Go 语言自身编写的，目前版本的 Go 实现了以下三个工具：
- `go install`：主要用于安装非标准库的包文件，将源代码编译成对象文件
- `go fix`：用于将 Go 代码从旧的发行版迁移到最新的发行版
  - Go 在标准库就提供生成抽象语法树和通过抽象语法树对代码进行还原的功能
- `go test`：轻量级的单元测试框架

与其他语言进行交互
- C：cgo 工具
- C++：SWIG（简化封装器和接口生成器）

# 4

驼峰命名法，区分大小写。

`_` 是特殊的标识符，被称为空白标识符。它可以像其他标识符那样用于变量的声明或赋值（任何类型都可以赋值给它），但任何赋给这个标识符的值都将被抛弃，因此这些值不能在后续的代码中使用，也不可以使用这个标识符作为变量对其它变量进行赋值或运算。

25 个关键字/保留字

break	    |default	    |func	    |interface	    |select
case	    |defer	        |go	        |map	        |struct
chan	    |else	        |goto	    |package	    |switch
const	    |fallthrough	|if	        |range	        |type
continue	|for	        |import	    |return	        |var

36 个预定义标识符，其中包含了基本类型的名称和一些基本的内置函数

append	|bool	    |byte	    |cap	    |close	|complex	|complex64	|complex128	|uint16
copy	|false	    |float32	|float64	|imag	|int	    |int8	    |int16	    |uint32
int32	|int64	    |iota	    |len	    |make	|new	    |nil	    |panic	    |uint64
print	|println	|real	    |recover	|string	|true	    |uint	    |uint8	    |uintptr

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号 `()`，中括号 `[]` 和大括号 `{}` 。

程序中可能会使用到这些标点符号：`.`、`,`、`;`、`:` 和 `…` 。

程序的代码通过语句来实现结构化，不需要以分号 `;` 结尾；如果将多个语句写在一行，则必须使用分号 `;` 区分。


包是结构化代码的一种方式：每个程序都由包（通常简称为 pkg）的概念组成，可以使用自身的包或者从其它包中导入内容。每个 Go 文件都属于且仅属于一个包，一个包可以由许多以 .go 为扩展名的源文件组成，因此文件名和包名一般来说都是不相同的。

必须在源文件中非注释的第一行指明这个文件属于哪个包，`package main` 表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 `main` 的包。此外，所有的包名都应该使用小写字母。

在 Go 的安装文件里包含了一些可以直接使用的包，即标准库。

如果想要构建一个程序，则包和包内的文件都必须以正确的顺序进行编译，包的依赖关系决定了其构建顺序。属于同一个包的源文件必须全部被一起编译，一个包就是编译时的一个单元，根据惯例，每个目录都只包含一个包。如果对一个包进行更改或重新编译，所有引用了这个包的客户端程序都必须全部重新编译。

Go 中的包模型采用了显式依赖关系的机制来达到快速编译的目的，编译器会从后缀名为 `.o` 的对象文件（需要且只需要这个文件）中提取传递依赖类型的信息。每一段代码只会被编译一次。

一个 Go 程序是通过 `import` 关键字将一组包链接在一起。导入多个包时，可以使用因式分解关键字，而且最好按照字母顺序排列包名。

```go
import (
   "fmt"
   "os"
)
```

导入包即等同于包含了这个包的所有的代码对象。除了符号 `_` ，包中所有代码对象的标识符必须是唯一的，以避免名称冲突。不过相同的标识符可以在不同的包中使用，因为可以使用包名来区分。

可见性规则
- 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，那么这些标识符对象就可以被外部包的代码所使用
  - Pascal 命名法
- 当标识符以小写字母开头，那它们对包外是不可见的，但是在整个包的内部是可见并且可用的
  - 骆驼命名法，第一个单词的首字母小写，其余单词的首字母大写

因此，在导入一个外部包后，能够且只能够访问该包中导出的对象，即大写字母开头的对象。包也可以作为命名空间使用，帮助避免命名冲突（名称冲突）。

可以在使用 `import` 导入包之后定义或声明 0 个或多个常量（const）、变量（var）和类型（type），这些对象的作用域都是全局的（在本包范围内），所以可以被本包中所有的函数调用，然后声明一个或多个函数（func）。

定义一个函数
- 可以在括号 `()` 中写入 0 个或多个函数的参数（使用逗号 `,` 分隔）
  - 每个参数的名称后面必须紧跟着该参数的类型
- main 函数是每一个可执行程序所必须包含的
  - 一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
  - main 函数既没有参数，也没有返回类型
- 函数里的代码（函数体）使用大括号 `{}` 括起来
  - 左大括号 `{` 必须与方法的声明放在同一行
  - 右大括号 `}` 需要被放在紧接着函数体的下一行，放在同一行也可以

```go
func functionName()

func Sum(a, b int) int { return a + b }

func functionName(parameter_list) (return_value_list) {
   …
}
```

注释不会被编译，但可以通过 godoc 来使用（godoc 工具会收集注释并生成技术文档）
- 单行注释：`//`
- 多行注释：`/*` 开头，`*/` 结尾，不可嵌套

每一个包应该有相关注释，在 package 语句之前的块注释将被默认认为是这个包的文档说明，其中应该提供一些相关信息并对整体功能做简要的介绍。

```go
// Package superman implements methods for saving the world.
//
// Experience has shown that a small number of procedures can prove
// helpful when attempting to save the world.
package superman
```

几乎所有全局作用域的类型、常量、变量、函数和被导出的对象都应该有一个合理的注释。如果这种注释（称为文档注释）出现在函数前面，例如函数 Abcd，则要以 `"Abcd..."` 作为开头。

```go
// enterOrbit causes Superman to fly into low Earth orbit, a position
// that presents several possibilities for planet salvation.
func enterOrbit() error {
   ...
}
```

使用 `var` 声明的变量的值会自动初始化为该类型的零值，结构化的类型没有真正的值，它使用 `nil` 作为默认值。

函数返回值：

```go
// 单个返回值
func FunctionName (a typea, b typeb) typeFunc

// 多个返回值
func FunctionName (a typea, b typeb) (t1 type1, t2 type2)
```

使用 type 关键字可以定义自己的类型，也可以定义一个已经存在的类型的别名
- 这里并不是真正意义上的别名，因为使用这种方法定义之后的类型可以拥有更多的特性，而且在类型转换时必须显式转换

```go
// 别名
type IZ int

// 声明变量
var a IZ = 5

// 因式分解关键字
type (
   IZ int
   FZ float64
   STR string
)
```

Go 程序的执行（程序启动）顺序如下：

1. 按顺序导入所有被 main 包引用的其它包，然后在每个包中执行如下流程：
2. 如果该包又导入了其它的包，则从第一步开始递归执行，但是每个包只会被导入一次。
3. 然后以相反的顺序在每个包中初始化常量和变量，如果该包含有 init 函数的话，则调用该函数。
4. 在完成这一切之后，main 也执行同样的过程，最后调用 main 函数开始执行程序。

在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。精度丢失的情况下将引发运行时错误。

```go
valueOfTypeB = typeB(valueOfTypeA)

// 必须显式说明
a := 5.0
b := int(a)

// 具有相同底层类型的变量之间可以相互转换
var a IZ = 5
c := int(a)
d := IZ(c)
```

返回某个对象的函数或方法的名称一般都是使用名词，如果是用于修改某个对象，则使用 `SetName` 。有必须要的话可以使用大小写混合的方式，如 MixedCaps 或 mixedCaps，而不是使用下划线来分割多个名称。

常量使用关键字 `const` 定义，用于存储不会改变的数据。存储在常量中的数据类型只能是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：`const identifier [type] = value`

```go
const Pi = 3.14159
```

在 Go 语言中，可以省略类型说明符 `[type]`，因为编译器可以根据变量的值来推断其类型。
- 显式类型定义： `const b string = "abc"`
- 隐式类型定义： `const b = "abc"`

一个没有指定类型的常量被使用时，会根据其使用环境而推断出它所需要具备的类型。换句话说，未定义类型的常量会在必要时刻根据上下文来获得相关类型。

```go
var n int
f(n + 5) // 无类型的数字型常量 “5” 它的类型在这里变成了 int
```

常量的值必须是能够在编译时就能够确定的，因为在编译期间自定义函数均属于未知，所以无法用于常量的赋值，但内置函数可以使用，如：`len()` 。

数字型的常量是没有大小和符号的，并且可以使用任何精度而不会导致溢出。

反斜杠 `\` 可以在常量表达式中作为多行的连接符使用。

```go
const Ln2 = 0.693147180559945309417232121458\
			176568075500134360255254120680009
const Log2E = 1/Ln2 // this is a precise reciprocal
const Billion = 1e9 // float constant
const hardEight = (1 << 100) >> 97

// 并行赋值
const beef, two, c = "eat", 2, "veg"
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const (
	Monday, Tuesday, Wednesday = 1, 2, 3
	Thursday, Friday, Saturday = 4, 5, 6
)

// 用作枚举
const (
	Unknown = 0
	Female = 1
	Male = 2
)
```

`iota` 可以被用作枚举值。第一个 `iota` 等于 0，每当 `iota` 在新的一行被使用时，它的值都会自动加 1，没有赋值的常量默认会应用上一行的赋值表达式。

```go
const (
	a = iota    // 0
	b = iota    // 1
	c = iota    // 2
)

// 赋值一个常量时，之后没赋值的常量都会应用上一行的赋值表达式
const (
	a = iota  // a = 0
	b         // b = 1
	c         // c = 2
	d = 5     // d = 5   
	e         // e = 5
)

// 赋值两个常量，iota 只会增长一次，而不会因为使用了两次就增长两次
const (
	Apple, Banana = iota + 1, iota + 2 // Apple=1 Banana=2
	Cherimoya, Durian                  // Cherimoya=2 Durian=3
	Elderberry, Fig                    // Elderberry=3, Fig=4

)

// 使用 iota 结合 位运算 表示资源状态的使用案例
const (
	Open = 1 << iota  // 0001
	Close             // 0010
	Pending           // 0100
)

const (
	_           = iota             // 使用 _ 忽略不需要的 iota
	KB = 1 << (10 * iota)          // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)
```

`iota` 也可以用在表达式中，如：`iota + 50` 。在每遇到一个新的常量块或单个常量声明时， `iota` 都会重置为 0 。简单地讲，每遇到一次 const 关键字，iota 就重置为 0 。

声明变量的一般形式是使用 `var` 关键字：`var identifier type` ，按照从左至右的顺序阅读，代码更加容易理解

```go
var a int
var b bool
var str string

// 因式分解关键字的写法一般用于声明全局变量
var (
	a int
	b bool
	str string
)
```

当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil。所有的内存在 Go 中都是经过初始化的。

声明与赋值（初始化）语句也可以组合起来。

```go
var identifier [type] = value
var a int = 15
var i = 5
var b bool = false
var str string = "Go says hello to the world!"

// 编译时自动推断
var a = 15
var b = false
var str = "Go says hello to the world!"

var (
	a = 15
	b = false
	str = "Go says hello to the world!"
	numShips = 50
	city string
)

// 显示指定类型
var n int64 = 2

// 运行时自动推断
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
)
```

在函数体内声明局部变量时，应使用简短声明语法 `:=`

```go
a := 1
```

值类型的变量，当使用等号 `=` 将一个变量的值赋值给另一个变量时，如：`j = i`，实际上是在内存中将 i 的值进行了拷贝。可以通过 `&i` 来获取变量 i 的内存地址。

引用类型的变量，当使用赋值语句 `r2 = r1` 时，只有引用（地址）被复制。

函数 `Printf` 可以在 fmt 包外部使用，这是因为它以大写字母 P 开头，该函数主要用于打印输出到控制台，通常使用的格式化字符串作为第一个参数。

函数 `fmt.Print` 和 `fmt.Println` 会自动使用格式化标识符 `%v` 对字符串进行格式化，两者都会在每个参数之间自动增加空格，而后者还会在字符串的最后加上一个换行符。

使用操作符 `:=` 可以高效地创建一个新的变量，称之为初始化声明。只能被用在函数体内，而不可以用于全局变量的声明与赋值。

在相同的代码块中，不可以对相同名称的变量再次使用初始化声明，但可以赋予新的值（使用 `=`）。全局变量允许声明但不使用，但局部变量必须先声明后使用，而且必须使用。

同一类型的多个变量可以声明在同一行，多变量可以在同一行进行赋值（并行赋值/同时赋值）。并行赋值也被用于一个函数返回多个返回值的情况。

```go
var a, b, c int

// 声明新变量
a, b, c := 5, 7, "abc"

// 变量已声明
a, b, c = 5, 7, "abc"
```

交换变量的值很容易，空白标识符 `_` 被用于抛弃值

```go
// 交换值
a, b = b, a

// 5 被抛弃
_, b = 5, 7
```

变量除了可以在全局声明中初始化，也可以在 init 函数中初始化。这是一类非常特殊的函数，它不能够被人为调用，而是在每个包完成初始化后自动执行，并且执行优先级比 main 函数高。

每个源文件都只能包含一个 init 函数，初始化总是以单线程执行，并且按照包的依赖关系顺序执行。一个可能的用途是在开始执行程序之前对数据进行检验或修复，以保证程序状态的正确性。

init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine

```go
func init() {
   // setup preparations
   go backend()
}
```

布尔型的值只可以是常量 true 或者 false 。两个类型相同的值可以使用相等 == 或者不等 != 运算符来进行比较并获得一个布尔型的值。

```go
var aVar = 10
aVar == 5 -> false
aVar == 10 -> true

var aVar = 10
aVar != 5 -> true
aVar != 10 -> false
```

Go 对于值之间的比较有非常严格的限制，只有两个类型相同的值才可以进行比较，如果值的类型是接口，它们也必须都实现了相同的接口。如果其中一个值是常量，那么另外一个值的类型必须和该常量类型相兼容的。

Go 语言中包含以下逻辑运算符，与 `&&`、或 `||` 与相等 `==` 或不等 `!=` 属于二元运算符，非 `!` 属于一元运算符。

```go
// 非运算符 !
!T -> false
!F -> true

// 与运算符 &&
T && T -> true
T && F -> false
F && T -> false
F && F -> false

// 或运算符 ||
T || T -> true
T || F -> true
F || T -> true
F || F -> false
```

Go 语言支持整型和浮点型数字，int 型是计算最快的一种类型。整型的零值为 0，浮点型的零值为 0.0 ；float32 精确到小数点后 7 位，float64 精确到小数点后 15 位。
- 整数：
  - int8（-128 -> 127）
  - int16（-32768 -> 32767）
  - int32（-2,147,483,648 -> 2,147,483,647）
  - int64（-9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807）
- 无符号整数：
  - uint8（0 -> 255）
  - uint16（0 -> 65,535）
  - uint32（0 -> 4,294,967,295）
  - uint64（0 -> 18,446,744,073,709,551,615）
- 浮点型（IEEE-754 标准）：
  - float32（+- 1e-45 -> +- 3.4 * 1e38）
  - float64（+- 5 * 1e-324 -> 107 * 1e308）


`math` 包中所有有关数学运算的函数都要求接收 float64 类型。

Go 中不允许不同类型之间的混合使用，但是对于常量的类型限制非常少，因此允许常量之间的混合使用。

```go
// type_mixing.go
package main

func main() {
	var a int
	var b int32
	a = 15
	b = a + a	 // 编译错误
	b = b + 5    // 因为 5 是常量，所以可以通过编译
}
```


格式化说明符
- `%d` 用于格式化整数（`%x` 和 `%X` 用于格式化 16 进制表示的数字）
- `%g` 用于格式化浮点型（`%f` 输出浮点数，`%e` 输出科学计数表示法）
- `%0nd` 用于规定输出长度为 n 的整数，其中开头的数字 0 是必须的
- `%n.mg` 用于表示数字 n 并精确到小数点后 m 位
  - 除了使用 g 之外，还可以使用 e 或者 f

当进行类似 `a32bitInt = int32(a32Float)` 的转换时，小数点后的数字将被丢弃。下面展示了如何安全地进行转换，如果实际存的数字超出要转换到的类型的取值范围的话，则会引发 panic 。

```go
// int -> int8 
func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 { // conversion is safe
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

// float64 -> int
func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 { // x lies in the integer range
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}
```

Go 拥有以下复数类型：
```
complex64 (32 位实数和虚数)
complex128 (64 位实数和虚数)
```

复数使用 `re+imI` 表示，其中 `re` 代表实数部分，`im` 代表虚数部分，`I` 代表根号负 1 。

```go
var c1 complex64 = 5 + 10i
fmt.Printf("The value is: %v", c1)
// 输出： 5 + 10i
```

如果 `re` 和 `im` 的类型均为 float32 ，那么类型为 complex64 的复数 c 可以通过以下方式来获得

```go
c = complex(re, im)

// 实数部分
real(c)

// 虚数部分
imag(c)
```

位运算只能用于整数类型的变量，且需当它们拥有等长位模式时。`%b` 是用于表示位的格式化标识符。
- 二元运算符
  - 按位与 `&`
  - 按位或 `|`
  - 按位异或 `^`
  - 位清除 `&^`：将指定位置上的值设置为 0
- 一元运算符
  - 按位补足 `^`：该运算符与异或运算符一同使用，即 `m^x`，对于无符号 x 使用“全部位设置为 1”，对于有符号 x 时使用 `m=-1`
    - `^10 = -01 ^ 10 = -11`
  - 位左移 `<<`
  - 位右移 `>>`

位左移常见实现存储单位的用例

```go
type ByteSize float64
const (
	_ = iota // 通过赋值给空白标识符来忽略值
	KB ByteSize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)
```

在通讯中使用位左移表示标识的用例

```go
type BitFlag int
const (
	Active BitFlag = 1 << iota // 1 << 0 == 1
	Send // 1 << 1 == 2
	Receive // 1 << 2 == 4
)

flag := Active | Send // == 3
```

Go 中拥有以下逻辑运算符：`==`、`!=`、`<`、`<=`、`>`、`>=`，它们的运算结果总是为布尔值 `bool` 。

```go
b3 := 10 > 5 // b3 is true
```

常见可用于整数和浮点数的二元运算符有 `+`、`-`、`*` 和 `/` ，Go 在进行字符串拼接时允许使用对运算符 `+` 的重载。

`/` 对于整数运算而言，结果依旧为整数，取余运算符 `%` 只能作用于整数。

整数除以 0 可能导致 panic ，浮点数除以 0.0 会返回一个无穷尽的结果，使用 `+Inf` 表示。

可以将语句 `b = b + a` 简写为 `b += a`，同样的写法也可用于 `-=`、`*=`、`/=`、`%=`。

对于整数和浮点数，你可以使用一元运算符 `++`（递增）和 `--`（递减），但只能用于后缀。同时，带有 `++` 和 `--` 的只能作为语句，不能作为表达式，因此 `n = i++` 这种写法是无效的。

```
i++ -> i += 1 -> i = i + 1
i-- -> i -= 1 -> i = i - 1
```

在运算时溢出不会产生错误，Go 会简单地将超出位数抛弃。

`rand` 包实现了伪随机数的生成
- 函数 `rand.Float32` 和 `rand.Float64` 返回介于 [0.0, 1.0) 之间的伪随机数
- 函数 `rand.Intn` 返回介于 [0, n) 之间的伪随机数
- 函数 `rand.Seed(value)` 用于提供伪随机数的生成种子

二元运算符的运算方向均是从左至右，由上至下代表优先级由高到低，可以通过使用括号来临时提升某个表达式的整体运算优先级。

```
优先级 	运算符
 7 		^ !
 6 		* / % << >> & &^
 5 		+ - | ^
 4 		== != < <= >= >
 3 		<-
 2 		&&
 1 		||
```

在 `type TZ int` 中，TZ 就是 int 类型的新名称（用于表示程序中的时区），然后就可以使用 TZ 来操作 int 类型的数据。实际上，类型别名得到的新类型并非和原类型完全相同，新类型不会拥有原类型所附带的方法；TZ 可以自定义一个方法用来输出更加人性化的时区信息。

`byte` 类型是 `uint8` 的别名，对于只占用 1 个字节的传统 ASCII 编码的字符来说，完全没有问题。例如：`var ch byte = 'A'`，字符使用单引号括起来。


在 ASCII 码表中，A 的值是 65，而使用 16 进制表示则为 41，所以下面的写法是等效的：
- `\x` 后总是跟着长度为 2 的 16 进制数
- `\` 后总是跟着长度为 3 的 8 进制数

```go
var ch byte = 65

// 等价
var ch byte = '\x41'
```

Go 同样支持 Unicode（UTF-8），因此字符同样称为 Unicode 代码点或者 runes ，并在内存中使用 int 来表示。
- 在文档中，一般使用格式 U+hhhh 来表示，其中 h 表示一个 16 进制数
- 其实 `rune` 也是 Go 当中的一个类型，并且是 `int32` 的别名

书写 Unicode 字符时，需要在 16 进制数之前加上前缀 `\u` 或者 `\U` ，因为 Unicode 至少占用 2 个字节，所以我们使用 `int16` 或者 `int` 类型来表示。
- 前缀 `\u` 后总是紧跟着长度为 4 的 16 进制数
- 前缀 `\U` 后总是跟着长度为 8 的 16 进制数

包 `unicode` 包含了一些对测试字符非常有用的函数（其中 `ch` 代表字符），这些函数返回一个布尔值。包 `utf8` 包含更多与 `rune` 类型相关的函数。
- 判断是否为字母：`unicode.IsLetter(ch)`
- 判断是否为数字：`unicode.IsDigit(ch)`
- 判断是否为空白符号：`unicode.IsSpace(ch)`

字符串是 UTF-8 字符的一个序列，字符串是一种不可变的值类型，创建文本后无法修改该文本的内容，其实字符串就是字节的定长数组。

Go 支持 2 种形式的字面值
- 解释字符串：该类字符串使用双引号括起来，其中的相关的转义字符将被替换
  - `\n`：换行符
  - `\r`：回车符
  - `\t`：tab 键
  - `\u` 或 `\U`：Unicode 字符
  - `\\`：反斜杠
- 非解释字符串：该类字符串使用反引号括起来，支持换行
```
  `This is a raw string \n` 中的 `\n\` 会被原样输出。
```

Go 中的字符串是根据长度限定，`string` 类型的零值为长度为零的字符串，即空字符串 `""` 。

一般的比较运算符（`==`、`!=`、`<`、`<=`、`>=`、`>`）通过在内存中按字节比较来实现字符串的对比，可以通过函数 `len()` 来获取字符串所占的字节长度。

字符串的内容（纯字节，纯 ASCII 码）可以通过标准索引法来获取，在中括号 `[]` 内写入索引，索引从 0 开始：
- 字符串 str 的第 1 个字节：`str[0]`
- 第 i 个字节：`str[i - 1]`
- 最后 1 个字节：`str[len(str)-1]`

获取字符串中某个字节的地址的行为是非法的，例如：`&str[i]` 。

两个字符串 `s1` 和 `s2` 可以通过 `s := s1 + s2` 拼接在一起，可以通过以下方式来对代码中多行的字符串进行拼接：
- 由于编译器行尾自动补全分号的缘故，加号 + 必须放在第一行

```go
str := "Beginning of the string " +
	"second part of the string"
```

拼接的简写形式 `+=` 也可以用于字符串

```go
s := "hel" + "lo,"
s += "world!"
fmt.Println(s) //输出 “hello, world!”
```

在循环中使用加号 `+` 拼接字符串并不是最高效的做法，更好的办法是使用函数 `strings.Join()` ，更好的办法是使用 `bytes.Buffer` 。

Go 中使用 `strings` 包来完成对字符串的主要操作

```go
// 判断字符串 s 是否以 prefix 开头
strings.HasPrefix(s, prefix string) bool

// 判断字符串 s 是否以 suffix 结尾
strings.HasSuffix(s, suffix string) bool

// 判断字符串 s 是否包含 substr
strings.Contains(s, substr string) bool

// 返回字符串 str 在字符串 s 中首次出现位置的索引，-1 表示不包含
strings.Index(s, str string) int

// 返回字符串 str 在字符串 s 中最后出现位置的索引，-1 表示不包含
strings.LastIndex(s, str string) int

// 查询非 ASCII 编码
strings.IndexRune(s string, r rune) int

// 将字符串 str 中的前 n 个字符串 old 替换为字符串 new，并返回一个新的字符串
// n = -1 替换全部
strings.Replace(str, old, new, n) string

// 计算字符串 str 在字符串 s 中出现的非重叠次数
strings.Count(s, str string) int

// 重复 count 次字符串 s 并返回一个新的字符串
strings.Repeat(s, count int) string

// 将字符串中的 Unicode 字符全部转换为相应的小写字符
strings.ToLower(s) string

// 将字符串中的 Unicode 字符全部转换为相应的大写字符
strings.ToUpper(s) string

// 剔除字符串开头和结尾的空白符号
strings.TrimSpace(s)

// 剔除字符串开头和结尾的指定字符串
strings.Trim(s, "cut")

// 剔除开头或者结尾的指定字符串
strings.TrimLeft(s,, "cut")
strings.TrimRight(s, "cut")

// 利用 1 个或多个空白符号来作为动态长度的分隔符将字符串分割成若干小块，并返回一个 slice
// 如果字符串只包含空白符号，则返回一个长度为 0 的 slice
strings.Fields(s)

// 自定义分割符号来对指定字符串进行分割，同样返回 slice
strings.Split(s, sep)

// 将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串
strings.Join(sl []string, sep string) string
```

函数 `strings.NewReader(str)` 用于生成一个 `Reader` 并读取字符串中的内容，然后返回指向该 `Reader` 的指针，从其它类型读取内容的函数还有：
- `Read()` 从 []byte 中读取内容
- `ReadByte()` 和 `ReadRune()` 从字符串中读取下一个 byte 或者 rune

与字符串相关的类型转换都是通过 `strconv` 包实现的，任何类型转换为字符串总是成功的。

从数字类型转换到字符串：

```go
// 返回数字 i 所表示的字符串类型的十进制数
strconv.Itoa(i int) string

// 将 64 位浮点型的数字转换为字符串
// fmt 表示格式，其值可以是 b、e、f 或 g
// prec 表示精度，bitSize 则使用 32 表示 float32，用 64 表示 float64
strconv.FormatFloat(f float64, fmt byte, prec int, bitSize int) string
```

从字符串类型转换为数字类型：

```go
// 将字符串转换为 int 型
strconv.Atoi(s string) (i int, err error)

// 字符串转换为 float64 型
strconv.ParseFloat(s string, bitSize int) (f float64, err error)
```

利用多返回值的特性，这些函数会返回 2 个值，第 1 个是转换后的结果（如果转换成功），第 2 个是可能出现的错误，因此，我们一般使用以下形式来进行从字符串到其它类型的转换：

```go
val, err = strconv.Atoi(s)
```

`time` 包为提供了一个数据类型 `time.Time`（作为值使用）以及显示和测量时间和日期的功能函数。
- Duration 类型表示两个连续时刻所相差的纳秒数，类型为 int64
- Location 类型映射某个时区的时间，UTC 表示通用协调世界时间

Go 语言的取地址符是 `&`，放到一个变量前使用就会返回相应变量的内存地址。不能获取字面量或常量的地址

```go
const i = 5
ptr := &i       // error: cannot take the address of i
ptr2 := &10     // error: cannot take the address of 10
```

一个指针变量可以指向任何一个值的内存地址，使用一个指针引用一个值被称为间接引用。

可以在指针类型前面加上 `*` 号（前缀）来获取指针所指向的内容，这被称为反引用（或者内容或者间接引用）操作符，另一种说法是指针转移。
- 这里的 `*` 号是一个类型更改器

当一个指针被定义后没有分配到任何变量时，它的值为 `nil` 。一个指针变量通常缩写为 `ptr` 。

对于任何一个变量 var， 该表达式都是正确的：`var == *(&var)`

指针的一个高级应用是可以传递一个变量的引用（如函数的参数），这样不会传递变量的拷贝。指针也可以指向另一个指针，并且可以进行任意深度的嵌套，因此可以有多级的间接引用，但在大多数情况这会使代码的结构不清晰。

对一个空指针的反向引用是不合法的，并且会使程序崩溃。

# 5

Go 提供了下面这些条件结构和分支结构：
- if-else 结构
- switch 结构
- select 结构，用于 channel 的选择

可以使用迭代或循环结构来重复执行一次或多次某段代码（任务）：
- for (range) 结构

一些如 `break` 和 `continue` 这样的关键字可以用于中途改变循环的状态。还可以使用 `return` 来结束某个函数的执行，或者用 `goto` 和标签来调整程序的执行位置。


即使当代码块之间只有一条语句时，大括号也不能省略。关键字 `if` 和 `else` 之后的左大括号 `{` 必须和关键字在同一行，如果使用了 else-if 结构，则前段代码块的右大括号 `}` 必须和 `else-if` 关键字在同一行。

```go
if condition {
	// do something	
}

// 两个分支
if condition {
	// do something	
} else {
	// do something	
}

// 三个分支
if condition1 {
	// do something	
} else if condition2 {
	// do something else	
} else {
	// catch-all or default
}
```

当 if 结构内有 break、continue、goto 或者 return 语句时，Go 代码的常见写法是省略 else 部分。无论满足哪个条件都会返回 x 或者 y 时，一般使用以下写法：

```go
if condition {
	return x
}
return y
```

1. 判断一个字符串是否为空

```go
if str == "" { ... }

// 或
if len(str) == 0 {...}
```

2. 判断运行 Go 程序的操作系统类型

这段代码一般放在 `init()` 函数中执行

```go
if runtime.GOOS == "windows" {
	...
} else { // Unix-like
	...
}
```

3. 函数 `Abs()` 用于返回一个整型数字的绝对值

```go
func Abs(x int) int {
	if x < 0 {
 		return -x
 	}
 	return x	
}
```

4. 函数 `isGreater` 用于比较两个整型数字的大小

```go
func isGreater(x, y int) bool {
	if x > y {
 		return true	
 	}
 	return false
}
```

`if` 可以包含一个初始化语句（如：给一个变量赋值），这种写法具有固定的格式：在初始化语句后方必须加上分号

```go
if initialization; condition {
	// do something
}
```

使用 `:=` 声明的变量的作用域只存在于 if 结构中，并且会遮盖已有的同名变量，最简单的解决方案就是不在初始化语句中声明变量。

```go
val := 10
if val > max {
	// do something
}

// 等价
if val := 10; val > max {
	// do something
}

// 在初始化语句中获取函数 process() 的返回值
if value := process(data); value > max {
	...
}
```

Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败。
- 不使用 true 或 false 的时候，也可以使用一个 error 类型的变量作为第二个返回值：成功执行的话，error 的值为 nil，否则为相应的错误信息
- `comma,ok` 模式（逗号 ok 模式）

```go
value, err := pack1.Function1(param1)
if err != nil {
	fmt.Printf("An error occured in pack1.Function1 with parameter %v", param1)
	return err
}
// 未发生错误，继续执行

if err != nil {
	fmt.Printf("Program stopping with error %v", err)
	os.Exit(1)
}
// 发生错误，终止程序运行
```

可以将错误的获取放置在 if 语句的初始化部分，也可以将 ok-pattern 的获取放置在 if 语句的初始化部分

```go
if err := file.Chmod(0664); err != nil {
	fmt.Println(err)
	return err
}

if value, ok := readData(); ok {
	...
}
```

多返回值的函数必须用相同数量的变量存放结果，否则将得到编译错误 `multiple-value mySqrt() in single-value context` 。

```go
func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 { return } // error case
	return math.Sqrt(f),true
}

func main() {
	// 错误
	t := mySqrt(25.0)
	fmt.Println(t)

	// 正确
	t, ok := mySqrt(25.0)
	if ok { fmt.Println(t) }
}
```

将字符串转换为整数，并且确定转换一定能成功时，可以对 Atoi 函数加一层忽略错误的封装

```go
func atoi (s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}
```

实际上，`fmt` 包最简单的打印函数也有 2 个返回值。当打印到控制台时，可以将该函数返回的错误忽略；但当输出到文件流、网络流等具有不确定因素的输出对象时，应该始终检查是否产生错误。

```go
count, err := fmt.Println(x) // number of bytes printed, nil or 0, error
```

Go 语言中的 switch 结构接受任意形式的表达式。其中的变量必须是相同的类型，或者最终结果为相同类型的表达式。前花括号 `{` 必须和 switch 关键字在同一行。同时测试多个可能符合条件的值时，使用逗号分隔。

```go
switch var1 {
	case val1:
		...
	case val2:
		...
	default:
		...
}
```

每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块。如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。

```go
switch i {
	case 0: // 空分支，只有当 i == 0 时才会进入分支
	case 1:
		f() // 当 i == 0 时函数不会被调用
}

// 继续执行
switch i {
	case 0: fallthrough
	case 1:
		f() // 当 i == 0 时函数也会被调用
}
```

在 `case ...:` 语句之后，不需要使用花括号将多行语句括起来，当代码块只有一行时，可以直接放置在 case 语句之后。

同样可以在 switch 结构中使用 return 语句来提前结束代码块的执行。可选的 default 分支可以出现在任何顺序，但最好放在最后，它表示不符合任何已给出条件时，执行相关语句。


switch 语句的第二种形式是不提供任何被判断的值，然后在每个 case 分支中进行测试不同的条件。当任一分支的测试结果为 true 时，该分支的代码会被执行。这看起来非常像链式的 if-else 语句，但是在测试条件非常多的情况下，提供了可读性更好的书写方式。

```go
switch {
	case condition1:
		...
	case condition2:
		...
	default:
		...
}
```

任何支持进行相等判断的类型都可以作为测试表达式的条件，包括 int、string、指针等。

```go
switch {
	case i < 0:
		f1()
	case i == 0:
		f2()
	case i > 0:
		f3()
}
```

switch 语句的第三种形式包含一个初始化语句。

```go
switch initialization {
	case val1:
		...
	case val2:
		...
	default:
		...
}
```

这种形式可以非常优雅地进行条件判断

```go
switch result := calculate(); {
	case result < 0:
		...
	case result > 0:
		...
	default:
		// 0
}

// 变量 a 和 b 被平行初始化，然后作为判断条件
switch a, b := x[i], y[j]; {
	case a < b: t = -1
	case a == b: t = 0
	case a > b: t = 1
}
```

switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。

如果想要重复执行某些语句，Go 语言中只有 for 结构可以使用。基于计数器的迭代，基本形式为
- 使用分号 `;` 分隔，但不需要括号 `()`
- 左花括号 `{` 必须和 for 语句在同一行，计数器的生命周期在遇到右花括号 `}` 时便终止

```
for 初始化语句; 条件语句; 修饰语句 {}
```

由花括号括起来的代码块会被重复执行已知次数，次数根据计数器决定。可以在循环中同时使用多个计数器（平行赋值）

```go
for i, j := 0, N; i < j; i, j = i+1, j-1 {}
```

for 结构的第二种形式是没有头部的条件判断迭代，基本形式为 `for 条件语句 {}` ，可以看作这是没有初始化语句和修饰语句的 for 结构。

条件语句是可以被省略的，想要直接退出循环体，可以使用 break 语句或 return 语句直接返回。无限循环的经典应用是服务器，用于不断等待和接受新的请求。

```go
for t, err = p.Token(); err == nil; t, err = p.Token() {
	...
}
```

Go 特有的 for-range 结构可以迭代任何集合，一般形式为 `for ix, val := range coll { }` 。`val` 始终为集合中对应索引的值拷贝，因此一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值。
- 如果 `val` 为指针，则会产生指针的拷贝，可以修改集合中的原值

一个字符串是 Unicode 编码的字符（或称之为 `rune`）集合，因此可以用它迭代字符串。每个 `rune` 字符和索引在 for-range 循环中是一一对应的，它能够自动根据 UTF-8 规则识别 Unicode 编码的字符。

```go
for pos, char := range str {
	...
}
```

一个 break 的作用范围为该语句出现后的最内部的结构，可以被用于任何形式的 for 循环（计数器、条件判断等）。但在 switch 或 select 语句中，break 语句的作用结果是跳过整个代码块，执行后续的代码。在嵌套的循环体中，break 只会退出最内层的循环。

关键字 continue 忽略剩余的循环体而直接进入下一次循环的过程，但不是无条件执行下一次循环，执行之前依旧需要满足循环的判断条件。关键字 continue 只能被用于 for 循环中。

for、switch 或 select 语句都可以配合标签（label）形式的标识符使用，即某一行第一个以冒号（`:`）结尾的单词（gofmt 会将后续代码自动移至下一行）。
- 标签的名称是大小写敏感的，为了提升可读性，一般建议全部用大写字母
- 可以使用 goto 语句和标签配合使用来模拟循环
- 定义但未使用标签会导致编译错误 `label … defined and not used`

如果必须使用 goto，应该只用正序的标签（标签位于 goto 语句之后），而且标签和 goto 语句之间不能定义新变量，否则会导致编译失败。

# 6

Go是编译型语言，所以函数编写的顺序是无关紧要的；鉴于可读性的需求，最好把 `main()` 函数写在文件的前面，其他函数按照一定逻辑顺序进行编写（例如函数被调用的顺序）。

DRY 原则：Don't Repeat Yourself

当函数执行到代码块最后一行（`}` 之前）或者 `return` 语句的时候会退出，其中 `return` 语句可以带有零个或多个参数；这些参数将作为返回值供调用者使用。简单的 `return` 语句也可以用来结束 for 死循环，或者结束一个协程（goroutine）。

Go 有三种类型的函数：
- 普通的带有名字的函数
- 匿名函数或者 lambda 函数
- 方法（Methods）

除了 `main()`、`init()` 函数外，其它所有类型的函数都可以有参数和返回值。函数参数、返回值以及它们的类型被统称为函数签名。Go 不支持函数重载（同名函数，签名不同）。

```go
// 定义
func g() {
	...
}

// 调用 pack1 包里的 Function 函数
pack1.Function(arg1, arg2, …, argn)
```

如果需要申明一个在外部定义的函数，只要给出函数名与函数签名，不需要给出函数体

```go
func flushICache(begin, end uintptr) // implemented externally
```

函数也可以以申明的方式被使用，作为一个函数类型，这里不需要函数体 `{}` 。

```go
type binOp func(int, int) int
```

函数是一等值（first-class value）：它们可以赋值给变量，就像 `add := binOp` 一样。函数值（functions value）之间可以相互比较：如果它们引用的是相同的函数或者都是 nil 的话，则认为它们是相同的函数。函数不能在其它函数里面声明（不能嵌套），不过可以使用匿名函数来破除这个限制。

函数能够接收参数供自己使用，也可以返回零个或多个值（我们通常把返回多个值称为返回一组值）。任何一个有返回值（单个或多个）的函数都必须以 `return` 或 `panic` 结尾。

在函数块中，`return` 之后的语句都不会执行。如果一个函数需要返回值，那么这个函数里面的每一个代码分支（code-path）都要有 `return` 语句。

函数定义时，它的形参一般是有名字的，不过也可以定义没有形参名的函数，只有相应的形参类型：`func f(int, int, float64)` 。没有参数的函数通常被称为 niladic 函数（niladic function），就像 `main.main()` 。

Go 默认使用按值传递来传递参数，也就是传递参数的副本。在变量名前加 `&` 按引用传递，此时传递给函数的是一个指针。

在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）。

命名返回值作为结果形参（result parameters）被初始化为相应类型的零值，需要返回时，只要一条简单的不带参数的 `return` 语句。需要注意的是，即使只有一个命名返回值，也要用 `()` 括起来。即使函数使用了命名返回值，也可以无视它而返回明确的值。

任何非命名返回值在 `return` 语句里都要明确指出包含返回值的变量或是一个可计算的值。因此，建议尽量使用命名返回值：会使代码更清晰、更简短，更容易读懂。

空白符 `_` 用来匹配一些不需要的值，然后丢弃掉。

传递指针给函数不但可以节省内存（因为没有复制变量的值），而且赋予了函数直接修改外部变量的能力，所以被修改的变量不需要再用 `return` 返回。

如果函数的最后一个参数是 `...type` 的形式，那么这个函数就可以处理一个变长的参数，长度可以为 0，这样的函数称为变参函数。

```go
func myFunc(a, b, arg ...int) {}
```

在 Greeting 函数中，变量 `who` 的值为 `[]string{"Joe", "Anna", "Eileen"}`

```go
// 变参函数
func Greeting(prefix string, who ...string)

// 调用
Greeting("hello:", "Joe", "Anna", "Eileen")
```

如果参数被存储在一个 slice 类型的变量 `slice` 中，则可以通过 `slice...` 的形式来传递参数，调用变参函数。

如果变长参数的类型不同，有 2 种解决方案：

1. 使用结构

定义一个结构类型，假设为 `Options` ，用于存储所有可能的参数。

函数 F1 可以使用正常的参数 a 和 b，以及一个没有任何初始化的 Options 结构： `F1(a, b, Options {})` 。如果需要对选项进行初始化，则可以使用 `F1(a, b, Options {par1:val1, par2:val2})` 。

```go
type Options struct {
	par1 type1,
	par2 type2,
	...
}
```

2. 使用空接口

如果一个变长参数的类型没有被指定，则可以使用默认的空接口 `interface{}` ，这样就可以接受任何类型的参数。该方案不仅可以用于长度未知的参数，还可以用于任何不确定类型的参数。
- 通常使用 for-range 循环以及 switch 结构对每个参数的类型进行判断

```go
func typecheck(..,..,values … interface{}) {
	for _, value := range values {
		switch v := value.(type) {
			case int: …
			case float: …
			case string: …
			case bool: …
			default: …
		}
	}
}
```

关键字 defer 允许推迟到函数返回之前（或任意位置执行 `return` 语句之后）的那一刻才执行某个语句或函数，一般用于释放某些已分配的资源。

使用 defer 的语句同样可以接受参数，下面这个例子就会在执行 defer 语句时打印 0：

```go
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
```

当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：

```go
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
```

关键字 defer 通常用于进行一些函数执行完成后的收尾工作

```go
// 关闭文件流
// open a file  
defer file.Close()

// 解锁一个加锁的资源
mu.Lock()  
defer mu.Unlock() 

// 打印最终报告
printHeader()  
defer printFooter()

// 关闭数据库链接
// open a database connection  
defer disconnectFromDB()
```

defer 语句可以用来实现代码追踪，也可以用来记录函数的参数与返回值。

一个基础但十分实用的实现代码执行追踪的方案就是在进入和离开某个函数打印相关的消息，可以概括为两个函数：

```go
func trace(s string) { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }
```

Go 语言拥有一些不需要进行导入操作就可以使用的内置函数。

|名称|说明|
|---|---|
|close|用于管道通信|
|len、cap|len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）|
|new、make|new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：`v := new(int)`。make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作 **new() 是一个函数，不要忘记它的括号**|
|copy、append|用于复制和连接切片|
|panic、recover|两者均用于错误处理机制|
|print、println|底层打印函数，在部署环境中建议使用 fmt 包|
|complex、real imag|用于创建和操作复数|

当一个函数在其函数体内调用自身，则称之为递归。最经典的例子便是计算斐波那契数列，即前两个数为1，从第三个数开始每个数均为前两个数之和。

```
1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …
```

在使用递归函数时经常会遇到的一个重要问题就是栈溢出，一般出现在大量的递归调用导致的程序栈内存分配耗尽。这个问题可以通过一个名为[惰性求值](https://zh.wikipedia.org/wiki/%E6%83%B0%E6%80%A7%E6%B1%82%E5%80%BC)的技术解决，在 Go 语言中，可以使用管道（channel）和 goroutine 来实现。

Go 语言中也可以使用相互调用的递归函数：多个函数之间相互调用形成闭环。因为 Go 语言编译器的特殊性，这些函数的声明顺序可以是任意的。

函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调。

将函数作为参数的最好的例子是函数 `strings.IndexFunc()` ，该函数的签名是 `func IndexFunc(s string, f func(c rune) bool) int` ，它的返回值是在函数 `f(c)` 返回 true、-1 或从未返回时的索引值。
- `strings.IndexFunc(line, unicode.IsSpace)` 返回第一个空白字符的索引值

当不想给函数起名字的时候，可以使用匿名函数，例如：`func(x, y int) int { return x + y }`

匿名函数不能独立存在（编译器会返回错误 `non-declaration statement outside function body`），但可以赋给某个变量，即将函数的地址保存到变量中，然后通过变量名对函数进行调用。

表示参数列表的第一对括号必须紧挨着关键字 `func`，因为匿名函数没有名称。花括号 `{}` 涵盖着函数体，最后的一对括号表示对该匿名函数的调用。

```go
func() {
	sum := 0
	for i := 1; i <= 1e6; i++ {
		sum += i
	}
}()
```

关键字 defer 经常配合匿名函数使用，它可以用于改变函数的命名返回值。匿名函数还可以配合 `go` 关键字来作为 goroutine 使用。

匿名函数同样被称之为闭包：它们被允许调用定义在其它环境下的变量。闭包可使得某个函数捕捉到一些外部状态，例如：函数被创建时的状态。另一种表示方式为：一个闭包继承了函数所声明时的作用域。这种状态（作用域内的变量）都被共享到闭包的环境中，因此这些变量可以在闭包中被操作，直到被销毁。闭包经常被用作包装函数：它们会预先定义一个或多个参数用于包装。另一个不错的应用就是使用闭包来完成简洁的错误检查。

闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量，这些局部变量同样可以是参数。

在闭包中使用到的变量可以是在闭包函数体内声明的，也可以是在外部函数声明的。这样闭包函数就能够被应用到整个集合的元素上，并修改它们的值。

```go
var g int
go func(i int) {
	s := 0
	for j := 0; j < i; j++ { s += j }
	g = s
}(1000) // Passes argument 1000 to the function literal.
```

一个返回值为另一个函数的函数可以被称之为工厂函数，在需要创建一系列相似的函数的时候非常有用：书写一个工厂函数而不是针对每种情况都书写一个函数。

```go
// 动态返回追加后缀的函数：
func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 生成函数
addBmp := MakeAddSuffix(".bmp")
addJpeg := MakeAddSuffix(".jpeg")

// 调用函数
addBmp("file") // returns: file.bmp
addJpeg("file") // returns: file.jpeg
```

可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点。函数也是一种值，因此很显然 Go 语言具有一些函数式语言的特性。闭包在 Go 语言中非常常见，常用于 goroutine 和管道操作。

使用闭包调试时，可以用 `runtime` 或 `log` 包中的特殊函数来实现功能，`runtime` 包中的函数 `Caller()` 提供了相应的信息，可以在需要的时候实现一个 `where()` 闭包函数来打印函数执行的位置。

```go
where := func() {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
}
where()
// some code
where()
// some more code
where()
```

也可以设置 `log` 包中的 flag 参数来实现

```go
log.SetFlags(log.Llongfile)
log.Print("")
```

或使用一个更加简短的 `where` 函数实现

```go
var where = log.Print
func func1() {
where()
... some code
where()
... some code
where()
}
```

有时候知道一个计算执行消耗的时间是非常有意义的，尤其是在对比和基准测试中。最简单的测量办法就是在计算开始之前设置一个起始时间，再记录计算结束时的结束时间，它们的差值就是这个计算所消耗的时间。

可以使用 `time` 包中的 `Now()` 和 `Sub` 函数实现。

```go
start := time.Now()
longCalculation()
end := time.Now()
delta := end.Sub(start)
fmt.Printf("longCalculation took this amount of time: %s\n", delta)
```

当在进行大量的计算时，提升性能最直接有效的一种方式就是避免重复计算。通过在内存中缓存和重复利用相同计算的结果，称之为内存缓存。

# 7

容器是可以包含大量条目（item）的数据结构，例如数组、切片和 map 。

数组是具有相同类型的一组已编号且长度固定的数据项序列，数组长度必须是一个常量表达式，并且必须是一个非负整数。数组长度也是数组类型的一部分，所以 `[5]int` 和 `[10]int` 是属于不同类型的。

如果想让数组元素类型为任意类型的话可以使用空接口作为类型，使用值时必须先做一个类型判断。

数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。元素的数目（也称为长度或者数组大小）必须是固定的，并且在声明该数组时就给出（编译时需要知道数组长度以便分配内存）；数组长度最大为 2GB。

数组声明的格式如下：

```go
var identifier [len]type
```

arr1 的长度是 5，索引范围从 0 到 `len(arr1)-1` 。每个元素是一个整型值，声明数组时所有的元素都会被自动初始化为默认值 0。

```go
var arr1 [5]int
```

由于索引的存在，遍历数组的方法自然就是使用 for 结构，for 循环中的条件非常重要，如果写成 `i <= len(arr1)` 会产生越界错误。
- 通过 for 初始化数组项
- 通过 for 打印数组元素
- 通过 for 依次处理元素

```go
// for
for i:=0; i < len(arr1); i++｛
	arr1[i] = ...
}

// for-range
for i,_:= range arr1 {
	...
}
```

Go 语言中的数组是一种值类型，所以可以通过 `new()` 来创建：`var arr1 = new([5]int)`
- `var arr2 [5]int`：arr1 的类型是 `*[5]int`，而 arr2 的类型是 `[5]int`

把一个数组赋值给另一个时，需要做一次数组内存的拷贝操作。这样两个数组就有了不同的值，在赋值后修改 arr2 不会对 arr1 生效。

```go
arr2 := *arr1
arr2[2] = 100
```

如果数组值已经提前知道了，那么可以通过 数组常量 的方法来初始化数组，而不用依次使用 `[]=` 方法

```go
// [10]int {1, 2, 3} 有 10 个元素，除了前三个元素外其他元素都为 0
var arrAge = [5]int{18, 20, 15, 22, 16}

// ... 可以忽略
// 从技术上说其实变成了切片
var arrLazy = [...]int{5, 6, 7, 8, 22}

// key: value 语法
// 只有索引 3 和 4 被赋予实际的值
// 组长度同样可以写成 ...
var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
```

几何点（或者数学向量）是一个使用数组的经典例子，为了简化代码通常使用一个别名

```go
type Vector3D [3]float32
var vec Vector3D
```

Go 语言的多维数组是矩形式的，唯一的例外是切片的数组，多维数组例如 `[3][5]int`，`[2][2][2]float64` 。

把一个大数组传递给函数会消耗很多内存，有两种方法可以避免这种情况：
- 传递数组的指针
- 使用数组的切片

切片（slice）是对数组一个连续片段的引用（称为相关数组，通常是匿名的），所以切片是一个引用类型。切片可以是整个数组，或者是由起始和终止索引标识的子集，终止索引标识的项不包括在切片内。

切片是可索引的，并且可以用 `len()` 函数获取长度。

和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度，因此切片是一个长度可变的数组。

切片提供了计算容量的函数 `cap()`，可以测量切片最长可达到多少：等于切片的长度 + 数组除切片之外的长度。如果 s 是一个切片，`cap(s)` 就是从 `s[0]` 到数组末尾的数组长度。切片的长度永远不会超过它的容量，所以对于切片 s 来说该不等式永远成立：`0 <= len(s) <= cap(s)`。

多个切片如果表示同一个数组的片段，那它们可以共享数据，一个切片和相关数组的其他切片是共享存储的。相反，不同的数组总是代表不同的存储，数组实际上是切片的构建块。

因为切片是引用，所以不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中切片比数组更常用。

声明切片的格式是： `var identifier []type`（不需要说明长度）。一个切片在未初始化之前默认为 nil，长度为 0 。

切片的初始化格式是：`var slice1 []type = arr1[start:end]`
- 表示 slice1 是由数组 arr1 从 start 索引到 end-1 索引之间的元素构成的子集
- `slice1[0]` 就等于 `arr1[start]`

```go
// slice1 等于完整的 arr1 数组
var slice1 []type = arr1[:]
slice1 = arr1[0:len(arr1)]
slice1 = &arr1

// 等价
arr1[2:] == arr1[2:len(arr1)]
arr1[:3] == arr1[0:3]

// 去掉最后一个元素
slice1 = slice1[:len(slice1)-1]

// 由数字 1、2、3 组成的切片
s := [3]int{1,2,3}[:]
s := []int{1,2,3}

// 用切片组成的切片，指向相同的相关数组
s2 := s[:]

// 扩展切片大小
s = s[:cap(s)]
```

对于每一个切片（包括 string），以下状态总是成立的：

```
s == s[:i] + s[i:] // i是一个整数且: 0 <= i <= len(s)
len(s) <= cap(s)
```

切片在内存中的组织方式实际上是一个有 3 个域的结构体：指向相关数组的指针，切片长度以及切片容量。

切片只能向后移动，切片不能被重新分片以获取数组的前一个元素。不要用指针指向 slice，因为切片本身已经是一个引用类型。

如果函数需要对数组进行操作，可能总是要把参数声明为切片。调用该函数时，创建一个切片引用并传递给该函数。

当相关数组还没有定义时，可以使用 `make()` 函数来创建一个切片，同时创建好相关数组：`var slice1 []type = make([]type, len)`
- 可以简写为 `slice1 := make([]type, len)` ，`len` 是数组的长度并且也是切片的初始长度

所以定义 `s2 := make([]int, 10)` ，那么 `cap(s2) == len(s2) == 10` 。


```go
make([]int, 50, 100)

// 等价
new([100]int)[0:50]
```

`new()` 和 `make()` 的区别：new 函数分配内存，make 函数初始化
- new(T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 \*T 的内存地址：这种方法返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体，相当于 `&T{}`
- make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型：切片、map 和 channel

```go
// 分配内存
var p *[]int = new([]int) // *p == nil; with len and cap 0
p := new([]int)

// 初始化
p := make([]int, 0)

// 50 个 int 值的数组
// 长度为 10 容量为 50 的切片
var v []int = make([]int, 10, 50)
v := make([]int, 10, 50)
```

如何理解 new、make、slice、map、channel 的关系
- slice、map 以及 channel 都是 golang 内建的一种引用类型，三者在内存中存在多个组成部分， 需要对内存组成部分初始化后才能使用，而 make 就是对三者进行初始化的一种操作方式
- new 获取的是存储指定变量内存地址的一个变量，对于变量内部结构并不会执行相应的初始化操作， 所以 slice、map、channel 需要 make 进行初始化并获取对应的内存地址，而非 new 简单的获取内存地址

切片通常是一维的，但可以由一维组合成高维，通过分片的分片（或者切片的数组），长度可以任意动态变化，所以 Go 语言的多维切片可以任意切分。不过内层的切片必须通过 make 函数单独分配。

类型 `[]byte` 的切片十分常见，Go 语言有一个 bytes 包专门用来提供这种类型的操作方法。

bytes 包提供 Buffer 类型，长度可变并且支持 Read 和 Write 方法，读写长度未知的 bytes 时使用。

```go
import "bytes"

type Buffer struct {
	...
}
```

Buffer 的定义

```go
// 定义
var buffer bytes.Buffer

// 使用 new 获得一个指针
var r *bytes.Buffer = new(bytes.Buffer)

// 使用函数创建 Buffer 对象，用 buf 初始化
// NewBuffer 最好用在从 buf 读取的时候使用
func NewBuffer(buf []byte) *Buffer
```

通过 Buffer 串联字符串，这种实现方式比使用 `+=` 要更节省内存和 CPU 。

```go
var buffer bytes.Buffer
for {
	if s, ok := getNextString(); ok { //method getNextString() not shown here
		buffer.WriteString(s)
	} else {
		break
	}
}
fmt.Print(buffer.String(), "\n")
```

for-range 结构可以用于数组和切片，`_` 可以用于忽略索引，如果只需要索引，可以忽略第二个变量

```go
for ix, value := range slice1 {
	...
}
```

通过计算行数和矩阵值可以很方便的写出多维切片的 for 循环

```go
for row := range screen {
	for column := range screen[row] {
		screen[row][column] = 1
	}
}
```

切片创建的时候通常比相关数组小，这么做的好处是切片在达到容量上限后可以扩容。
- `start_length` 为切片初始长度
- `capacity` 为相关数组的长度

```go
slice1 := make([]type, start_length, capacity)
```

改变切片长度的过程称之为切片重组（reslicing），做法如下：`slice1 = slice1[0:end]` ，其中 end 是新的末尾索引（即长度）。切片可以反复扩展直到占据整个相关数组。

```go
// 将切片扩展 1 位
sl = sl[0:len(sl)+1]
```

如果想增加切片的容量，必须创建一个新的更大的切片并把原始切片的内容拷贝过来。

`func append(s[]T, x ...T) []T` 其中 append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后并且返回新的切片。如果 s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储。因此，返回的切片可能已经指向一个不同的相关数组了。append 方法总是返回成功，除非系统内存耗尽了。

如果想将切片 y 追加到切片 x 后面，只要将第二个参数扩展成一个列表即可：`x = append(x, y...)`

`func copy(dst, src []T) int copy` 方法将类型为 T 的切片从源地址 src 拷贝到目标地址 dst，覆盖 dst 的相关元素，并且返回拷贝的元素个数。
- 源地址和目标地址可能会有重叠
- 拷贝个数是 src 和 dst 的长度最小值
- 如果 src 是字符串那么元素类型就是 byte
- 如果还想继续使用 src，在拷贝结束后执行 src = dst 即可。

```go
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
```

假设 s 是一个字符串（本质上是一个字节数组），那么可以直接通过 `c := []byte(s)` 来获取一个字节的切片 c ，也可以通过 copy 函数来达到相同的目的：`copy(dst []byte, src string)` 。

可以使用 for-range 来获得字符串中的每个元素。

Unicode 字符会占用 2 个字节，有些甚至需要 3 个或 4 个字节来表示。如果发现错误的 UTF-8 字符，则该字符会被设置为 U+FFFD 并且索引向前移动一个字节。

```go
// 转换为整数
c := []int32(s)

// 转换为 rune 切片
r := []rune(s)

// 字符串中的字符数量
len([]int32(s))
utf8.RuneCountInString(s)	// 更高效

// 将一个字符串追加到某一个字节切片的尾部
var b []byte
var s string
b = append(b, s...)
```

使用 `substr := str[start:end]` 可以从字符串 str 获取到从索引 start 开始到 end-1 位置的子字符串。
- `str[start:]` 则表示获取从 start 开始到 len(str)-1 位置的子字符串
- `str[:end]` 表示获取从 0 开始到 end-1 的子字符串

在内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数。因为指针对用户来说是完全不可见的，因此可以把字符串看做是一个值类型，也就是一个字符数组。

Go 语言中的字符串是不可变的，如果要修改字符串中的某个字符，必须先将字符串转换成字节数组，然后再修改数组中的元素值，最后将字节数组转换回字符串格式。

```go
// 将字符串 "hello" 转换为 "cello"
s := "hello"
c := []byte(s)
c[0] = 'c'
s2 := string(c) // s2 == "cello"
```

返回两个字节数组字典顺序的整数对比结果，即 `0 if a == b, -1 if a < b, 1 if a > b`

```go
func Compare(a, b[]byte) int {
    for i:=0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    // 数组的长度可能不同
    switch {
    case len(a) < len(b):
        return -1
    case len(a) > len(b):
        return 1
    }
    return 0 // 数组相等
}
```

标准库提供了 `sort` 包来实现常见的搜索和排序操作


```go
// 对切片排序
func Ints(a []int)
func Float64s(a []float64)
func Strings(a []string)

// 检查数组是否已排序
func IntsAreSorted(a []int) bool

// 对已排序的数组进行搜索，返回索引值
func SearchInts(a []int, n int) int
func SearchFloat64s(a []float64, x float64) int
func SearchStrings(a []string, x string) int
```

append 函数常见操作

1. 将切片 b 的元素追加到切片 a 之后：`a = append(a, b...)`
2. 复制切片 a 的元素到新的切片 b 上：`b = make([]T, len(a)); copy(b, a)`
3. 删除位于索引 i 的元素：`a = append(a[:i], a[i+1:]...)`
4. 切除切片 a 中从索引 i 至 j 位置的元素：`a = append(a[:i], a[j:]...)`
5. 为切片 a 扩展 j 个元素长度：`a = append(a, make([]T, j)...)`
6. 在索引 i 的位置插入元素 x：`a = append(a[:i], append([]T{x}, a[i:]...)...)`
7. 在索引 i 的位置插入长度为 j 的新切片：`a = append(a[:i], append(make([]T, j), a[i:]...)...)`
8. 在索引 i 的位置插入切片 b 的所有元素：`a = append(a[:i], append(b, a[i:]...)...)`
9. 取出位于切片 a 最末尾的元素 x：`x, a = a[len(a)-1], a[:len(a)-1]`
10. 将元素 x 追加到切片 a：`a = append(a, x)`

从数学的角度来看，切片相当于向量，如果需要的话可以定义一个向量作为切片的别名来进行操作。

切片的底层指向一个数组，该数组的实际容量可能要大于切片所定义的容量。只有在没有任何切片指向的时候，底层的数组内存才会被释放，这种特性有时会导致程序占用多余的内存。

函数 FindDigits 将一个文件加载到内存，然后搜索其中所有的数字并返回一个切片

```go
// 找到第一个匹配正则表达式的数字串
func FindDigits(filename string) []byte {
   b, _ := ioutil.ReadFile(filename)
   b = digitRegexp.Find(b)	// 返回 []byte 指向的底层是整个文件的数据
   c := make([]byte, len(b))
   copy(c, b)
   return c		// 释放底层占用
}

// 找到所有匹配的数字
func FindFileDigits(filename string) []byte {
   fileBytes, _ := ioutil.ReadFile(filename)
   b := digitRegexp.FindAll(fileBytes, len(fileBytes))
   c := make([]byte, 0)
   for _, bytes := range b {
      c = append(c, bytes...)
   }
   return c
}
```

# 8

map 是引用类型，可以使用如下声明。声明的时候不需要知道 map 的长度，map 是可以动态增长的。未初始化的 map 的值是 nil 。
- `[keytype]` 和 `valuetype` 之间允许有空格，但是 gofmt 会移除空格

```go
var map1 map[keytype]valuetype
var map1 map[string]int
```

key 可以是任意可以用 == 或者 != 操作符比较的类型，比如 string、int、float ，所以数组、切片和结构体不能作为 key ，但是指针和接口类型可以。含有数组切片的结构体不能作为 key，只包含内建类型的 struct 是可以作为 key 的。
- 如果要用结构体作为 key ，必须提供 `Key()` 和 `Hash()` 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key。

value 可以是任意类型，使用空接口类型可以存储任意值，但是使用这种类型作为值时需要先做一次类型断言。

通过 key 在 map 中寻找值是很快的，比线性查找快得多，但是仍然比从数组和切片的索引中直接读取要慢。

map 也可以用函数作为自己的值，这样就可以用来做分支结构，key 用来选择要执行的函数。

`v := map1[key1]` 可以将 key1 对应的值赋值给 v；如果 map 中没有 key1 存在，那么 v 将被赋值为 map1 的值类型的空值。

map 可以用 {key1: val1, key2: val2} 的描述方法来初始化，就像数组和结构体一样。map 是引用类型的，内存用 make 方法来分配，永远不要使用 new。

```go
// 初始化
var map1 = make(map[keytype]valuetype)
map1 := make(map[keytype]valuetype)

// 等价
mapCreated := make(map[string]float32)
mapCreated := map[string]float32{}

// 将音阶和对应的音频
noteFrequency := map[string]float32 {
	"C0": 16.35, "D0": 18.35, "E0": 20.60, "F0": 21.83,
	"G0": 24.50, "A0": 27.50, "B0": 30.87, "A4": 440}
```

如果使用 `new()` 分配了一个引用对象，将会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取得了它的地址。

```go
mapCreated := new(map[string]float32)

// 编译错误
// invalid operation: mapCreated["key1"] (index of type *map[string]float32).
mapCreated["key1"] = 4.5
```

map 可以根据新增的 key-value 对动态的伸缩，因此它不存在固定长度或者最大限制，不过也可以选择标明 map 的初始容量 capacity

```go
map2 := make(map[string]float32, 100)
```

一个 key 要对应多个值时，将 value 定义为 `[]int` 类型或者其他类型的切片，就可以优雅地实现。

```go
mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)
```

判断某个 key 是否存在

```go
_, ok := map1[key1] // 如果key1存在则ok == true，否则ok为false

if _, ok := map1[key1]; ok {
	// ...
}
```

从 map1 中删除 key1：`delete(map1, key1)` ，即使 key1 不存在，该操作也不会产生错误。

可以使用 for 循环构造 map 。注意 map 既不是按照 key 的顺序排列，也不是按照 value 的序排列。

```go
for key, value := range map1 {
	...
}

// 只使用值
for _, value := range map1 {
	...
}
```

如果想获取一个 map 类型的切片，那么必须使用两次 `make()` 函数，第一次分配切片，第二次分配切片中每个 map 元素。

如果想对 map 排序，需要将 key 或者 value 拷贝到一个切片，然后再对切片排序，最后用 for-range 打印出所有的 key 和 value 。

如果需要一个排序的列表，最好使用结构体切片：

```go
type name struct {
	key string
	value int
}
```

如果 map 的值类型可以作为 key ，并且所有的 value 是唯一的，那么 key 和 value 可以调换。

# 9

Go 语言中具有常用功能的内置包称为标准库
- `unsafe`：包含了一些打破 Go 语言“类型安全”的命令，一般的程序中不会使用，可用在 C/C++ 程序的调用中。
- `syscall` - `os`-`os/exec`
	- `os`：提供一个平台无关性的操作系统功能接口，采用类 Unix 设计，隐藏了不同操作系统间的差异，让不同的文件系统和操作系统对象表现一致。
	- `os/exec`：提供运行外部操作系统命令和程序的方式。
	- `syscall`：底层的外部包，提供了操作系统底层调用的基本接口。
- `archive/tar` - `/zip-compress`：压缩（解压缩）文件功能。
- `fmt` - `io` - `bufio` - `path/filepath` - `flag`
	- `fmt`：提供了格式化输入输出功能。
	- `io`：提供了基本输入输出功能，大多数对系统功能的封装。
	- `bufio`：缓冲输入输出功能的封装。
	- `path/filepath`：用来操作在当前系统中的目标文件名路径。
	- `flag`：对命令行参数的操作。　　
- `strings` - `strconv` - `unicode` - `regexp` - `bytes`
	- `strings`：提供对字符串的操作。
	- `strconv`：提供将字符串转换为基础类型的功能。
	- `unicode`：为 unicode 型的字符串提供特殊的功能。
	- `regexp`：正则表达式功能。
	- `bytes`：提供对字符型分片的操作。
	- `index/suffixarray`：子字符串快速查询。
- `math` - `math/cmath` - `math/big` - `math/rand-sort`
	- `math`：基本的数学函数。
	- `math/cmath`：对复数的操作。
	- `math/rand`：伪随机数生成。
	- `sort`：为数组排序和自定义集合。
	- `math/big`：大数的实现和计算。 　　
- `container`-`list-ring-heap`：实现对集合的操作。
	- `list`：双链表
	- `ring`：环形链表

```go
// 遍历链表 *List
for e := l.Front(); e != nil; e = e.Next() {
	//do something with e.Value
}
```


- `time` - `log`
  - time: 日期和时间的基本操作。
  - log: 记录程序运行时产生的日志
- `encoding/json` - `encoding/xml` - `text/template`
  - `encoding/json`：读取并解码和写入并编码 JSON 数据。
  - `encoding/xml`：简单的 XML1.0 解析器。
  - `text/template`：生成像 HTML 一样的数据与文本混合的数据驱动模板。
- `net` - `net/http` - `html`
  - `net`：网络数据的基本操作。
  - `http`：提供了一个可扩展的 HTTP 服务器和基础客户端，解析 HTTP 请求和回复。
  - `html`：HTML5 解析器。
- `runtime`：Go 程序运行时的交互操作，例如垃圾回收和协程创建。
- `reflect`：实现通过程序运行时反射，让程序操作任意类型的变量。

在字符串中对正则表达式模式（pattern）进行匹配，简单模式使用 `Match` 和 `MatchString` 方法即可。

```go
ok, _ := regexp.Match(pat, []byte(searchIn))
ok, _ := regexp.MatchString(pat, searchIn)
```

通常需要先将正则模式通过 `Compile` 方法返回一个 Regexp 对象，然后再进行匹配，查找，替换等相关操作。`Compile` 函数也可能返回一个错误，不过一般在使用时忽略对错误的判断。可以用 `MustCompile` 方法检验正则的有效性，当正则不合法时程序将 panic 。

在 Go 语言中锁的机制是通过 `sync` 包中 `Mutex` 来实现的。`sync.Mutex` 是一个互斥锁，用于确保同一时间只能有一个线程进入临界区。

```go
import  "sync"

// 共享变量
type Info struct {
	mu sync.Mutex
	// ... other fields, e.g.: Str string
}

// 修改共享变量
func Update(info *Info) {
	info.mu.Lock()
    // critical section:
    info.Str = // new value
    // end critical section
    info.mu.Unlock()
}
```

使用 Mutex 实现一个可以上锁的共享缓冲区

```go
type SyncedBuffer struct {
	lock 	sync.Mutex
	buffer  bytes.Buffer
}
```

在 `sync` 包中的 `RWMutex` 实现读写锁，通过 RLock() 允许同一时间多个线程对变量进行读操作，但是只允许一个线程进行写操作。

`sync` 包中还有一个方便的 `Once` 类型变量的方法 `once.Do(call)` 用于确保被调用函数只被调用一次。

对于整数的高精度计算， Go 语言提供了 big 包（包含在 math 包中），可以实现任意位的数字，缺点是更大的内存和处理开销使它们使用起来要比内置的数字类型慢很多
- `big.Int`：表示大整数
- `big.Rat`：表示大有理数，例如 2/5 或 3.1416

大的整型数字通过 `big.NewInt(n)` 构造的，其中 n 为 int64 类型整数；大的有理数通过 `big.NewRat(n, d)` 构造，其中 n（分子）和 d（分母）都是 int64 型整数。因为 Go 语言不支持运算符重载，所以所有大数字类型都有像 `Add()` 和 `Mul()` 这样的方法。

包是 Go 语言中代码组织和代码编译的主要方式，当写自己包的时候，要使用短小的不含有 `_`（下划线）的小写单词来为文件命名。

import 的一般格式如下，主程序使用的包必须在主程序编写之前被编译，不同包存放在不同的目录下，每个包（所有属于这个包中的 go 文件）都存放在和包名相同的子目录下。

```
import "包的路径或 URL 地址" 

// 当前目录的相对路径
import "./pack1"

// url
import "github.com/org1/pack1”
```

当使用 `.` 作为包的别名时，可以不通过包名来使用

```go
import . "./pack1"
test := ReturnStr()
```

当使用 `_` 作为包的别名时，只执行包的 init 函数并初始化其中的全局变量

```go
import _ "./pack1"
```

导入外部安装包

```
// 安装
go install codesite.ext/author/goExample/goex

// 安装并导入
import goex "codesite.ext/author/goExample/goex"
```

程序的执行开始于导入包，初始化 main 包然后调用 main 函数。一个没有导入的包将通过分配初始值给所有的包级变量和调用源码中定义的包级 init 函数初始化。一个包可能有多个 init 函数在同一个源码文件中，它们的执行是无序的。

init 函数是不能被调用的，导入的包在包自身初始化前被初始化，而一个包在程序执行中只能初始化一次。

`go get` 可以根据要求和实际情况从互联网上下载或更新指定的代码包及其依赖包，并对它们进行编译和安装；`go install` 用于编译并安装指定的代码包及它们的依赖包。

godoc 能够创建本地浏览的 go 文档站点，用于展示代码包的文档。[Go Walker](https://gowalker.org/) 是一个可以在线生成并浏览 Go 项目 API 文档的 Web 服务器，目前已支持包括 GitHub 等代码托管平台。

```bash
// 安装 godoc
go get -v  golang.org/x/tools/cmd/godoc

// 本地在线浏览 godoc
godoc -http=:6060
```

# 10

Go 通过类型别名（alias types）和结构体的形式支持用户自定义类型，或者叫定制类型。结构体是复合类型（composite types），当需要定义一个由一系列属性组成的类型，并且每个属性都有自己的类型和值时，应该使用结构体把数据聚集在一起。结构体是值类型，因此可以通过 new 函数来创建。

组成结构体类型的那些数据称为字段（fields），每个字段都有一个类型和一个名字；在一个结构体中，字段的名字必须是唯一的。

结构体定义的一般方式如下，不会用到的字段名字可以命名为 `_` 。

```go
type identifier struct {
    field1 type1
    field2 type2
    ...
}
```

`type T struct {a, b int}` 也是合法的语法，更适用于简单的结构体。用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针。

```go
// 定义并分配
var t *T = new(T)

// 先定义后分配
var t *T
t = new(T)

// 习惯用法
t := new(T)

// 类型 T 的实例/对象
var t T
```
使用 `fmt.Println` 打印一个结构体的默认输出可以很好的显示它的内容，类似使用 %v 选项。

可以使用点号符给字段赋值：`structname.fieldname = value`。也可以用来获取结构体字段的值：`structname.fieldname` 。

在 Go 语言中这叫选择器（selector）。无论变量是一个结构体类型还是一个结构体类型指针，都使用同样的选择器符（selector-notation）来引用结构体的字段

```go
type myStruct struct { i int }
var v myStruct    // v 是结构体类型变量
var p *myStruct   // p 是指向一个结构体类型变量的指针
v.i
p.i

// 初始化结构体实例
ms := &struct1{10, 15.5, "Chris"}	// 此时 ms 的类型是 *struct1

// 等价
var ms struct1
ms = struct1{10, 15.5, "Chris"}
```

混合字面量语法（composite literal syntax）`&struct1{a, b, c}` 是一种简写，底层仍然会调用 `new ()` ，这里值的顺序必须按照字段顺序来写。表达式 `new(Type)` 和 `&Type{}` 是等价的。

```go
// 时间间隔
type Interval struct {
    start int
    end   int
}

// 初始化
// (A) 值必须以字段在结构体定义时的顺序给出，& 不是必须的
intr := Interval{0, 3}
// (B) 值的顺序不必一致
intr := Interval{end:5, start:1}
// (C) 某些字段还可以被忽略掉
intr := Interval{end:5}
```

下图说明了结构体类型实例和一个指向它的指针的内存布局

```go
type Point struct { x, y int }
```

使用 new 初始化

![](images/10.1_fig10.1-1.jpg)

作为结构体字面量初始化

![](images/10.1_fig10.1-2.jpg)

可以直接通过指针 `pers2.lastName = "Woodward"` 给结构体字段赋值，也可以通过解指针的方式来设置值：`(*pers2).lastName = "Woodward"` 。

Go 语言中，结构体和它所包含的数据在内存中是以连续块的形式存在的，即使结构体中嵌套有其他的结构体，这在性能上带来了很大的优势。

```go
type Rect1 struct {Min, Max Point }
type Rect2 struct {Min, Max *Point }
```

![](images/10.1_fig10.2.jpg)


结构体类型可以通过引用自身来定义，在定义链表或二叉树的元素（通常叫节点）时特别有用，此时节点包含指向临近节点的链接（地址）。

链表中的第一个元素叫 `head` ，它指向第二个元素；最后一个元素叫 `tail` ，它没有后继元素，所以它的 `su` 为 nil 值。

```go
// 链表
type Node struct {
    data    float64		// 存放数据
    su      *Node		// 指向后继节点
}

// 双向链表
type Node struct {
    pr      *Node		// 前驱节点
    data    float64
    su      *Node		// 后继节点
}
```

![](images/10.1_fig10.3.jpg)


二叉树中每个节点最多能链接至两个节点：左节点（`le`）和右节点（`ri`），这两个节点本身又可以有左右节点，以此类推。树的顶层节点叫根节点（`root`），底层没有子节点的节点叫叶子节点（leaves），叶子节点的 `le` 和 `ri` 指针为 nil 值。

```go
// 二叉树
type Tree struct {
    le      *Tree		// 左节点
    data    float64
    ri      *Tree		// 右节点
}
```

![](images/10.1_fig10.4.jpg)


Go 中的类型转换遵循严格的规则，当为结构体定义了一个 alias 类型时，此结构体类型和它的 alias 类型都有相同的底层类型，非法赋值或转换将引发编译错误。

在 Go 语言中常常在工厂方法里使用初始化来实现构造函数，工厂的名字以 new 或 New 开头。
- `File` 是一个结构体类型，表达式 `new(File)` 和 `&File{}` 是等价的

```go
// 结构体类型
type File struct {
    fd      int     // 文件描述符
    name    string  // 文件名
}

// 对应的工厂方法，返回指向结构体实例的指针
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }

    return &File{fd, name}
}

// 调用
f := NewFile(10, "test.txt")
```

想知道结构体类型 T 的一个实例占用了多少内存，可以使用 `size := unsafe.Sizeof(T{})` 。

强制使用工厂方法

```go
type matrix struct {
    ...
}

func NewMatrix(params) *matrix {
    m := new(matrix) // 初始化 m
    return m
}
```

在其他包里使用工厂方法

```go
package main
import "matrix"
...
wrong := new(matrix.matrix)     // 编译失败（matrix 是私有的）
right := matrix.NewMatrix(...)  // 实例化 matrix 的唯一方式
```

试图 `make()` 一个结构体变量会引发一个编译错误，但是 `new()` 一个 map 并试图向其填充数据，将会引发运行时错误。因为 `new(Foo)` 返回的是一个指向 nil 的指针，它尚未被分配内存。

结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）：它是一个附属于字段的字符串，可以是文档或其他的重要标记。标签的内容只有 `reflect` 包能获取，`reflect` 包可以在运行时自省类型、属性和方法。在一个变量上调用 `reflect.TypeOf()` 可以获取变量的类型，如果变量是一个结构体类型，就可以通过 Field 来索引结构体的字段，然后就可以使用 Tag 属性。

结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字。匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体。不过，在一个结构体中对于每一种数据类型只能有一个匿名字段。

Go 语言中的继承是通过内嵌或组合来实现的。内层结构体被简单的插入或者内嵌进外层结构体，即可以从另外一个或一些类型继承部分或全部实现。

当两个字段拥有相同的名字（可能是继承来的名字）时：

1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；

2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发错误（不使用没关系）。没有办法解决这种问题引起的二义性，必须由程序员修正。

```go
// c.a 引发编译错误
type A struct {a int}
type B struct {a, b int}
type C struct {A; B}
var c C

// d.b, d.B.b 都可以访问
type D struct {B; b float32}
var d D
```

Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量，因此方法是一种特殊类型的函数。接收者类型可以是（几乎）任何类型，结构体类型、函数类型、int、bool、string 或数组的别名类型。但不能是接口类型，因为接口是一个抽象定义，而方法却是具体实现，如果这样做会引发编译错误。另外，接收者也不能是指针类型，但可以是任何其他允许类型的指针。

一个类型加上它的方法等价于面向对象中的一个类。在 Go 中，类型的代码和绑定的方法的代码可以不放在一起，它们在不同的源文件中，但必须是同一个包的。

类型 T（或 *T）上的所有方法的集合叫做类型 T（或 *T）的方法集（method set）。

因为方法是函数，所以不允许重载，对于一个类型只能有一个给定名称的方法。基于接收者类型的可以重载，具有同样名字的方法可以在多个不同的接收者类型上存在。

```go
func (a *denseMatrix) Add(b Matrix) Matrix
func (a *sparseMatrix) Add(b Matrix) Matrix
```

别名类型没有原始类型上已经定义过的方法，定义方法的一般格式如下

在方法名之前，`func` 关键字之后的括号中指定 receiver
- 如果 recv 是 receiver 的实例，Method1 是它的方法名，那么方法调用遵循传统的 `object.name` 选择器符号：recv.Method1()。
- 如果 recv 是一个指针，Go 会自动解引用。
- 如果方法不需要使用 recv 的值，可以用 `_` 替换。

```go
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }

func (_ receiver_type) methodName(parameter_list) (return_value_list) { ... }
```

可以先定义类型的别名类型，然后再为别名类型定义方法，或者将它作为匿名类型嵌入在一个新的结构体中，当然方法只在这个别名类型上有效。

函数和方法的区别
- 函数将变量作为参数：`Function1(recv)`
- 方法在变量上被调用：`recv.Method1()`

在接收者是指针时，方法可以改变接收者的值（或状态）；当参数作为指针传递，即通过引用调用时，函数也可以改变参数的状态。

接收者必须有一个显式的名字，这个名字必须在方法中被使用。

receiver_type 叫做 （接收者）基本类型，这个类型必须在和方法同样的包中被声明。

在 Go 中，（接收者）类型关联的方法不写在类型结构里面，类型和方法之间的关联由接收者来建立。

方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。

鉴于性能的原因，recv 最常见的是一个指向 receiver_type 的指针，特别是在 receiver 类型是结构体时。

如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法，否则，就在普通的值类型上定义方法。

```go
type Point3 struct { x, y, z float64 }
// A method on Point3
func (p Point3) Abs() float64 {
    return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

// 指针
p3 := &Point{ 3, 4, 5}

// 等价方法
p3.Abs()
(*p3).Abs()
```

指针方法和值方法都可以在指针或非指针上被调用，Go 做了探测工作。

为了修改没有被导出的字段，可以提供 getter 和 setter 方法，对于 setter 方法使用 Set 前缀，对于 getter 方法只使用成员名。

并发访问对象，可以使用 `sync` 包中的方法。


当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型继承了这些方法：将父类型放在子类型中来实现亚型。这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果。

```go
// 接口类型
type Engine interface {
	Start()
	Stop()
}

// 结构体类型
type Car struct {
	Engine		// 匿名字段
}

// 方法
func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}
```

内嵌将一个已存在类型的字段和方法注入到了另一个类型里：匿名字段上的方法“晋升”成为了外层类型的方法。当然类型可以有只作用于本身实例而不作用于内嵌“父”类型上的方法。

可以覆写方法（像字段一样）：和内嵌类型方法具有同样名字的外层类型的方法会覆写内嵌类型对应的方法。

因为一个结构体可以嵌入多个匿名类型，所以实际上我们可以有一个简单版本的多重继承。

结构体内嵌和自己在同一个包中的结构体时，可以彼此访问对方所有的字段和方法。

主要有两种方法来实现在类型中嵌入功能：
- 聚合（或组合）：包含一个所需功能类型的具名字段
- 内嵌：内嵌（匿名地）所需功能类型

内嵌的类型不需要指针，内嵌类型可以嵌入其他类型，那些类型的方法可以直接在外层类型中使用。

在 Go 语言中，如果方法在此类型定义了，就可以调用它，和其他类型上是否存在这个方法没有关系。Go 不需要一个显式的类定义，类是通过提供一组作用于一个共同类型的方法集来隐式定义的，类型可以是结构体或者任何用户自定义类型。

```go
// 定义自己的 Integer 类型
type Integer int

// 添加转换成字符串的方法
func (i *Integer) String() string {
    return strconv.Itoa(int(*i))
}
```

在 Go 中，类型就是类（数据和关联的方法），代码复用通过组合和委托实现，多态通过接口的使用来实现：有时这也叫组件编程（Component Programming）。

使用一个自定义类型时，最好为它定义 String()方法，格式化描述符 `%T` 会给出类型的完全规格，`%#v` 会给出实例的完整输出，包括它的字段。

不要在 `String()` 方法里调用涉及 `String()` 方法的方法，它会导致意料之外的错误，比如下面的例子导致了无限递归调用，很快就会导致内存溢出：

```go
type TT float64

func (t TT) String() string {
    return fmt.Sprintf("%v", t)
}

t.String()
```

使用 `runtime` 包可以获取当前的内存状态：

```go
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("%d Kb\n", m.Alloc / 1024)
```

如果需要在一个对象 obj 被从内存移除前执行一些特殊操作，比如写到日志文件中，可以通过如下方式调用函数来实现：

`func(obj *typeObj)` 需要一个 `typeObj` 类型的指针参数 `obj`，特殊操作会在它上面执行，func 也可以是一个匿名函数。

在对象被 GC 进程选中并从内存中移除以前，SetFinalizer 都不会执行，无论程序是正常结束还是发生错误。

```go
runtime.SetFinalizer(obj, func(obj *typeObj))
```

# 11

接口提供了一种方式来说明对象的行为：如果谁能搞定这件事，它就可以用在这儿。接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的），接口里也不能包含变量。

```go
// 定义接口
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```

Go 语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。接口的名字由方法名加 `[e]r` 后缀组成，例如 `Printer`、`Reader`、`Writer`、`Logger`、`Converter` 等，还有一些不常用的方式，例如 `Recoverable`，此时接口名以 `able` 结尾，或者以 `I` 开头）。

在 Go 语言中接口可以有值，一个接口类型的变量或一个接口值：`var ai Namer`
- `ai` 是一个多字（multiword）数据结构，它的值是 nil
- 此处的方法指针表是通过运行时反射能力构建的

类型（比如结构体）可以实现某个接口的方法集，这个实现可以描述为，该类型的变量上的每一个具体方法所组成的集合，包含了该接口的方法集。

实现了 `Namer` 接口的类型的变量可以赋值给 `ai`（即 `receiver` 的值），方法表指针（method table ptr）就指向了当前的方法实现。当另一个实现了 `Namer` 接口的类型的变量被赋给 `ai`，`receiver` 的值和方法表指针也会相应改变。

![](11.1_fig11.1.jpg)


类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

实现某个接口的类型（除了实现接口方法外）可以有其他的方法。一个类型可以实现多个接口。接口类型可以包含一个实例的引用， 该实例的类型实现了此接口（接口是动态类型）。

即使接口在类型之后才定义，二者处于不同的包中，被单独编译，只要类型实现了接口中的方法，它就实现了此接口。

io 包里有一个接口类型 Reader

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

// r 右边的类型都实现了 Read() 方法，并且有相同的方法签名
// r 的静态类型是 io.Reader
var r io.Reader
r = os.Stdin    // see 12.1
r = bufio.NewReader(r)
r = new(bytes.Buffer)
f,_ := os.Open("test.txt")
r = bufio.NewReader(f)
```

一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。

```go
type ReadWrite interface {
    Read(b Buffer) bool
    Write(b Buffer) bool
}

type Lock interface {
    Lock()
    Unlock()
}

// 包含了 ReadWrite 和 Lock 的所有方法
// 额外有一个 Close() 方法
type File interface {
    ReadWrite
    Lock
    Close()
}
```

一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的动态类型，即运行时在变量中存储的值的实际类型。在执行过程中动态类型可能会有所不同，但是它总是可以分配给接口变量本身的类型。

可以使用类型断言来测试在某个时刻 varI 是否包含类型 T 的值，如果 varI 不是接口变量，编译器会报错。

```go
v := varI.(T)       // unchecked type assertion
```

如果转换在程序运行时失败会导致错误发生，更安全的方式是使用以下形式来进行类型断言，如果转换合法，v 是 varI 转换到类型 T 的值，ok 会是 true；否则 v 是类型 T 的零值，ok 是 false，没有运行时错误发生。

```go
if v, ok := varI.(T); ok {  // checked type assertion
    Process(v)
    return
}
// varI is not of type T


// 简单写法
if _, ok := varI.(T); ok {
    // ...
}
```

接口变量的类型也可以使用一种特殊形式的 switch 来检测：`type-switch`。可以用 `type-switch` 进行运行时类型分析，但是不允许有 `fallthrough` 。

变量 t 得到了 areaIntf 的值和类型，所有 case 语句中列举的类型（nil 除外）都必须实现对应的接口，如果被检测类型没有在 case 语句列举的类型中，就会执行 default 语句。

```go
switch t := areaIntf.(type) {
case *Square:
	fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
	fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
	fmt.Printf("nil value: nothing to check?\n")
default:
	fmt.Printf("Unexpected type %T\n", t)
}

// 仅测试变量类型，可以不用赋值语句
switch areaIntf.(type) {
case *Square:
	// TODO
case *Circle:
	// TODO
...
default:
	// TODO
}
```

类型分类函数，它有一个可变长度参数，可以是任意类型的数组，它会根据数组元素的实际类型执行不同的动作。

```go
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

// 调用
classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
```

在处理来自于外部的、类型未知的数据时，比如解析 JSON 或 XML 编码的数据，类型测试和转换会非常有用。

Print 函数检测类型是否可以打印

```go
type Stringer interface {
    String() string
}

// 测试 v 是否实现 Stringer 接口
if sv, ok := v.(Stringer); ok {
    fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
```

接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。接口将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现出不同的行为，这就是多态的本质。

编写参数是接口变量的函数，这使得它们更具有一般性，使用接口使代码更具有普适性。

在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以根据具体类型 P 辨识的：
- 指针方法可以通过指针调用
- 值方法可以通过值调用
- 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
- 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

将一个值赋值给一个接口时，编译器会确保所有可能的接口方法都能在此值上被调用，因此不正确的赋值在编译期间就会失败。

Go 语言规范定义了接口方法集的调用规则：
- 类型 T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
- 类型 *T 的可调用方法集包含接受者为 *T 的所有方法
- 类型 *T 的可调用方法集不包含接受者为 T 的方法

要对一组数字或字符串排序，只需实现三个方法：反映元素个数的 `Len()` 方法、比较第 i 和 j 个元素的 `Less(i, j)` 方法以及交换第 i 和 j 个元素的 `Swap(i, j)` 方法。

同样的原理，排序函数可以用于浮点型数组，字符串数组，或者表示一周各天的结构体。

```go
// 接口类型
type Sorter interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// 冒泡排序
func Sort(data Sorter) {
    for pass := 1; pass < data.Len(); pass++ {
        for i := 0;i < data.Len() - pass; i++ {
            if data.Less(i+1, i) {
                data.Swap(i, i + 1)
            }
        }
    }
}

// 对 int 数组排序
type IntArray []int
func (p IntArray) Len() int           { return len(p) }
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// 调用排序函数
data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
a := sort.IntArray(data) //conversion to type IntArray from package sort
sort.Sort(a)
```

对于基本类型的排序，标准库已经提供了相关的排序函数，对于一般性的排序，`sort` 包定义了一个接口，总结了需要用于排序的抽象方法。

```go
type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}
```

`io` 包提供了用于读和写的接口 `io.Reader` 和 `io.Writer`。`io` 包里的 Readers 和 Writers 都是不带缓冲的，`bufio` 包里提供了对应的带缓冲的操作，在读写 UTF-8 编码的文本文件时尤其有用。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

只要类型实现了读写接口，提供 Read 和 Write 方法，就可以从它读取数据，或向它写入数据。
- 一个对象要是可读的，它必须实现 `io.Reader` 接口，这个接口只有一个签名是 `Read(p []byte) (n int, err error)` 的方法，它从调用它的对象上读取数据，并把读到的数据放入参数中的字节切片中，然后返回读取的字节数和一个 error 对象，如果没有错误发生返回 nil，如果已经到达输入的尾端，会返回 `io.EOF("EOF")`，如果读取的过程中发生了错误，就会返回具体的错误信息。
- 一个对象要是可写的，它必须实现 `io.Writer` 接口，这个接口也只有一个签名是 `Write(p []byte) (n int, err error)` 的方法，它将指定字节切片中的数据写入调用它的对象里，然后返回实际写入的字节数和一个 error 对象（如果没有错误发生就是 nil）。

空接口（最小接口）不包含任何方法，它对实现不做任何要求。任何其他类型都实现了空接口，`any` 或 `Any` 是空接口一个很好的别名或缩写。

```go
type Any interface {}
```

可以给一个空接口类型的变量 `var val interface {}` 赋任何类型的值。每个 `interface {}` 变量在内存中占据两个字长：一个用来存储它包含的类型，另一个用来存储它包含的数据或者指向数据的指针。

给空接口定一个别名类型 Element：`type Element interface{}` ，然后定义一个容器类型的结构体 Vector，它包含一个 Element 类型元素的切片。Vector 里能放任何类型的变量，因为任何类型都实现了空接口，实际上 Vector 里放的每个元素可以是不同类型的变量，因此需要使用类型断言。

```go
type Vector struct {
	a []Element
}

// 返回第 i 个元素
func (p *Vector) At(i int) Element {
	return p.a[i]
}

// 设置第 i 个元素
func (p *Vector) Set(i int, e Element) {
	p.a[i] = e
}
```

假设有一个 `myType` 类型的数据切片，直接将切片中的数据复制到空接口切片中会引发编译错误，因为它们在内存中的布局是不一样的，必须使用 for-range 显式赋值。

```go
var dataSlice []myType = FuncReturnSlice()

// 编译错误
var interfaceSlice []interface{} = dataSlice

// 正确做法：显式赋值
var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
for i, d := range dataSlice {
    interfaceSlice[i] = d
}
```

一个接口的值可以赋值给另一个接口变量，只要底层类型实现了必要的方法。这个转换是在运行时进行检查的，转换失败会导致运行时错误。

```go
var ai AbsInterface // declares method Abs()
type SqrInterface interface {
    Sqr() float
}
var si SqrInterface
pp := new(Point) // say *Point implements Abs, Sqr
var empty interface{}

// 合法
empty = pp                // everything satisfies empty
ai = empty.(AbsInterface) // underlying value pp implements Abs()
// (runtime failure otherwise)
si = ai.(SqrInterface) // *Point has Sqr() even though AbsInterface doesn’t
empty = si             // *Point implements empty set
// Note: statically checkable so type assertion not necessary.
```

x 转换为 myPrintInterface 类型是完全动态的：只要 x 的底层类型（动态类型）定义了 print 方法这个调用就可以正常运行。若 x 的底层类型未定义 print 方法，此处类型断言会导致 panic，最佳实践应该为 `if mpi, ok := x.(myPrintInterface); ok { mpi.print() }`。

```go
// 函数调用
type myPrintInterface interface {
	print()
}

func f3(x myInterface) {
	x.(myPrintInterface).print() // type assertion to myPrintInterface
}
```

反射是用程序检查其所拥有的结构，尤其是类型的一种能力；这是元编程的一种形式。反射可以在运行时检查类型和变量，例如它的大小、方法和动态的调用这些方法。

变量的最基本信息就是类型和值：反射包的 `Type` 用来表示一个 Go 类型，反射包的 `Value` 为 Go 值提供了反射接口。

实际上，反射是通过检查一个接口的值，变量首先被转换成空接口。接口的值包含一个 type 和 value 。反射可以从接口值反射到对象，也可以从对象反射回接口值。

```go
func TypeOf(i interface{}) Type
func ValueOf(i interface{}) Value

// 检查类型
type MyInt int
var m MyInt = 5
v := reflect.ValueOf(m)

// 总是返回底层类型，reflect.Int
v.Kind()

// 还原（接口）值
fmt.Println(v.Interface())
```

当 `v := reflect.ValueOf(x)` 函数通过传递一个 x 拷贝创建了 v，那么 v 的改变并不能更改原始的 x。要想 v 的更改能作用到 x，那就必须传递 x 的地址 `v = reflect.ValueOf(&x)`。

`Elem()` 方法返回接口 v 所包含或指针 v 所指向的值，可以用 `CanSet()` 方法测试反射值是否可设置，最后再用 `Set...` 方法设置新的值。

有些时候需要反射一个结构类型，`v.NumField()` 方法返回结构内的字段数量，通过一个 for 循环用索引取得每个字段的值 `v.Field(i)`。如果要调用签名在结构上的方法，可以使用索引 n 来调用 `Method(n).Call(nil)`。注意，结构体中只有被导出字段（首字母大写）是可以设置的。

Printf 的函数声明为 `func Printf(format string, args ... interface{}) (n int, err error)` ，其中 `...` 参数为空接口类型，Printf 使用反射包来解析这个参数列表。

在 Go 语言中，任何提供了接口方法实现代码的类型都隐式地实现了该接口，而不用显式声明。Go 是唯一结合了接口值，静态类型检查（是否该类型实现了某个接口），运行时动态转换的语言，并且不需要显式声明类型满足某个接口。该特性允许程序员能够在不改变已有的代码的情况下定义和使用新接口。

接收一个（或多个）接口类型作为参数的函数，实参可以是任何实现了该接口的类型的变量，实现了某个接口的类型可以被传给任何以此接口为参数的函数。

Go 通常需要编译器静态检查的支持：当变量被赋值给一个接口类型的变量时，编译器会检查其是否实现了该接口的所有函数。如果方法调用作用于像 `interface{}` 这样的“泛型”上，可以通过类型断言来检查变量是否实现了相应接口。

```go
// XML 写接口
type xmlWriter interface {
	WriteXML(w io.Writer) error
}

// Exported XML streaming function.
func StreamXML(v interface{}, w io.Writer) error {
	if xw, ok := v.(xmlWriter); ok {
		// It’s an  xmlWriter, use method of asserted type.
		return xw.WriteXML(w)
	}
	// No implementation, so we have to use our own function (with perhaps reflection):
	return encodeToXML(v, w)
}

// Internal XML encoding function.
func encodeToXML(v interface{}, w io.Writer) error {
	// ...
}
```

Go 的接口提高了代码的分离度，改善了代码的复用性，使得代码开发过程中的设计模式更容易实现，用 Go 接口还能实现依赖注入模式。

提取接口是非常有用的设计模式，可以减少需要的类型和方法数量，而且不需要维护整个的类层次结构。

Go 接口可以让开发者找出自己写的程序中的类型。假设有一些拥有共同行为的对象，并且开发者想要抽象出这些行为，这时就可以创建一个接口来使用。整个设计可以持续演进，而不用废弃之前的决定。类型要实现某个接口，它本身不用改变，只需要在这个类型上实现新的方法。


如果希望满足某个接口的类型显式地声明它们实现了这个接口，可以向接口的方法集中添加一个具有描述性名字的方法。大部分代码并不使用这样的约束，因为它限制了接口的实用性，但是有些时候，这样的约束在大量相似的接口中被用来解决歧义。

```go
type Fooer interface {
	Foo()
	ImplementsFooer()
}

// 必须实现 ImplementsFooer 方法以满足 Fooer 接口
type Bar struct{}
func (b Bar) ImplementsFooer() {}
func (b Bar) Foo() {}
```

在 Go 语言中函数重载可以用可变参数 `...T` 作为函数最后一个参数来实现。如果把 T 换为空接口，那么任何类型的变量都是满足 T (空接口）类型的，这样就允许传递任何数量任何类型的参数给函数，这就是 Go 语言中的函数重载。

函数 `fmt.Printf` 就是这样做的，通过枚举 slice 类型的实参动态确定所有参数的类型，并查看每个类型是否实现了 `String()` 方法，如果是就用于生成输出信息。

```go
fmt.Printf(format string, a ...interface{}) (n int, errno error)
```

当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，这个类型就可以使用（另一个类型）所有的接口方法。

```go
// 结构体
type Task struct {
	Command string
	*log.Logger
}

// 工厂方法
func NewTask(command string, logger *log.Logger) *Task {
	return &Task{command, logger}
}

// log.Logger 实现 Log 方法
// Task 实例 task 就可以调用
task.Log()
```

类型可以通过继承多个接口来提供像多重继承一样的特性：

```go
type ReaderWriter struct {
	*io.Reader
	*io.Writer
}
```

有用的接口可以在开发的过程中被归纳出来。添加新接口非常容易，因为已有的类型不用变动（仅仅需要实现新接口的方法）。已有的函数可以扩展为使用接口类型的约束性参数：通常只有函数签名需要改变。

OO 语言最重要的三个方面分别是：封装，继承和多态
- 封装（数据隐藏）：和别的 OO 语言有 4 个或更多的访问层次相比，Go 把它简化为了 2 层:
	- 包范围内的：通过标识符首字母小写，对象 只在它所在的包内可见
	- 可导出的：通过标识符首字母大写，对象 对所在包以外也可见

类型只拥有自己所在包中定义的方法。
- 继承：用组合实现。内嵌一个（或多个）包含想要的行为（字段和方法）的类型，多重继承可以通过内嵌多个类型实现
- 多态：用接口实现。某个类型的实例可以赋给它所实现的任意接口类型的变量。类型和接口是松耦合的，并且多重继承可以通过实现多个接口实现。

# 12

