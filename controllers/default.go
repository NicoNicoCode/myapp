package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"net"
	"log"
	"strings"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"crypto/md5"
	"encoding/hex"

	"flag"
	"github.com/golang/glog"

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

type SignupController struct{
	beego.Controller
}

type LoginController struct{
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Oem"] =  os.Getenv("OEM")+":"
	c.Data["Version"] = os.Getenv("VER")
	c.TplName = "index.tpl"
}

func (u *UserController) Get(){
	u.Ctx.WriteString("欢迎你使用本系统\n")
}

func (u *UserController) Profile(){

}

func (s *SignupController) Get(){
	s.Data["Oem"] =  os.Getenv("OEM")+":"
	s.Data["Version"] = os.Getenv("VER")
	s.TplName = "signup.tpl"
}

func (s *SignupController) Post(){
	username := s.GetString("u")
	password := s.GetString("p")
	info := s.GetString("i")

	flag.Parse()


	db,err := sql.Open("mysql",
		"root:123456@tcp(192.168.34.53:3306)/fun")
	if err!=nil{
		log.Fatal("数据库无法打开")
	}
	
	rows,err := db.Query("SELECT username FROM users WHERE username = ?",username)
	if err !=nil{
		s.TplName = "signup.tpl"
	}else{
		if rows.Next() == false{
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(password+"salt"))
			cipherStr := md5Ctx.Sum(nil)
			db.Exec("INSERT INTO users values(?,?,?)",username,hex.EncodeToString(cipherStr),info)
			s.Data["Scripts"]="注册成功"
			s.TplName = "signup.tpl"
		}else{
			s.Data["Scripts"]="注册失败"
			s.TplName = "signup.tpl"
		}
	}

}



func (l *LoginController) Get(){
	l.Data["Oem"] =  os.Getenv("OEM")+":"
	l.Data["Version"] = os.Getenv("VER")
	l.TplName = "login.tpl"
}

func (l *LoginController) Post(){
	username := l.GetString("u")
	password := l.GetString("p")

	db,err := sql.Open("mysql",
		"root:123456@tcp(192.168.34.53:3306)/fun")

	userController := UserController{}

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password+"salt"))
	cipherStr := md5Ctx.Sum(nil)

	rows,err := db.Query("SELECT username FROM users WHERE username = ? and password = ?",username,hex.EncodeToString(cipherStr))
	l.Data["Scripts"]=""

	if err != nil {
		log.Fatal(err)
	}else{
		if rows.Next()==true{
			var userName string
			var userInfo string
			err = rows.Scan(&userName,&userInfo)

			beego.Router("/user/profile",&userController,"*:Profile")
			url :=l.URLFor("UserController.Profile","username",username,"userinfo",userInfo)
		
			l.Redirect(url,302)

		}else{
			l.TplName="login.tpl"
			l.Data["Scripts"]="登陆失败"
		}

	}

}

func (p *ProfileController) Get(){
	host, err1 := os.Hostname()
	conn, err2:= net.Dial("udp","baidu.com:80")
	ua := p.Ctx.Request.UserAgent()

	username := p.GetString("username")
	userinfo := p.GetString("userinfo")
	if username==""{
		p.Ctx.WriteString("用户名:"+"无"+"\n")
	}else{
		p.Ctx.WriteString("用户名:"+username+"\n")
	}

	if userinfo==""{
		p.Ctx.WriteString("个人介绍:"+"无"+"\n")
	}else{
		p.Ctx.WriteString("个人介绍:"+userinfo+"\n")
	}
	
        p.Ctx.WriteString(ua+"\n")

	defer conn.Close()

	str:=strings.Split(conn.LocalAddr().String(),":")[0]

	if err2 != nil{
		p.Ctx.WriteString("The network is unconnect")
	}else{
		p.Ctx.WriteString(str+"\n")
	}

	if err1 != nil{
		fmt.Println()
	}else{
		p.Ctx.WriteString(host+"\n")
	}

	glog.Infoln("client's host:"+host)
	glog.Infoln("client's ip:"+str)

	p.Ctx.WriteString("\n\n\n\n\n\n\n\n\n\n"+os.Getenv("OEM")+":"+os.Getenv("VER"))
}

