package routers

import (
	"server/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/gift/create",&controllers.GiftCodeCreate{})
	beego.Router("/gift/get",&controllers.GiftGet{})
}
