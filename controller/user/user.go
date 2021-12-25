package user

import (
	"c_blog/entity"
	"c_blog/global"
	"c_blog/global/source"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello1(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func Login(c *gin.Context) {
	//var userTemp entity.User
	//c.ShouldBind(&userTemp)
	//fmt.Printf("%v", userTemp)
	if c.PostForm("username") != "" && c.PostForm("password") != "" {
		user := entity.User{Username: c.PostForm("username"), Password: c.PostForm("password")}
		rowsAffected := source.DB.Where("user_name = ? AND user_pwd = ?", user.Username, user.Password).Find(&user).RowsAffected
		fmt.Println(rowsAffected)
		if rowsAffected == 0 {
			c.JSON(http.StatusOK, global.ResultVO{0, "用户名或密码不正确", nil})
			return
		}

		if rowsAffected == 1 {
			c.JSON(http.StatusOK, global.ResultVO{1, "登录成功", nil})
			//TODO 设置cookie
			return
		}

	} else {
		c.JSON(http.StatusOK, global.ResultVO{0, "用户名或密码不能为空", nil})
		return
	}
	c.JSON(http.StatusInternalServerError, global.ResultVO{0, "服务器异常，请联系管理员", nil})
}

func Register(c *gin.Context) {
	println(c.PostForm("username"))
	println(c.PostForm("password"))
	//重名检查
	user1 := entity.User{Username: c.PostForm("username")}
	result := source.DB.Find(&user1)
	println("--发生错误:", result.Error)
	resultCount := result.RowsAffected

	if resultCount == 0 {
		user2 := entity.User{Username: c.PostForm("username"), Password: c.PostForm("password")}
		result := source.DB.Create(&user2).RowsAffected
		if result == 1 {
			c.JSON(http.StatusOK, global.ResultVO{1, "注册成功", user2})
		} else {
			c.JSON(http.StatusOK, global.ResultVO{0, "注册失败，请联系管理员解决", nil})
		}
	} else {
		c.JSON(http.StatusOK, global.ResultVO{0, "注册失败，用户名已经存在", nil})
	}
}

func EditUserInfo(c *gin.Context) {

	//判断当前登录用户，获取用户id
	cookie, err := c.Cookie("userid")
	if err != nil {
		panic(err)
	}
	user_id1 := c.PostForm("user_id")
	if cookie == user_id1 {
		//根据表单中ID，判断是否是本人修改
		//	source.DB := global.GetDBConn()

	} else {

	}

}
