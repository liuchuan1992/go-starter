package main

import "github.com/astaxie/beego"
var (
	BConfig *Config
)
type Config struct {
	AppName string
	ServerName string
}

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get()  {
	this.Ctx.WriteString("hello world")
}
func main() {
	beego.Router("/",&HomeController{})
	beego.Run()
}

func init() {
	BConfig = newConfig()
}

func newConfig() *Config {
	return &Config{
		AppName : "appName",
		ServerName :"serverName"}
}