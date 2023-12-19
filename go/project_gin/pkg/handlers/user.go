package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"lovelake.cn/app/pkg/models"
	"lovelake.cn/app/pkg/utils"
)

// Login 登录用户信息
func Login(ctx *gin.Context) {
	// 获取登录的username和password
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	// 对password进行sha256加密
	pwdSha256 := utils.HashBySha256(password, "")

	// 获取数据库连接
	db := ctx.MustGet("DB").(*gorm.DB)

	// 验证登录信息
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 处理记录不存在的情况
			ctx.JSON(200, gin.H{"code": 1, "msg": "用户不存在"})
			return
		} else {
			// 处理其他错误
			ctx.JSON(200, gin.H{"code": 1, "msg": "未知错误" + err.Error()})
			return
		}
	}
	if user.Password != pwdSha256 {
		ctx.JSON(200, gin.H{"code": 1, "msg": "密码错误"})
		return
	}

	session := sessions.Default(ctx)
	session.Set("user", username)
	session.Save()
	ctx.JSON(200, gin.H{"code": 0, "msg": "登录成功"})
	return

}

// Registry User 注册用户
func Registry(ctx *gin.Context) {
	// 获取用户名 密码 邮箱
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	retryPassword := ctx.PostForm("retryPassword")
	email := ctx.PostForm("email")

	if password != retryPassword {
		ctx.JSON(200, gin.H{"code": 1, "msg": "两次密码不一致"})
		return
	}

	// 对password进行sha256加密
	pwdSha256 := utils.HashBySha256(password, "")

	// 获取数据库连接
	db := ctx.MustGet("DB").(*gorm.DB)

	// 验证登录信息
	var user models.User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 处理记录不存在的情况
			// 创建用户
			user := models.User{Username: username, Password: pwdSha256, Email: email}
			db.Create(&user)
			ctx.JSON(200, gin.H{"code": 0, "msg": "注册成功"})
			return
		} else {
			// 处理其他错误
			ctx.JSON(200, gin.H{"code": 1, "msg": "未知错误" + err.Error()})
			return
		}
	} else {
		ctx.JSON(200, gin.H{"code": 1, "msg": "用户已存在"})
		return
	}
}

// Profile 登录用户信息
func Profile(ctx *gin.Context) {
	// 获取登录的username和password

	// 获取数据库连接
	db := ctx.MustGet("DB").(*gorm.DB)

	// 验证登录信息

	session := sessions.Default(ctx)
	sessionUser := session.Get("user")
	var user models.User
	db.Where("username = ?", sessionUser).First(&user)
	ctx.JSON(200, gin.H{"code": 0, "msg": user})
	return

}
