package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Student struct {
	NIM       uint64
	FirstName string
	LastName  string
}

var students = []Student{}

func main() {
	// membuat aplikasi server:
	// API WEB SERVER

	// GOLANG: support untuk membuat API
	// dengan menggunakan NATIVE PACKAGE (net/http)
	// dengan menggunakan WEB FRAMEWORK (gin gonic)

	port := "8080"
	version := "/api/v1"

	http.HandleFunc(version+"/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello From My Awesome WEB SERVER API")
	})
	// student api
	http.HandleFunc(version+"/students", func(w http.ResponseWriter, r *http.Request) {
		// METHOD GET
		if r.Method == http.MethodGet {
			// check apakah ada param ID / NIM
			path := strings.Split(r.URL.Path, "/students")
			if len(path) > 1 {
				if path[1] != "" {
					id := path[1]
					// search student with same nim
					for _, student := range students {
						if fmt.Sprintf("%v", student.NIM) == id {
							json.NewEncoder(w).Encode(student)
							return
						}
					}
					http.Error(w, "Student Not Found", http.StatusNotFound)
					return
				}
			}
			// return seluruh data student
			// dalam bentuk / format JSON
			json.NewEncoder(w).Encode(students)
			return
		}

		// METHOD POST
		if r.Method == http.MethodPost {
			// akan menambahkan data student
			// dari masukan API (form)
			nim := r.FormValue("nim")
			firstName := r.FormValue("first_name")
			lastName := r.FormValue("last_name")

			// error handling
			nimNumber, err := strconv.Atoi(nim)
			if err != nil {
				http.Error(w, "invalid nim", http.StatusBadRequest)
				return
			}
			// append to slice
			students = append(students, Student{
				NIM:       uint64(nimNumber),
				FirstName: firstName,
				LastName:  lastName,
			})
			fmt.Fprintf(w, "ok")
			return
		}
	})
	// student api with id
	http.HandleFunc(version+"/students/", func(w http.ResponseWriter, r *http.Request) {
		// METHOD GET
		if r.Method == http.MethodGet {
			// check apakah ada param ID / NIM
			path := strings.Split(r.URL.Path, "/students/")
			if len(path) > 1 {
				if path[1] != "" {
					id := path[1]
					// search student with same nim
					for _, student := range students {
						if fmt.Sprintf("%v", student.NIM) == id {
							json.NewEncoder(w).Encode(student)
							return
						}
					}
					http.Error(w, "Student Not Found", http.StatusNotFound)
					return
				}
			}
			// return seluruh data student
			// dalam bentuk / format JSON
			http.Error(w, "Student Not Found", http.StatusNotFound)
			return
		}
	})

	// html template
	// /students
	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		// SERVER SIDE RENDERING (SSR)
		tpl, err := template.ParseFiles("./student.html")
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
		err = tpl.Execute(w, students)
		if err != nil {
			http.Error(w, "something went wrong", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":"+port, nil)
} // end of program block

func ErrorSession() {
	// error handling, panic, recover
	// menghandle error tsb dengan baik

	strNum := "10"
	num, err := strconv.Atoi(strNum)
	// graceful error handling
	if err != nil {
		fmt.Println("something wrong with your input")
	} else {
		fmt.Println("got some input number", num)
	}

	// custom error
	// userInput harus >= 0
	// selain itu akan memunculkan error
	var err2 error
	userInput := -100
	if userInput < 0 {
		// err2 = errors.New("invalid user input, must be positif number")
		err2 = fmt.Errorf("invalid user input, must be positif number: %v", userInput)
	}
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	// biasanya error yang tidak dihandling dengan baik
	// atau kondisi yang tidak valid
	// akan menghasilkan panic di golang
	// panic => kondisi program akan berhenti secara paksa karena suatu kondisi (bukan os.Exit)

	// cara menghindari panic:
	// 1. coding dengan benar
	// 2. membuat unit test dengan baik
	// 3. mengahandle panic
	// 		a. recover
	defer func() {
		if rec := recover(); rec != nil {
			fmt.Println("got panic", rec)
		} else {
			fmt.Println("ok!")
		}
	}()
	var err3 error                                 // masih nil
	fmt.Println("value of error is", err3.Error()) // dipanggil method dari nil, sehingga panic
}

func DeferSession() {
	num := 100
	if num == 100 {
		return
	}
	defer fmt.Println("end of main block")
	defer fmt.Println("end of main block 2")
	defer fmt.Println("end of main block 3")
	// behavior dari defer:
	// last in first out
	otherFunction()
	for i := 0; i < 100; i++ {
		fmt.Println("long process", i)
		// if i == 64 {
		// 	return // defer akan terpanggil di sini
		// } // end of program block
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if i == 974 {
				return
			}
		}
		return
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			if i == 123 {
				return
			}
		}
		return
	}()
	wg.Wait()

	// exit
	// => akan keluar dari program
	os.Exit(0)
}

func otherFunction() {
	defer fmt.Println("defer from other function")
} // end of program block otherFunction
