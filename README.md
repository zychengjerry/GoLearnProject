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

## Go 0 - 数据类型 Data type

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

## Go 1 - 变量 Var

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


## Go 2-3 - 常量 Const

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


## Go 4-9 - 运算符 Operator

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


## Go 10 - 条件语句 If

条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。


    语句	        描述
    if 语句	        if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
    if...else 语句	if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
    if 嵌套语句	你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。
    switch 语句	switch 语句用于基于不同条件执行不同动作。
    select 语句	select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。


## Go 11 - 循环语句 For

在不少实际问题中有许多具有规律性的重复操作，因此在程序中就需要重复执行某些语句。

Go 语言提供了以下几种类型循环处理语句：

    循环类型	        描述
    for 循环	重复执行语句块
    循环嵌套	        在 for 循环中嵌套一个或多个 for 循环

**循环控制语句**

循环控制语句可以控制循环体内语句的执行过程。

GO 语言支持以下几种循环控制语句：

    控制语句	        描述
    break 语句	经常用于中断当前 for 循环或跳出 switch 语句
    continue 语句	跳过当前循环的剩余语句，然后继续进行下一轮循环。
    goto 语句	将控制转移到被标记的语句。

**无限循环**

如果循环中条件语句永远不为 false 则会进行无限循环，我们可以通过 for 循环语句中只设置一个条件表达式来执行无限循环


## Go 12a - 函数 Function

函数是基本的代码块，用于执行一个任务。

Go 语言最少有个 main() 函数。

你可以通过函数来划分不同功能，逻辑上每个函数执行的是指定的任务。

函数声明告诉了编译器函数的名称，返回类型，和参数。

Go 语言标准库提供了多种可动用的内置的函数。例如，len() 

**函数定义**

Go 语言函数定义格式如下：

    func function_name( [parameter list] ) [return_types] {
        函数体
    }

函数定义解析：
    
    func：           函数由 func 开始声明
    function_name：  函数名称，参数列表和返回值类型构成了函数签名。
    parameter list： 参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
    return_types：   返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
    函数体：         函数定义的代码集合。


**函数参数**

函数如果使用参数，该变量可称为函数的形参。
形参就像定义在函数体内的局部变量。

调用函数，可以通过两种方式来传递参数：

    传递类型	  描述
    值传递	  值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
    引用传递	  引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。

默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。


**函数用法**

    函数用法	                     描述
    函数作为另外一个函数的实参     函数定义后可作为另外一个函数的实参数传入
    闭包	                     闭包是匿名函数，可在动态编程中使用
    方法	                     方法就是一个包含了接受者的函数



Go 语言可以很灵活的创建函数，并作为另外一个函数的实参。


Go 语言支持匿名函数，可作为**闭包**。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。


Go语言中同时有函数和**方法**。一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。所有给定类型的方法属于该类型的方法集。

    func (variable_name variable_data_type) function_name() [return_type]{
        /* 函数体*/
    }


## Go 12b - 变量作用域 Variable Scope


作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：

    - 函数内定义的变量称为局部变量
    - 函数外定义的变量称为全局变量
    - 函数定义中的变量称为形式参数


在函数体内声明的变量称之为**局部变量**，它们的作用域只在函数体内，参数和返回值变量也是局部变量。


在函数体外声明的变量称之为**全局变量**，全局变量可以在整个包甚至外部包（被导出后）使用。

    package main
    
    import "fmt"
    
    /* 声明全局变量 */
    var g int

    func main() {
      /* 声明局部变量 */
      var a, b, c int
    
      /* 初始化参数 */
      a = 10
      b = 20
      c = a + b
    
      fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
    }


Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑: 

    package main
    
    import "fmt"
    
    /* 声明全局变量 */
    var g int = 20
    
    func main() {
      /* 声明局部变量 */
      var g int = 10
      
      fmt.Printf ("结果： g = %d\n",  g)
    }

    result: g = 10


**形式参数**会作为函数的局部变量来使用。

    package main
    
    import "fmt"
    
    /* 声明全局变量 */
    var a int = 20;
    
    func main() {
      /* main 函数中声明局部变量 */
      var a int = 10
      var b int = 20
      var c int = 0
      
      fmt.Printf("main()函数中 a = %d\n",  a);
      c = sum( a, b);
      fmt.Printf("main()函数中 c = %d\n",  c);
    }
    
    /* 函数定义-两数相加 */
    func sum(a, b int) int {
      fmt.Printf("sum() 函数中 a = %d\n",  a);
      fmt.Printf("sum() 函数中 b = %d\n",  b);
      return a + b;
    }

    result: a = 10    a = 10    b = 20    c = 30


