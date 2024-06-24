package main

import (
	"fmt"
	"time"
)

func main() {
	// III. Array
	arr := [3]int{}
	arr[0] = 10
	// arr = append(arr, 1) tidak akan berlaku untuk array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	// Qn: apa saja tipe data yang bisa kita jadikan array/slice?
	// semua bisa (int, string, bool, interface)
	// array bisa kita buat dari array lainnya
	arrMultiDimension := [3][3]int{}
	// loop over array
	for idx, val := range arr {
		fmt.Printf("id: %v val: %v\n", idx, val)
	}
	// bagaimana caranya kita melakukan
	// looping untuk multi dimension array
	// I. solusi 1
	for i := 0; i < len(arrMultiDimension); i++ {
		for j := 0; j < len(arrMultiDimension[i]); j++ {
			fmt.Println(arrMultiDimension[i][j])
		}
	}

	for outId, outVal := range arrMultiDimension {
		for inId, inVal := range outVal {
			fmt.Println(outId, outVal, inId, inVal)
		}
	}

	arr2 := arr
	arr2[0] = 5
	arr2[1] = 100
	arr[2] = 99

	// KITA BAHAS SETELAH SLICE
	// Qn: apa output dari arr2? 5, 100, 0
	// apa output dari arr? 10, 0, 99

	// IV. Slice
	// len dan cap
	slc := []int{}
	slc = append(slc, 1, 2, 3, 4, 5, 6, 7)
	slc[0] = 100
	for i := 0; i < len(slc); i++ {
		fmt.Println(slc[i])
	}
	// aku mau memasukkan semua element dari arr ke slc
	// a. Solusi I
	for _, val := range arr {
		slc = append(slc, val)
	}
	fmt.Println(slc)
	// b. Solusi II
	slc = append(slc, arr[0], arr[1], arr[2])

	slc2 := []int{1, 2, 3, 4, 5}
	// kita masukkan seluruh elm slc2 ke dalam slc
	slc = append(slc, slc2...)

	// slice => memotong isi dari slice
	slc3 := slc[0:2]
	fmt.Println(slc3)
	slc3[0] = 99

	slc4 := []int{1, 2, 3, 4}
	slc5 := make([]int, 10)
	// untuk menggunakan copy
	// destination slice harus memiliki element terlebih dahulu
	copy(slc5, slc4)
	slc6 := []int{}
	slc6 = append(slc6, slc4...)
	slc5[0] = 99
	slc5[1] = 88
	slc4[2] = 77
	slc5 = append(slc5, 10)
	fmt.Println(slc4, slc5, slc6)
}

func condition() {
	// I. Condition
	// requirement tambahan pemerintah:
	//  - jika umur dari pengguna > 50, kamu harus test sim ulang
	currentYear := time.Now().Year() // 2024
	birthYear := 1960
	age := currentYear - birthYear // 19
	if age < 17 {
		fmt.Println("kamu belum bisa membuat sim")
	} else if age > 50 {
		fmt.Println("kamu harus test ulang untuk membuat sim")
	} else {
		// jika age >= 17
		fmt.Println("kamu sudah bisa membuat sim")
	}

	// bisa menggunakan temporary variable
	birthYear = 1980
	if userAge := currentYear - birthYear; userAge < 17 {
		fmt.Println("kamu belum bisa membuat sim")
	} else if userAge > 50 {
		fmt.Println("kamu harus test ulang untuk membuat sim")
	} else if userAge > 50 {
		fmt.Println("kamu harus test ulang untuk membuat sim")
	} else if userAge > 50 {
		fmt.Println("kamu harus test ulang untuk membuat sim")
	} else if userAge > 50 {
		fmt.Println("kamu harus test ulang untuk membuat sim")
	} else if userAge > 50 {
		fmt.Println("kamu harus test ulang untuk membuat sim")
	} else if userAge > 50 {
		fmt.Println("kamu harus test ulang untuk membuat sim")
	} else {
		// jika age >= 17
		fmt.Println("kamu sudah bisa membuat sim")
	}

	// switch
	// a. switch exact value
	// requirement tambahan dari product
	// - yang dikatakan abg adalah ketika age == 20 dan age == 17
	// - ketika abg, berarti user tersebut juga sudah dewasa
	age = 20
	switch age {
	case 10:
		// apakah age == 10?
		fmt.Println("kamu masih bocil")
	case 20, 17:
		// apakah age == 20?
		fmt.Println("kamu abg")
		fallthrough
	case 30:
		fmt.Println("umur kamu 100")
	default:
		// apakah diluar dari 10 dan 20
		fmt.Println("kamu sudah dewasa")
	}
	// b. switch with condition
	age = 23
	switch {
	case age < 10:
		fmt.Println("kamu masih bocil")
	case age < 17 && age >= 10:
		fmt.Println("kamu sudah mulai abg")
	case age < 24 && age >= 17:
		fmt.Println("kamu abg")
	default:
		fmt.Println("kamu sudah dewasa")
	}
	// Qn: apakah bisa kita buat
	// if di dalam if
	// atau swict di dalam switch
	if age > 10 {
		if age < 100 {

		}
	}

	switch age {
	case 100:
		switch age {
		case 200:

		}
	}
}

func looping() {
	// II. Looping
	// => definisi dan kegunaan
	// requirement:
	// tulis angka dari 1 - 500
	// fmt.Println(1)
	// fmt.Println(2)
	// fmt.Println(3)
	// fmt.Println(4)
	// fmt.Println(5)

	// a.
	// initial variable;
	// condition when to stop;
	// condition yang akan dilakukan setiap pengulangan
	for i := 1; i <= 500; i++ {
		fmt.Println(i)
	}

	// b.
	i := 0
	for i < 10 {
		fmt.Printf("pengulangan ke:%v\n", i)
		i++
	}

	// c. infinit loop
	num := 0
	for {
		if num == 1000 {
			break
		}
		num++
	}
	fmt.Println("end of looping")
}
