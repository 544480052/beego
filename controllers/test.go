package controllers


import (
	"beegoApi/models"
	"github.com/astaxie/beego"
)


type TestController struct {
	beego.Controller
}



func (this *TestController) GetAll() {
	this.Ctx.WriteString("什么鬼")
	test := models2.GetAllTests()
	this.Data["json"] = test
	this.ServeJSON()
}

func (this *TestController) Get()  {

	this.Ctx.WriteString("来耍耍")

}