## Go 13 - 数组 Array

数组是具有相同唯一类型的一组已编号且长度固定的数据项序列，这种类型可以是任意的原始类型例如整型、字符串或者自定义类型。

相对于去声明 number0, number1, ..., number99 的变量，使用数组形式 numbers[0], numbers[1] ..., numbers[99] 更加方便且易于扩展。

数组元素可以通过索引（位置）来读取（或者修改），索引从 0 开始，第一个元素索引为 0，第二个索引为 1，以此类推。

Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

    var variable_name [SIZE] variable_type

Eg. 初始化数组：

    var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

我们也可以通过字面量在声明数组的同时快速初始化数组：

    balance := [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：

    var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
    或
    balance := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

如果设置了数组的长度，我们还可以通过指定下标来初始化元素：

    //  将索引为 1 和 3 的元素初始化
    balance := [5]float32{1:2.0,3:7.0}

初始化数组中 {} 中的元素个数不能大于 [] 中的数字。

如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小：

    balance[4] = 50.0


**多维数组**

以下为常用的多维数组声明方式：

    var variable_name [SIZE1][SIZE2]...[SIZEN] 

    Eg: 
    var threedim [5][10][4]int    // 三维整型数组

**二维数组**

是最简单的多维数组，二维数组本质上是由一维数组组成的。二维数组定义方式如下：

    var arrayName [ x ][ y ] variable_type

variable_type 为 Go 语言的数据类型，arrayName 为数组名，二维数组可认为是一个表格，x 为行，y 为列。
二维数组中的元素可通过 a[ i ][ j ] 来访问。

**二维数组初始化**

多维数组可通过大括号来初始值。以下实例为一个 3 行 4 列的二维数组：

    a := [3][4]int{  
      {0, 1, 2, 3} ,   /*  第一行索引为 0 */
      {4, 5, 6, 7} ,   /*  第二行索引为 1 */
      {8, 9, 10, 11},   /* 第三行索引为 2 */
    }
    
    注意：
    以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样：
    a := [3][4]int{  
      {0, 1, 2, 3} ,   /*  第一行索引为 0 */
      {4, 5, 6, 7} ,   /*  第二行索引为 1 */
      {8, 9, 10, 11}}   /* 第三行索引为 2 */

Eg: 初始化一个2*2的二维数组
    
    // 创建二维数组
    sites := [2][2]string{}

    // 向二维数组添加元素
    sites[0][0] = "Name: "
    sites[0][1] = "Jerry"
    sites[1][0] = "ID: "
    sites[1][1] = "xxxxxxx"

**访问二维数组**

二维数组通过指定坐标来访问。如数组中的行索引与列索引，例如：

    val := a[2][3]
    或
    var value int = a[2][3]

以上实例访问了二维数组 val **第三行的第四个**元素。


**向函数传递数组**

形参设定数组大小：`void myFunction(param [10]int){...}`
形参未设定数组大小：`void myFunction(param []int){...}`


## Go 14 - 指针 Pointer

变量是一种使用方便的占位符，用于引用计算机内存地址。Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

Eg. 

    func main() {
    var a int = 10
    fmt.Printf("变量的地址: %x\n", &a  )
    }

    输出：
    变量的地址: c000016088


一个指针变量指向了一个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

    var var_name *var-type
    var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

    eg：
    var ip *int        /* 指向整型*/
    var fp *float32    /* 指向浮点型 */
    // 这是一个指向 int 和 float32 的指针。


**指针使用流程：**

1. 定义指针变量。
2. 为指针变量赋值。
3. 访问指针变量中指向地址的值。

在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

**空指针**

当一个指针被定义后没有分配到任何变量时，它的值为 nil。

nil 指针也称为空指针。

nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。

Eg：

    func main() {
    var  ptr *int
    fmt.Printf("ptr 的值为 : %x\n", ptr)
    }
    
    ptr 的值为 : 0

空指针判断：

    if(ptr != nil)     /* ptr 不是空指针 */
    if(ptr == nil)    /* ptr 是空指针 */

**指针的应用**

    内容	            描述
    指针数组	           你可以定义一个指针数组来存储地址
    指向指针的指针	   Go支持指向指针的指针
    向函数传递指针参数   通过引用或地址传参，在函数调用时可以改变其值

如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。

当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址。

指向指针的指针变量声明格式： `var ptr **int`

以上指向指针的指针变量为整型。访问指向指针的指针变量值需要使用两个 * 号。


## Go 15 - 结构体 Type - struct

Go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。

结构体是由一系列具有相同类型或不同类型的数据构成的数据集合。

结构体表示一项记录，比如保存图书馆的书籍记录，每本书有以下属性：
Title ：标题 || 
Author ： 作者 || 
Subject：学科 || 
ID：书籍ID


**结构体定义**

需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

    type struct_variable_type struct {
    member definition
    member definition
    ...
    member definition
    }

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

    variable_name := structure_variable_type {value1, value2...valuen}
    或
    variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}

