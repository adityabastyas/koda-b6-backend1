package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "koda-b6-backend1/docs"
)

// @title User API
// @version 1.0
// @description belajar swagger
// @host localhost:8888
// @BasePath /
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  any    `json:"result"`
}

type Users struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var ListUser []Users

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// root
	// @Summary root
	// @Tags Root
	// @Success 200 {object} Response
	// @Router /  [get]
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
		})
	})

	// get all users
	// @Summary get all users
	// @Tags Users
	// @Success 200 {object} Response
	// @Router /users [get]
	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
			Result:  ListUser,
		})
	})

	// create user
	// @Summary create user
	// @Tags User
	// @Accept json
	// @Produce json
	// @Param body body Users true "user data"
	// @Success 200 {object} Response
	// @Router /users [post]
	r.POST("/users", func(ctx *gin.Context) {
		data := Users{}

		err := ctx.ShouldBindJSON(&data)

		if err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "create user failed",
			})
			return
		}

		for x := range ListUser {
			if ListUser[x].Email == data.Email {
				ctx.JSON(400, Response{
					Success: false,
					Message: "email is ready",
				})
				return
			}
		}

		ListUser = append(ListUser, data)

		ctx.JSON(200, Response{
			Success: true,
			Message: "user created",
		})

	})

	// get user by id
	// @Summary get user by id
	// @Tags Users
	// @Param id path string true "user id"
	// @Success 200 {object} Response
	// @Router /users/{id} [get]
	r.GET("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "9" {
			ctx.JSON(200, Response{
				Success: true,
				Message: "welcome to the admind",
			})
		} else {
			ctx.JSON(400, Response{
				Success: false,
				Message: fmt.Sprintf("id kamu adalah %s", id),
			})
		}
	})

	// update user
	// @Summary update user
	// @Tags Users
	// @Accept json
	// @Param id path int true "user index"
	// @Param body body Users true "user data"
	// @Success 200 {object} Response
	// @Router /users/{id} [patch]
	r.PATCH("/users/:id", func(ctx *gin.Context) {
		i, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || i < 0 || i >= len(ListUser) {
			ctx.JSON(404, Response{
				Success: false,
				Message: "user not fonud",
			})
			return
		}

		data := Users{}
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{
				Success: false,
				Message: "invalid body",
			})
			return
		}

		ListUser[i] = data
		ctx.JSON(200, Response{
			Success: true,
			Message: "update succes",
		})
	})

	// delete user
	// @Summary delete user
	// @Tags Users
	// @Param id path int true "user index"
	// @Success 200 {object} Response
	// @Router /users/{id} [delete]
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		i, err := strconv.Atoi(ctx.Param("id"))
		if err != nil || i < 0 || i >= len(ListUser) {
			ctx.JSON(400, Response{
				Success: false,
				Message: "user not found",
			})
			return
		}

		ListUser = append(ListUser[:i], ListUser[i+1:]...)

		ctx.JSON(200, Response{
			Success: true,
			Message: "user deleted",
		})
	})

	/// register
	// @Summary register user
	// @Tags Auth
	// @Accept json
	// @Param body body Users true "register data"
	// @Success 200 {object} Response
	// @Router /register [post]
	r.POST("/register", func(ctx *gin.Context) {
		data := Users{}

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{false, "invalid body", nil})
			return
		}

		if strings.TrimSpace(data.Email) == "" {
			ctx.JSON(400, Response{false, "email wajib diisi", nil})
			return
		}

		if strings.TrimSpace(data.Password) == "" {
			ctx.JSON(400, Response{false, "password wajib diisi", nil})
			return
		}

		for _, u := range ListUser {
			if u.Email == data.Email {
				ctx.JSON(400, Response{false, "email sudah terdaftar", nil})
				return
			}
		}

		ListUser = append(ListUser, data)

		ctx.JSON(200, Response{true, "register success", nil})

	})

	// login
	// @Summary login user
	// @Tags Auth
	// @Accept json
	// @Param body body Users true "login data"
	// @Success 200 {object} Response
	// @Router /login [post]
	r.POST("/login", func(ctx *gin.Context) {
		data := Users{}

		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(400, Response{false, "invalid body", nil})
			return
		}

		if data.Email == "" || data.Password == "" {
			ctx.JSON(400, Response{false, "email & password wajib", nil})
			return
		}

		for _, u := range ListUser {
			if u.Email == data.Email && u.Password == data.Password {
				ctx.JSON(200, Response{true, "login success", u})
				return
			}
		}

		ctx.JSON(401, Response{false, "email / password salah", nil})
	})

	r.Run("localhost:8888")

}
