package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
)

type HttpResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func main() {
	// go routine
	go fmt.Println("hello from go routine")
	go PrintSeq(0, 10, 10)
	go PrintSeq(11, 20, 5)
	// print berapa banyak go routine yang sudah kita spawn
	fmt.Println("num of go routine", runtime.NumGoroutine())
	PrintSeq(21, 30, 30)

	// setiap kita menjalankan
	// go routine, apakah harus ada delay atau tidak?
	// tidak => selama kita bisa ensure MAIN GO ROUTINE belum selesai

	// kasih contoh, kapan go routine digunakan? => background process
	// async process / concurrent process
	// NOTE: concurrency design pattern

	// BATAS MAXIMUM GO ROUTINE
	// => CPU dan RAM 4KB / 2KB
	// go routine leaks

	// channel => caranya go routine berkomunikasi dengan go routine lainnya
	// HTTP SERVER (BACKEND) => Graceful Shutdown

	// pattern lain untuk memastikan
	// semua go routine telah selesai dijalankan,
	// sebelum main go routine selesai / end of program
	// WAIT GROUP
	fmt.Println("=======================")
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		PrintSeq(0, 10, 10)
		wg.Done()
	}()
	go func() {
		PrintSeq(11, 20, 7)
		wg.Done()
	}()
	go func() {
		PrintSeq(11, 20, 2)
		wg.Done()
	}()
	wg.Wait() // BLOCKING PROCESS
	time.Sleep(time.Second)

	// APA YANG TERJADI KALAU
	// KITA MEMBERIKAN WG.ADD(1)
	// TAPI YANG DIPANGGIL WG.DONE ada di 2 go routine?
	// dia hanya menunggu 1 go routine yang dijalankan
	// tapi hasil akhirnya akan negative call panic

	// APA YANG TERJADI KALAU
	// KITA MEMBERIKAN WG.ADD(4)
	// TAPI YANG DIPANGGIL WG.DONE ada di 2 go routine?
	// dia akan deadlock

	// fmt.Println("=======================")
	// wg2 := sync.WaitGroup{}
	// wg2.Add(4)
	// go func() {
	// 	PrintSeq(0, 10, 10)
	// 	wg2.Done()
	// }()
	// go func() {
	// 	PrintSeq(11, 20, 7)
	// 	wg2.Done()
	// }()
	// wg2.Wait() // BLOCKING PROCESS

	// arr := []int{1, 2, 3}
	// wg3 := sync.WaitGroup{}
	// wg3.Add(1)
	// for _, val := range arr {
	// 	wg3.Add(1)
	// 	go func(wg_ *sync.WaitGroup, num int) {
	// 		fmt.Println("num", num)
	// 		wg_.Done()
	// 	}(&wg3, val)
	// }
	// wg3.Wait() // BLOCKING PROCESS
}

func PrintSeq(start, end int, delayMs int) {
	for i := start; i < end; i++ {
		fmt.Println("print for seq:", i)
		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
}

func InterfaceSession() {
	// var iface interface{}
	// interface{} => dynamic data type
	// untuk go 1.18 ++
	// interface juga memiliki built in
	// alias variable => any
	var iface any
	fmt.Println(iface)
	// assign iface to int
	iface = 10 // interface{}
	fmt.Printf("iface int: %T %v\n", iface, iface)
	iface = "string"
	fmt.Printf("iface string: %T %v\n", iface, iface)

	// revert iface value
	iface = 10

	// casting type from interface
	// bagaimana caranya kita assign value dari interface
	// ke suatu variable yang tipenya bersesuaian
	str1, ok := iface.(string) // ok => menunjukkan bahwa value bisa di casting ke string atau tidak
	if !ok {
		fmt.Println("iface is not string")
	} else {
		fmt.Printf("str1, type:%T value:%v", str1, str1)
	}

	// NOTE: named interface => dependency injection

	// kita tau iface tadi memiliki value string
	// sehingga kita bisa yakin (casting string)

	// iface bisa tidak slice?
	slcIface := []any{1, "string", 1.4}
	fmt.Println(slcIface)
	// iface bisa tidak kita jadi map?
	mapIface := map[string]any{} // JSON
	mapIface["name"] = "user 1"
	mapIface["age"] = 10
	fmt.Println(mapIface)

	// NOTE: package reflect
	kindOfIface := reflect.TypeOf(iface).Kind()
	fmt.Println(kindOfIface)
}
