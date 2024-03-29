package main

import (
	"c_blog/controller/user"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	v1 := r.Group("user")
	{
		v1.POST("/register", user.Register)
		v1.POST("/login", user.Login)
		v1.POST("/edit", user.EditUserInfo)
	}

	r.Run(":8082")
}
