package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/helper"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/model"
	"github.com/Calmantara/go-prakerja-2024-batch5/sesi7/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService service.UserServiceInterface
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
	studentCreate := model.UserCreate{}
	if err := ctx.Bind(&studentCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
		return
	}
	hashedPassword, err := helper.HashPassword(studentCreate.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	// call service
	err = u.UserService.Create(&model.User{
		Email:    studentCreate.Email,
		Password: hashedPassword,
		Name:     studentCreate.Name,
		DoB:      studentCreate.DoB,
		Gender:   studentCreate.Gender,
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
		return
	}
	id, _ := strconv.Atoi(idStr)
	// binding payload
	studentUpdate := model.UserUpdate{}
	if err := ctx.Bind(&studentUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
		return
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

func (u *UserHandler) Login(ctx *gin.Context) {
	// binding payload
	payload := &model.UserLogin{}
	if err := ctx.Bind(&payload); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.Response{
			Message: "bad request param",
			Success: false,
		})
		return
	}
	// fetch user by email
	user, err := u.UserService.GetByEmail(payload.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.Response{
			Message: "something went wrong",
			Success: false,
		})
		return
	}
	if user.ID <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, model.Response{
			Message: "user not found",
			Success: false,
		})
		return
	}
	// compare password
	isMatched := helper.CheckPasswordHash(payload.Password, user.Password)
	if !isMatched {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, model.Response{
			Message: "invalid email or password",
			Success: false,
		})
		return
	}
	// generate TOKEN if password is correct
	authToken, _ := helper.GenerateUserJWT(user.Name, user.Email, 2*time.Hour)
	sessionToken, _ := helper.GenerateUserJWT(user.Name, user.Email, 48*time.Hour)
	// return
	ctx.JSON(http.StatusOK, model.Response{
		Message: "logged in",
		Success: true,
		Data: model.Token{
			AuthToken:    authToken,
			SessionToken: sessionToken,
		},
	})
}
