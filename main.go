package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

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

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Success: true,
			Message: "wellcome to the backend",
			Result:  ListUser,
		})
	})

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

	r.GET("users/:id", func(ctx *gin.Context) {
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

	r.Run("localhost:8888")

}
