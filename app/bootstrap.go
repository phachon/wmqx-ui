package app

import (
	"os"
	"log"
	"fmt"
	"flag"
	"net/http"
	"wmqx-ui/app/utils"
	"github.com/astaxie/beego"
	"wmqx-ui/app/controllers"
	"github.com/snail007/go-activerecord/mysql"
	"github.com/fatih/color"
	"wmqx-ui/app/models"
)

var (
	confPath = flag.String("conf", "conf/wmqx-ui.conf", "please set wmqx-ui conf path")
)

var (
	version = "v0.1"
)

func init() {
	poster()
	initConfig()
	initDB()
	initRouter()
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
"Version: "+version+"\r\n"+
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
	}
	for adapter, config := range logConfigs {
		beego.SetLogger(adapter, config)
	}
	beego.SetLogFuncCall(true)
}

func initRouter() {
	// router
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.RouterCaseSensitive = false
	
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/author", &controllers.AuthorController{}, "*:Index")
	beego.AutoRouter(&controllers.AuthorController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.NodeController{})
	beego.AutoRouter(&controllers.MessageController{})
	beego.AutoRouter(&controllers.ConsumerController{})
	beego.AutoRouter(&controllers.LogController{})
	beego.AutoRouter(&controllers.ProfileController{})
	beego.ErrorHandler("404", http_404)
	beego.ErrorHandler("500", http_500)

	// add template func
	beego.AddFuncMap("dateFormat", utils.NewDate().Format)
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
		os.Exit(100)
	}
}

func http_404(rs http.ResponseWriter, req *http.Request) {
	rs.Write([]byte("404 not found!"))
}

func http_500(rs http.ResponseWriter, req *http.Request) {
	rs.Write([]byte("500 server error!"))
}
