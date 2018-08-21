package controllers


import (
	"github.com/astaxie/beego"
)


type TestController struct {
	beego.Controller
}



func (this *TestController) GetAll() {
	this.Ctx.WriteString("什么鬼")

}

func (this *TestController) Get()  {

	this.Ctx.WriteString("来耍耍")

}
