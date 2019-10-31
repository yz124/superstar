package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log"
	"time"

	"github.com/yz124/superstar/models"
	"github.com/yz124/superstar/services"
)

type AdminController struct {
	Ctx     iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	datalist := c.Service.GetAll()
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Datalist": datalist,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	var data *models.StarInfo
	if err == nil {
		data = c.Service.Get(id)
	} else {
		data = &models.StarInfo{
			Id: 0,
		}
	}
	//fmt.Println(id, data)
	// set the model and render the view template.
	return mvc.View{
		Name: "admin/edit.html",
		Data: iris.Map{
			"Title": "管理后台",
			"info":  data,
		},
		Layout: "admin/layout.html", // 不要跟前端的layout混用
	}
}

func (c *AdminController) PostSave() mvc.Result {
	info := models.StarInfo{}
	err := c.Ctx.ReadForm(&info)
	//fmt.Printf("%v\n", info)
	if err != nil {
		log.Fatal(err)
	}
	if info.Id > 0 {
		info.SysUpdated = int(time.Now().Unix())
		c.Service.Update(&info, []string{"name_zh", "name_en", "avatar",
			"birthday", "height", "weight", "club", "jersy", "coutry",
			"birthaddress", "feature", "moreinfo", "sys_updated"})
	} else {
		info.SysCreated = int(time.Now().Unix())
		c.Service.Create(&info)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}

func (c *AdminController) GetDelete() mvc.Result {
	id, err := c.Ctx.URLParamInt("id")
	if err == nil {
		c.Service.Delete(id)
	}
	return mvc.Response{
		Path: "/admin/",
	}
}
