package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"lovelake.cn/app/pkg/handlers"
	"lovelake.cn/app/pkg/middlewares"
	"lovelake.cn/app/pkg/models"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
	shutdowns []func() error
)

func init() {
	// load environment variable
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup logrus
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}

	log.SetLevel(logLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	var (
		host   = os.Getenv("HTTP_HOST")
		port   = os.Getenv("HTTP_PORT")
		server = gin.New()
	)

	// 初始化数据库
	// 配置 MySQL 连接信息
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	// 连接到 MySQL 数据库
	conn := mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	})

	db, err := gorm.Open(conn, &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// 进行数据库迁移，创建 User 表
	db.AutoMigrate(&models.User{})

	// 使用中间件将数据库连接注入到每个请求的上下文中
	server.Use(func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	})

	// 从os 的env中，获取HTTP_SECRET，默认值是123
	httpSecret := os.Getenv("HTTP_SECRET")
	if httpSecret == "" {
		httpSecret = "Fu7MAh7UntEZcoNTWUEIEpunGOZIkmTk"
	}

	store := cookie.NewStore([]byte(httpSecret))

	// 设置session中间件，参数mysession，指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	server.Use(sessions.Sessions("sid", store))

	server.Use(gin.Recovery())
	server.Use(middlewares.Logging())

	// 路由注册在这里
	// 注册用户模块
	authGroup := server.Group("/api/v1/users")
	authGroup.GET("/profile", handlers.Profile)
	authGroup.POST("/login", handlers.Login)
	authGroup.POST("/registry", handlers.Registry)

	// 第1个group
	group1 := server.Group("/api/v1/module1")
	group1.GET("/ping/", handlers.Ping)

	// 第2个group
	group2 := server.Group("/api/v1/module2")
	group2.GET("/ip/", handlers.ClientIP)

	err = server.Run(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatal(err)
	}

}

func gracefulShutdown(ctx context.Context, server *http.Server, shutdown chan struct{}) {
	var (
		sigint = make(chan os.Signal, 1)
	)

	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
	<-sigint

	logger.Info("shutting down server gracefully")

	// stop receiving any request.
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("shutdown error", zap.Error(err))
	}

	// close any other modules.
	for i := range shutdowns {
		shutdowns[i]()
	}

	close(shutdown)
}
