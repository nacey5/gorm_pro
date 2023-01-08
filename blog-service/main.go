package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm_pro/blog-service/global"
	"gorm_pro/blog-service/internal/model"
	"gorm_pro/blog-service/internal/routers"
	"gorm_pro/blog-service/pkg/logger"
	"gorm_pro/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

// @title 博客系统
// @version 1.0
// @description Go博客项目
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
	s.ListenAndServe()
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupTranslations()
	if err != nil {
		log.Fatalf("init.setupTranslations err: %v", err)
	}

}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

var trans ut.Translator

func setupTranslations() error {
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	v, ok := binding.Validator.Engine().(*val.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
	return nil
}
