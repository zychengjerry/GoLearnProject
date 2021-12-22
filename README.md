# GoLearnProject

## GoLand安装及环境配置

* golang下载安装
  > https://golang.google.cn/
    version: go1.17.5



* GoLand下载安装
  > https://www.jetbrains.com/go/download/other.html
    目前最新版：
    Version: 2021.3
    Build: 213.5744.269
    Released: 2 December 2021

* golang及GoLand环境配置参考：
  > https://blog.csdn.net/qq_44702847/article/details/108597386?spm=1001.2101.3001.6650.2&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-2.nonecase&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-2.nonecase


# Go语言基础

## Go - 数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go 语言按类别有以下几种数据类型：

    序号	类型和描述
    1	布尔型：布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。
    2	数字类型：整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。
    3	字符串类型：字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
    4	派生类型:
            包括：
            (a) 指针类型（Pointer）
            (b) 数组类型
            (c) 结构化类型(struct)
            (d) Channel 类型
            (e) 函数类型
            (f) 切片类型
            (g) 接口类型（interface）
            (h) Map 类型

数字类型：
    Go 也有基于架构的类型，例如：int、uint 和 uintptr。
    
    序号	类型和描述
    1	uint8     无符号 8 位整型 (0 到 255)
    2	uint16    无符号 16 位整型 (0 到 65535)
    3	uint32    无符号 32 位整型 (0 到 4294967295)
    4	uint64    无符号 64 位整型 (0 到 18446744073709551615)

    5	int8      有符号 8 位整型 (-128 到 127)
    6	int16     有符号 16 位整型 (-32768 到 32767)
    7	int32     有符号 32 位整型 (-2147483648 到 2147483647)
    8	int64     有符号 64 位整型 (-9223372036854775808 到 9223372036854775807)

浮点型:

    序号	类型和描述
    1	float32     IEEE-754 32位浮点型数
    2	float64     IEEE-754 64位浮点型数
    3	complex64   32 位实数和虚数
    4	complex128  64 位实数和虚数

其他数字类型：

    序号	类型和描述
    1	byte      类似 uint8
    2	rune      类似 int32
    3	uint      32 或 64 位
    4	int       与 uint 一样大小
    5	uintptr   无符号整型，用于存放一个指针

## Go - 变量

- Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。

  声明变量的一般形式是使用 var 关键字：
`var identifier type`
可以一次声明多个变量：
`var identifier1, identifier2 type`


- 变量声明: 第一种，指定变量类型，如果没有初始化，则变量默认为零值。


    数值类型（包括complex64/128）为 0
    布尔类型为 false
    字符串为 ""（空字符串）
    以下几种类型为 nil：
    var a *int
    var a []int
    var a map[string] int
    var a chan int
    var a func(string) int
    var a error // error 是接口

- 第二种，根据值自行判定变量类型。 `var v_name = value`

- 第三种，如果变量已经使用 var 声明过了，再使用 := 声明变量，就产生编译错误，格式：`v_name := value`

  例如：
    
      var intVal int
      intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

  直接使用下面的语句即可：

      intVal := 1 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
      
  intVal := 1 相等于：
    
      var intVal int
      intVal =1 


## Go - 常量

常量是一个简单值的标识符，在程序运行时，不会被修改的量。常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：

    const identifier [type] = value
    你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。

    显式类型定义： const b string = "abc"
    隐式类型定义： const b = "abc"

多个相同类型的声明可以简写为：

`const c_name1, c_name2 = value1, value2`

常量还可以用作枚举：

    const (
    Unknown = 0
    Female = 1
    Male = 2
    )

常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过

**iota**，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

iota 可以被用作枚举值：

    const (
    a = iota
    b = iota
    c = iota
    )

第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

    const (
    a = iota
    b
    c
    )


## Go - 运算符

运算符用于在程序运行时执行数学或逻辑运算。

Go 语言内置的运算符有：

- 算术运算符: `+ - * / % ++ --`
- 关系运算符: `== != > < >= <=`
- 逻辑运算符: `&& || !`
- 位运算符: `& | ^ << >>`
- 赋值运算符: `= += -= *= /= %= <<= >>= &= |= ^=`
- 其他运算符: `& *`

有些运算符拥有较高的**优先级**，二元运算符的运算方向均是从左至右。下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

    优先级	运算符
    5	* / % << >> & &^
    4	+ - | ^
    3	== != < <= > >=
    2	&&
    1	||

