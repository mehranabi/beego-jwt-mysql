// @APIVersion 1.0.0
// @Title Beego with JWT
// @Description A minimal Beego API with JWT implementation and MySQL database
// @Contact mehran.ab80@gmail.com
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

package routers

import (
	"beego_jwt_sample/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Get("/", func(ctx *context.Context){
		_ = ctx.Output.Body([]byte("This is a Beego + JWT API - Creator: Mehran Abghari (mehran.ab80@gmail.com)"))
	})
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
