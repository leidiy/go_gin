package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg      *ini.File
	RunMode  string
	HTTPPort int

	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PageSize  int
	JWTSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("load conf err :%v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("parse app section error %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt()
	JWTSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("parse server error %v",err)
	}

	HTTPPort = sec.Key("HTTP_PORT").MustInt(9090)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second



}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
