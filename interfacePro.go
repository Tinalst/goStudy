package main

import "fmt"

// Human 1 - 定义Human结构体
type Human struct {
	name string
	age int
	phone string
}

// SayHi Human 实现SayHi方法
func (h Human) SayHi()  {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Sing Human 对象实现 Sing 方法
func(h Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

// Guzzle Human 对象实现 Guzzle 方法
func (h Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}


// Student 2 - 定义Student结构体
type Student struct {
	Human // 匿名字段
	school string
	loan float32
}

// BorrowMoney Student 实现 BorrowMoney 方法
func (s Student) BorrowMoney(amount float32) {
	s.loan += amount // (again and again and...)
}



// Employee 3 - 定义Employee结构体
type Employee struct {
	Human // 匿名字段
	company string
	money float32
}

// SpendSalary Employee 实现 SpendSalary 方法
func (e Employee) SpendSalary(amount float32) {
	e.money -= amount // More vodka please!!! Get me through the day!
}

// SayHi 重载SayHi
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) // 此句可以分成多行
}

// Men 定义 interface
// 被Human、Student 和 Employee 实现
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(song string)
	BorrowMoney(amount float32)
}

type ElderlyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

func main()  {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}

	// 定义Men变量，因为Men被Human、Student 和 Employee 实现
	// 所以i可以被Human、Student 和 Employee 类型变量赋值
	var i Men

	i = mike

	// i 能存储 Student
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	// i 也能存储 Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// 定义了 slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	// 这三个都是不同类型的元素，但是他们实现了 interface 同一个接口
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}

}