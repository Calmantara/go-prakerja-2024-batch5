package main

import (
	"fmt"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi3/order"
)

type Class struct {
}

type User struct {
	FirstName string
	LastName  string
	Email     string
	Age       uint
	Class     Class
}

func main() {
	// IV. Exported Unexported Package
	odr := order.CreateOrder()
	fmt.Println(odr.Merchant, odr.Number, odr.GetState())

	odr.Paid()
	fmt.Println(
		odr.Merchant,
		odr.Number,
		odr.GetState(),
		order.OrderMerchant)
}

func methodSession() {
	// III. Method
	// apa bedanya method dan function?

	// suatu struct, selain bisa memiliki property
	// struct juga bisa memiliki function (method)

	// Product:
	// ey kita harus bisa menuliskan
	// nama lengkap dari user: FirstName + LastName
	usr2 := User{
		FirstName: "Python",
		LastName:  "Prakerja",
		Email:     "python.prakerja@mail.com",
		Age:       1,
	}
	fmt.Println(name(usr2))
	fmt.Println(usr2.Name())

	fmt.Println(usr2.IncreaseAge())
	fmt.Println(usr2.Age)
}

func (u *User) IncreaseAge() uint {
	u.Age += 1
	return u.Age
}

func (u User) Name() string {
	// METHOD DARI USER
	return u.FirstName + " " + u.LastName
}

func name(in User) string {
	return in.FirstName + " " + in.LastName
}

func structSession() {
	// II. Struct

	var usr1 User
	usr1.FirstName = "Golang"
	usr1.LastName = "Prakerja"
	usr1.Email = "golang.prakerja@mail.com"
	usr1.Age = 1
	fmt.Printf("%+v \n", usr1)

	usr2 := User{
		FirstName: "Python",
		LastName:  "Prakerja",
		Email:     "python.prakerja@mail.com",
		Age:       1,
	}
	fmt.Printf("%+v \n", usr2)

	// Qn: apakah User bisa kita buat
	// sebagai array / slice?
	// 1. bisa menjadi array / slice
	// 2. dia bisa juga menjadi input / output dari suatu function
	users := []User{
		{
			FirstName: "Golang",
			LastName:  "Prakerja",
			Email:     "golang.prakerja@mail.com",
			Age:       10,
		},
		{
			FirstName: "Golang2",
			LastName:  "Prakerja2",
			Email:     "golang2.prakerja2@mail.com",
			Age:       20,
		},
	}
	users = append(users, User{
		FirstName: "Golang3",
		LastName:  "Prakerja3",
		Email:     "golang3.prakerja3@mail.com",
		Age:       30,
	})
	for _, user := range users {
		fmt.Println(user)
	}
	// pointer
	num := 1
	var num2 *int
	fmt.Println(num, num2)
	// assign value in num2
	num2 = &num
	fmt.Println(num, num2)
	// change value in num2
	*num2 = 10 // mengganti nilai dari pointer
	fmt.Println(num, num2)

	var usrPtr *User
	usrPtr = &usr1
	fmt.Println(*usrPtr)
	// cara lain untuk mendefinisikan
	// pointer of struct
	userPtr2 := &User{
		FirstName: "",
		LastName:  "",
		Email:     "",
		Age:       1,
	}
	fmt.Println(*userPtr2)
	// anonymous struct
	student := struct {
		FirstName string
		LastName  string
		Email     string
		Age       uint
	}{
		FirstName: "calman",
		LastName:  "tara",
	}
	fmt.Println(student)

	//  Product:
	// tipe dari user:
	// student
	// teacher
	// tapi kita tetep mengkategorikan sebagai User
	// Embedded Struct: membuat struct dari turunan struct lainnya
	type Student struct {
		NIM uint64
		User
	}
	// student adalah turunan dari User
	// tapi memiliki property sendiri juga
	// NIM
	var std Student
	std.FirstName = "std1"
	std.LastName = "std1"
	std.Email = "std1@mail.com"
	std.Age = 1
	std.NIM = 1000
	fmt.Println(std)

}

func functionSession() {
	// I. Function
	in := []int{1, 2, 3, 4, 5}
	out, out2, out3 := sum(in, in, in)
	out, _, _ = sum(in, in, in)
	fmt.Println(out, out2, out3)
	// Qn: bisa berapa banyak
	// input variable yang kita bisa masukkan?
	//	-> bisa sebanyak banyaknya

	// Qn: bisa berapa banyak
	// output dari suatu function?
	//	-> bisa sebanyak banyaknya

	// NB:
	// multi input dan output dari suatu function
	// bisa dipermudah dengan menggunakan struct

	// Product:
	// buat suatu feature sum
	// tidak bisa memastikan inputnya itu ada berapa
	// dan user bisa memberikan input sebanyak banyaknya
	outSU := sumUpgraded(1, 2, 3, 4, 5, 6, 7, 8, 9)
	outSU2 := sumUpgraded2([]int{1, 2, 3, 4})
	fmt.Println(outSU, outSU2)

	in = append(in, 1, 2, 3, 4)
	sumUpgraded(1, 2, 3, 4) // input optional (bisa kosong) dan bisa dimasukkan int
	// sumUpgraded2()          // input harus slice of integer

	// DESIGN PATTERN OPTION PARAMETER IN GOLANG

	// bisa ga menggunakan sumUpgraded
	// tapi inputnya adalah slice of integer??
	sumUpgraded(in...)

	// Qn: data type apa saja yang bisa dijadikan
	// input atau output dari suatu function?
	// semua tipe data
	// function bisa menerima function lainnya
	// sebagai input atau output:
	// 1. membuat alias dari suatu function
	// var fn1 callbackFn
	fn(func() string {
		return "prakerja batch 5"
	})

	oFN := fn2("calman")
	fmt.Println(oFN, oFN())
}

// alias
type callbackFn func() string

func fn(f callbackFn) {
	fmt.Println("hello:", f())
}

func fn2(str string) callbackFn {
	return func() string {
		return str
	}
}

func sumUpgraded(in ...int) (res int) {
	for _, val := range in {
		res += val
	}
	return
}

func sumUpgraded2(in []int) (res int) {
	for _, val := range in {
		res += val
	}
	return
}

func sum(slc, slc2, slc3 []int) (res1, res2, res3 int) {
	//	format:
	//
	// func -> menunjukkan bahwa type tsb adalah function
	// sum -> function name
	// slc -> input variable
	// []int -> input type
	// int -> output from function

	// Qn: bisa tidak output dari suatu
	// function, kita predefined
	// (int, int, int) => kita hanya memberitahu type dari outputs
	// (res1 int, res2 int, res3 int) => predefined output variable
	// (res1, res2, res3 int) => predefined output with same type

	for _, val := range slc {
		res1 += val
	}

	for _, val := range slc2 {
		res2 += val
	}

	for _, val := range slc3 {
		res3 += val
	}
	// return res1, res2, res3
	return
}
