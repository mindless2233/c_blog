package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type result struct {
	code int8
	msg  string
}

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
func Login(c *gin.Context) {
	//types := c.DefaultPostForm("type", "post")
	//username := c.PostForm("username")
	//password := c.PostForm("userpassword")
	//c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	//result2 := result{1, "ok"}
	allUsers := []user{{ID: 123, Name: "张三", Age: 20}, {ID: 456, Name: "李四", Age: 25}}
	c.IndentedJSON(200, allUsers)
	//c.JSON(http.StatusOK, gin.H{"res": result{1, "ok"}})
}
