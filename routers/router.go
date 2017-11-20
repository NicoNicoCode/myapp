package routers

import (
	"myapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user",&controllers.UserController{})
	beego.Router("/user/signup",&controllers.SignupController{})
	beego.Router("/user/login",&controllers.LoginController{})
	beego.Router("/user/profile",&controllers.ProfileController{})
}
