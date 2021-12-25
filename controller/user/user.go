package user

import (
	"c_blog/entity"
	"c_blog/global"
	_const "c_blog/global/const"
	"c_blog/global/source"
	"c_blog/global/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello1(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}

func Login(c *gin.Context) {

	var user entity.User
	c.ShouldBind(&user)
	//fmt.Printf("%v", user)
	rowsAffected := source.DB.Where("user_name = ? AND user_pwd = ?", user.Username, user.Password).Find(&user).RowsAffected
	fmt.Println(rowsAffected)
	if rowsAffected == 0 {
		c.JSON(http.StatusOK, global.ResultVO{0, "用户名或密码不正确", nil})
		return
	}

	if rowsAffected == 1 {
		c.SetCookie("user_id", util.MD5(user.Userid), _const.COOKIE_MAX_AGE, "/", "localhost", false, false)
		c.JSON(http.StatusOK, global.ResultVO{1, "登录成功", nil})
		return
	}
	c.JSON(http.StatusInternalServerError, global.ResultVO{0, "服务器异常，请联系管理员", nil})
}

/*
第一个参数 key
第二个参数 value
第三个参数 过期时间.如果只想设置 Cookie 的保存路径而不想设置存活时间，可以在第三个 参数中传递 nil
第四个参数 cookie 的路径
第五个参数 cookie 的路径 Domain 作用域 本地调试配置成 localhost , 正式上线配置成域名
第六个参数是 secure ，当 secure 值为 true 时，cookie 在 HTTP 中是无效，在 HTTPS 中 才有效
第七个参数 httpOnly，是微软对 COOKIE 做的扩展。如果在 COOKIE 中设置了“httpOnly”属性， 则通过程序（JS 脚本、applet 等）将无法读取到 COOKIE 信息，防止 XSS 攻击产生
*/

func Register(c *gin.Context) {
	println(c.PostForm("username"))
	println(c.PostForm("password"))
	//重名检查
	user1 := entity.User{Username: c.PostForm("username")}
	result := source.DB.Find(&user1)
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
	//判断当前登录用户，获取用户id的加密后的值,暂不做处理，直接存储
	/*
		cookieValue, err := c.Cookie("userid")
		if err != nil {
			panic(err)
		}
		var user entity.User
		c.ShouldBind(&user)
		userId := user.Userid
		if cookieValue == userId {

		} else {

		}
	*/
	var user entity.User
	c.ShouldBind(&user)
	user.Userid = "1111"
	source.DB.Where("user_id = ?", user.Userid).Updates(&user)

}
