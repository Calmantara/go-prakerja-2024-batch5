// syarat untuk membuat program golang
// 1. package main
// 2. function main

package main

// [OPTIONAL]
// package fmt -> package built in
// untuk manipulasi input outut
import "fmt"

func main() {
	// Intepreter vs Compiler
	fmt.Println("Hello World")

	// I. Variables
	// memiliki konvensi variable:
	// 1. menggunakan camelCase
	// 2. golang memiliki exported and unexported variable

	// cara mendeklarasi variable
	var name string // sering digunakan
	var number int = 100

	var firstName = "calman"
	fmt.Println(firstName)

	name = "user"
	email := "user@mail.com" // sering digunakan
	// fmt.Println(name, number)

	fmt.Printf("variable name, tipe:%T value:%v\n", name, name)
	fmt.Printf("variable email, tipe:%T value:%v\n", email, email)
	fmt.Printf("variable number, tipe:%T value:%v\n", number, number)

	// fmt.Println vs fmt.Printf
	// println => tidak bisa mengandung verb, auto enter
	// printf => bisa mengandung verb / format (%T, %v, %s, \n)

	// Multi declare variables
	username, age := "username1", 19
	fmt.Println(username, age)
	// function => bisa memberikan output lebih dari 1
	_ = "some value"
	// underscore variable => deklarasi variable yang bisa kita ignore

	// II. DATA TYPE
	// golang => merupakan bahasa pemrograman yang strict dengan data type
	// 1 + "1" = "11"
	// ada apa aja data type di go:
	// 1. integer (int) basic
	// 2. float (float32, float64) basic
	// 3. string basic
	// 4. boolean basic
	// 5. usigned integer (uint) basic
	// 6. byte basic alias
	// 7. interface
	// 8. nil -> value
	// 9. map, slice, pointer, function, channel

	// varchar
	// object

	// membaca:
	//	a. mutable
	//	b. immutable

	// int8, int16, int32, int64, float32, float64
	// panjang bit (banyak memory yang diperlukan)

	// kita mau menyimpan data umur
	// umur paling panjang 200th
	// int8
	var age8 uint8 = 200
	fmt.Printf("age8, tipe:%T value:%v\n", age8, age8)
	// string: word
	// bool => true atau false
	bool1 := true
	bool2 := false
	fmt.Println(bool1, bool2)
	// biasanya digunakan untuk flag, conditional statement
	var b1 byte
	// data type alias (byte => uint8, rune => uint32)
	b1 = 10
	fmt.Println(b1)

	// type dynamic
	// abstraksi dari semua data type
	var iface1 interface{} // any alias dari interface{}
	iface1 = 100
	fmt.Printf("iface1=> tipe:%T value:%v\n", iface1, iface1)
	iface1 = "some string"
	fmt.Printf("iface1=> tipe:%T value:%v\n", iface1, iface1)

	// nil => bukan string kosong, bukan 0
	iface1 = nil
	fmt.Printf("iface1=> tipe:%T value:%v\n", iface1, iface1)
	var iface2 interface{}
	fmt.Printf("iface2 => type:%T value:%v", iface2, iface2)
	// akan dibahas lebih lanjut di bahasan POINTER

	// varchar || text === string
	// database type untuk representasi word

	// object => dictionary di python / object di Javascript / class di Java
	// object => map / struct
	// Golang sendiri itu bukan OOP (object oriented programming)
}
