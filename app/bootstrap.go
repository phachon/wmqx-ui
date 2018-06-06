package app

import (
	"os"
	"log"
	"fmt"
	"flag"
	"wmqx-ui/app/utils"
	"github.com/astaxie/beego"
	"github.com/snail007/go-activerecord/mysql"
	"github.com/fatih/color"
	"wmqx-ui/app/models"
)

var (
	confPath = flag.String("conf", "conf/wmqx-ui.conf", "please set wmqx-ui conf path")
)

var (
	Version = "v0.1"
)

func init() {
	poster()
	initConfig()
	initDB()
}

// poster logo
func poster() {
	fg := color.New(color.FgBlue)
	logo := `
__        __  __  __    ___   __  __          _   _   ___ 
\ \      / / |  \/  |  / _ \  \ \/ /         | | | | |_ _|
 \ \ /\ / /  | |\/| | | | | |  \  /   _____  | | | |  | | 
  \ V  V /   | |  | | | |_| |  /  \  |_____| | |_| |  | | 
   \_/\_/    |_|  |_|  \__\_\ /_/\_\          \___/  |___|
`+
"Author: phachon\r\n"+
"Version: "+Version+"\r\n"+
"Link: https://github.com/phachon/wmqx-ui"
	fg.Println(logo)
}

// init beego config
func initConfig()  {

	flag.Parse()

	if *confPath == "" {
		log.Println("conf file not empty!")
		os.Exit(1)
	}
	ok, _ := utils.NewFile().PathIsExists(*confPath)
	if ok == false{
		log.Println("conf file "+*confPath+" not exists!")
		os.Exit(1)
	}
	//init config file
	beego.LoadAppConfig("ini", *confPath)

	// init name
	beego.AppConfig.Set("sys.name", "wmqx-ui")
	beego.BConfig.AppName = beego.AppConfig.String("sys.name")
	beego.BConfig.ServerName = beego.AppConfig.String("sys.name")

	// set static path
	beego.SetStaticPath("/static/", "static")

	// views path
	beego.BConfig.WebConfig.ViewsPath = "views/"

	// session
	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = ".session"
	beego.BConfig.WebConfig.Session.SessionName = "wmqxssid"
	beego.BConfig.WebConfig.Session.SessionOn = true

	// log
	logConfigs, err := beego.AppConfig.GetSection("log")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	for adapter, config := range logConfigs {
		beego.SetLogger(adapter, config)
	}
	beego.SetLogFuncCall(true)
}

//init db
func initDB()  {
	host := beego.AppConfig.String("db::host")
	port, _ := beego.AppConfig.Int("db::port")
	user := beego.AppConfig.String("db::user")
	pass := beego.AppConfig.String("db::pass")
	dbname := beego.AppConfig.String("db::name")
	dbTablePrefix := beego.AppConfig.String("db::table_prefix")
	maxIdle, _ := beego.AppConfig.Int("db::conn_max_idle")
	maxConn, _ := beego.AppConfig.Int("db::conn_max_connection")
	models.G = mysql.NewDBGroup("default")
	cfg := mysql.NewDBConfigWith(host, port, dbname, user, pass)
	cfg.SetMaxIdleConns = maxIdle
	cfg.SetMaxOpenConns = maxConn
	cfg.TablePrefix = dbTablePrefix
	cfg.TablePrefixSqlIdentifier = "__PREFIX__"
	err := models.G.Regist("default", cfg)
	if err != nil {
		beego.Error(fmt.Errorf("database error:%s,with config : %v", err, cfg))
		os.Exit(1)
	}
}