如果要**访问结构体成员**，需要使用点号 . 操作符，格式为：`结构体.成员名` 。
结构体类型变量使用 struct 关键字定义。

可以像其他数据类型一样将结构体类型作为参数传递给函数。

可以定义指向结构体的指针，类似于其他指针变量，格式如下：

`var struct_pointer *Books
`

以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：

`struct_pointer = &Book1
`

使用结构体指针访问结构体成员，使用 "." 操作符：

`struct_pointer.title`


## Go 16 - 切片 Slice

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

**定义切片**

    // 切片不需要说明长度。
    var identifier []type

    // 或使用 make() 函数来创建切片:
    var slice1 []type = make([]type, len)
    
    // 也可以简写为
    slice1 := make([]type, len)
    
    // 也可以指定容量，其中 capacity 为可选参数。
    make([]T, length, capacity)
    这里 len 是数组的长度并且也是切片的初始长度。

**切片初始化**

    s :=[] int {1,2,3 }
    直接初始化切片，[] 表示是切片类型，{1,2,3} 初始化值依次是 1,2,3，其 cap=len=3。
    
    s := arr[:]
    初始化切片 s，是数组 arr 的引用。
    
    s := arr[startIndex:endIndex]
    将 arr 中从下标 startIndex 到 endIndex-1 下的元素创建为一个新的切片。
    
    s := arr[startIndex:]
    默认 endIndex 时将表示一直到arr的最后一个元素。
    
    s := arr[:endIndex]
    默认 startIndex 时将表示从 arr 的第一个元素开始。
    
    s1 := s[startIndex:endIndex]
    通过切片 s 初始化切片 s1。
    
    s :=make([]int,len,cap)
    通过内置函数 make() 初始化切片s，[]int 标识为其元素类型为 int 的切片。

**len()和cap()**：切片是可索引的，并且可以由 len() 方法获取**长度**。切片提供了计算容量的方法 cap() 可以测量切片**最长可以达到多少**。

    var numbers = make([]int,3,5)
    
    // len(numbers) = 3
    // cap(numbers) = 5
    // numbers = [0 0 0]

一个切片在未初始化之前默认为 nil，长度为 0

    var numbers []int
    
    // len(numbers) = 0
    // cap(numbers) = 0
    // numbers == nil

**append()和copy()**：如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。


## Go 17 - 范围 Range

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

eg: 

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // sum: 9


## Go 18 - 集合 Map

Map 是一种**无序的键值对的集合**。Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，我们无法决定它的返回顺序，这是因为 Map 是使用 hash 表来实现的。

可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

    /* 声明变量，默认 map 是 nil */
    var map_variable map[key_data_type]value_data_type
    
    /* 使用 make 函数 */
    map_variable := make(map[key_data_type]value_data_type)
    
    // 如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对

Eg: 

    var countryCapitalMap map[string]string /*创建集合 */
    countryCapitalMap = make(map[string]string)

delete() 函数用于删除集合的元素, 参数为 map 和其对应的 key。

    delete(countryCapitalMap, "France")


## Go 19 - 递归 Recursion

运行的过程中调用自己。Go 语言支持递归。但在使用递归时需要设置退出条件，否则递归将陷入无限循环中。

递归函数对于解决数学上的问题是非常有用的，就像计算阶乘，生成斐波那契数列等。

Eg: 

    func recursion() {
      recursion() /* 函数调用自身 */
    }
    
    func main() {
      recursion()
    }


## Go 20 - 类型转换 Type expression

类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：

    type_name(expression)
    // type_name 为类型，expression 为表达式。


## Go 21 - 接口 Interface

Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

定义接口：

    type interface_name interface {
        method_name1 [return_type]
        method_name2 [return_type]
        method_name3 [return_type]
        ...
        method_namen [return_type]
    }

定义结构体：

    type struct_name struct {
        /* variables */
    }

实现接口方法：

    func (struct_name_variable struct_name) method_name1() [return_type] {
        /* 方法实现 */
    }
    ...
    func (struct_name_variable struct_name) method_namen() [return_type] {
        /* 方法实现*/
    }
