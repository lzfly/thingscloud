package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type SessionController struct {
	BaseController
}

func (this *SessionController) Post() {
	form := models.SessionPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseSessionPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseSessionPost:", &form)

	if err := this.VerifyForm(&form); err != nil {
		beego.Debug("ValidLoginForm:", err)
		this.Data["json"] = models.NewErrorInfo(ErrInputData)
		this.ServeJSON()
		return
	}

	user := models.User{}
	if code, err := user.FindByUserName(form.Username); err != nil {
		beego.Error("FindUserByUserName:", err)
		if code == models.ErrNotFound {
			this.Data["json"] = models.NewErrorInfo(ErrNoUser)
		} else {
			this.Data["json"] = models.NewErrorInfo(ErrDatabase)
		}
		this.ServeJSON()
		return
	}
	beego.Debug("UserInfo:", &user)

	if form.Password != user.Password {
		this.Data["json"] = models.NewErrorInfo(ErrPass)
		this.ServeJSON()
		return
	}
	user.ClearPass()

	this.SetSession("user_id", form.Username)

	this.Data["json"] = &models.SessionPostInfo{Status:0, Code:0, Sessionid: form.Username, Userid: user.Id}
	this.ServeJSON()
}

func (this *SessionController) Delete() {

	idStr := this.Ctx.Input.Param(":id")

	if this.GetSession("user_id") != idStr {
		this.Data["json"] = models.NewErrorInfo(ErrInvalidUser)
		this.ServeJSON()
		return
	}

	this.DelSession("user_id")

	this.Data["json"] = &models.SessionDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
