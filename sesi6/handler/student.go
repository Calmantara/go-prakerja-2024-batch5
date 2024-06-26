package handler

import (
	"net/http"
	"strconv"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/model"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi6/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func (u *UserHandler) Get(ctx *gin.Context) {
	users, err := u.UserService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	ctx.JSON(http.StatusOK, model.Response{
		Message: "users fetched",
		Success: true,
		Data:    users,
	})
}

func (u *UserHandler) Create(ctx *gin.Context) {
	// binding payload
	studentCreate := model.StudentCreate{}
	if err := ctx.Bind(&studentCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	// call service
	err := u.UserService.Create(&model.Student{
		Email:  studentCreate.Email,
		Name:   studentCreate.Name,
		DoB:    studentCreate.DoB,
		Gender: studentCreate.Gender,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	// response

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "users created",
		Success: true,
	})
}

func (u *UserHandler) Update(ctx *gin.Context) {
	// bind id from path param
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	id, _ := strconv.Atoi(idStr)
	// binding payload
	studentUpdate := model.StudentUpdate{}
	if err := ctx.Bind(&studentUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
	}
	// call service
	err := u.UserService.Update(uint64(id), &studentUpdate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	// response

	ctx.JSON(http.StatusCreated, model.Response{
		Message: "users updated",
		Success: true,
	})
}
