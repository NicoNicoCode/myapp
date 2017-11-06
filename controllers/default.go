package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"net"
	"strings"
)

type MainController struct {
	beego.Controller
}

type UserController struct{
	beego.Controller
}

type ProfileController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.TplName = "index.tpl"
}

func (u *UserController) Get(){
	u.Ctx.WriteString("请进入Profile文件夹\n")
}

func (p *ProfileController) Get(){
	host, err1 := os.Hostname()
	conn, err2:= net.Dial("udp","baidu.com:80")
	ua := p.Ctx.Request.UserAgent()

	p.Ctx.WriteString(ua+"\n")

	defer conn.Close()

	if err2 != nil{
		fmt.Println()
	}else{
		p.Ctx.WriteString(strings.Split(conn.LocalAddr().String(),":")[0]+"\n")
	}

	if err1 != nil{
		fmt.Println()
	}else{
		p.Ctx.WriteString(host+"\n")
	}
}

