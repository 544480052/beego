package controllers


import (
	"github.com/astaxie/beego"
)


type TestController struct {
	beego.Controller
}



func (this *TestController) GetAll() {

	u := beego.BConfig.RunMode
	this.Ctx.WriteString(u)
	//this.Ctx.WriteString("what")
	//test := models.GetAllTests()
	//this.Data["json"] = test
	//this.ServeJSON()
}

func (this *TestController) Get()  {

	this.Ctx.WriteString("来耍")

}
