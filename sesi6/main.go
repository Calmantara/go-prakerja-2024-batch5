package main

import (
	"net/http"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/handler"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/model"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/repository"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PRODUCT:
// membutuhkan system yang bisa record
// data siswa:
// 	- ID
// 	- Name
// 	- Email
// 	- Gender
//	- DoB (Date of Birth)

// Technical Spec:
//   - Domain => Students
//	 - GET /students
//	 - POST /students
//	 - PUT /students/:id
//	 - DELETE /students/:id

// ketika membuat suatu aplikasi / web server
// tetaplah menganut concept domain driven design (DDD)
func main() {
	// membuat HTTP Server / Web Server dengan menggunakan net/http
	// kita akan membuat web server dengan menggunakan Gin Gonic

	// cara menggunakan gin:
	// install gin di go mod: go get -u github.com/gin-gonic/gin
	// import di main
	ge := gin.New()

	ge.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,
			map[string]any{
				"status": "OK!",
			})
	})

	// Step by Step:
	// 1. membuat model struct yang bisa digunakan di GORM
	// 2. connect ke database dengan menggunakan GORM
	// 3. membuat migrasi DDL (kalau kita lakukan dengan manual) (CI/CD)
	// 4. membuat logic code pada repository layer

	// connect to gorm database
	// https://gorm.io/docs/connecting_to_the_database.html
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=35432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// melakukan migrasi / DDL
	// membuat table user
	err = db.AutoMigrate(&model.Student{})
	if err != nil {
		panic(err)
	}
	// dependencies
	userLocalRepo := &repository.StudentLocalRepo{}
	userPgRepo := &repository.StudentPgRepo{DB: db}
	userService := &service.UserService{UserLocalRepo: userLocalRepo, UserPgRepo: userPgRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	apiV1 := ge.Group("/api/v1")
	// /students [POST, GET]
	studentGroup := apiV1.Group("/students")
	studentGroup.GET("", userHandler.Get)
	studentGroup.POST("", userHandler.Create)

	// /students/:id [PUT, DELETE]
	studentGroup.PUT("/:id", userHandler.Update)
	if err := ge.Run(":8080"); err != nil {
		panic(err)
	}
}
