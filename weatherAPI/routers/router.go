package routers

// here we implements the static files using beego.SetStaticPath("/down1", "download1")

import (
	"bootcamp-go/weatherAPI/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.WeatherController{})
}
