package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/handler"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/middleware"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/model"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/repository"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/service"
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
//   - Domain => Users
//	 - GET /users
//	 - POST /users
//	 - PUT /users/:id
//	 - DELETE /users/:id

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
	pgHost := os.Getenv("PG_HOST")
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDB := os.Getenv("PG_DB")
	pgPort := os.Getenv("PG_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		pgHost,
		pgUser,
		pgPassword,
		pgDB,
		pgPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// melakukan migrasi / DDL
	// membuat table user
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
	// dependencies
	userLocalRepo := &repository.UserLocalRepo{}
	userPgRepo := &repository.UserPgRepo{DB: db}
	userService := &service.UserService{UserLocalRepo: userLocalRepo, UserPgRepo: userPgRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	// skenario middleware:
	// 1. setiap /api/v1 harus melewati middleware 1
	// 2. setiap yang memiliki /users harus melewati middleware 2
	// 3. setial yang memiliki /users/:id harus melewati middleware 3

	apiV1 := ge.Group("/api/v1")
	// /users [POST, GET]
	userGroup := apiV1.Group("/users")
	// action
	userGroup.POST("/login", userHandler.Login) // midware1 midware2 login

	// PRODUCT:
	// untuk mendapatkan data users
	// untuk create users
	// untuk edit users
	// SEMUANYA HARUS LOGIN DULU!

	// harus menambahkan MIDDLEWARE VALIDATE JWT
	// domain
	userGroup.POST("", userHandler.Create) // midware1 midware2 create
	userGroup.Use(middleware.BearerAuthorization())
	userGroup.GET("", userHandler.Get) // midware1 midware2 get

	// /users/:id [PUT, DELETE]
	// PRODUCT:
	// yang bisa mengedit user hanyalah user yang bersesuaian
	// id: 1 email:student@edu.com
	// /users/1

	// Step:
	// 1. validate JWT
	// 2. get email from JWT
	// 		a. bisa dimasukkan ke context, baru diquery di handler
	//		b. bisa langsung query user by email di Middleware
	// 3. get id from param
	// 4. compare id from param and user id

	servicePort := os.Getenv("PORT")
	if servicePort == "" {
		servicePort = "8080"
	}

	userGroup.PUT("/:id", userHandler.Update) // midware1 midware2 midware3 update
	if err := ge.Run(":" + servicePort); err != nil {
		panic(err)
	}
}

func Add(a, b int) int {
	// kita sudah ada test cases
	// sekarang bagaimana caranya
	// kita coding sehingga
	// semua test cases SUCCESS
	return a + b
}
