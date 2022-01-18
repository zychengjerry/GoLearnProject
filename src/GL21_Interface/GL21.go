package main

import "fmt"

type Phone interface {
	call()
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("Hi I'm IPhone!")
}

type IPhoneX struct {
}

func (iPhoneX IPhoneX) call() {
	fmt.Println("Hi I'm IPhoneX!")
}

func main() {
	var phone Phone

	phone = new(IPhone)
	phone.call()

	phone = new(IPhoneX)
	phone.call()
}

/*
在上面的例子中，我们定义了一个接口Phone，
接口里面有一个方法call()。
然后我们在main函数里面定义了一个Phone类型变量，
并分别为之赋值为IPhone和IPhoneX。
然后调用call()方法
*/

/*
Hi I'm IPhone!
Hi I'm IPhoneX!
*/
