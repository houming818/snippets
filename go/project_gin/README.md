# go-gin 脚手架说明

## 配置文件

`etc/env.txt`是配置文件，按照要求填写

## 入口目录 `cmd/main.go`

main.go 中做路由注册,详细见代码注册

## 代码目录 `pkg/*'


## 启用的基本中间件

gorm已配置，获取方式

```
	db := ctx.MustGet("DB").(*gorm.DB)
	var user models.User
	db.Where("username = ?", sessionUser).First(&user)
```

session已配置,获取方式

```
	session := sessions.Default(ctx)
	sessionUser := session.Get("user")
```

