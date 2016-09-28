package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type UserBindGatewayController struct {
	BaseController
}

func (this *UserBindGatewayController) Post() {
	form := models.UserBindGatewayPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseUserBindGatewayPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseUserBindGatewayPost:", &form)

	regDate := time.Now()
	
	form_g := models.GatewayPostForm{}
	gateway1 := models.NewGateway(&form_g, regDate)
	if _, err1 := gateway1.FindByGatewaySN(form.Gateway_sn); err1 != nil{
        this.RetError(errDupUser)
		return
	}
	
	
	userBindGateway1 := models.NewUserBindGateway(&form, regDate)
	if _, err1 := userBindGateway1.FindByUsernameGatewaySN(form.Username, form.Gateway_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	userBindGateway := models.NewUserBindGateway(&form, regDate)
	beego.Debug("NewUserBindGateway:", userBindGateway)

	if code, err := userBindGateway.Insert(); err != nil {
		beego.Error("InsertUserBindGateway:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	userBindGateway.ClearPass()

	this.Data["json"] = &models.UserBindGatewayPostInfo{Status:0, Code:0, UserBindGatewayInfo: userBindGateway}
	this.ServeJSON()
}

func (this *UserBindGatewayController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserBindGatewayId:", err)
		this.RetError(errInputData)
		return
	}

	userBindGateway := models.UserBindGateway{}
	if code, err := userBindGateway.FindById(id); err != nil {
		beego.Error("FindUserBindGatewayById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("UserBindGatewayInfo:", &userBindGateway)

	userBindGateway.ClearPass()

	this.Data["json"] = &models.UserBindGatewayGetOneInfo{Status:0, Code:0, UserBindGatewayInfo: &userBindGateway}
	this.ServeJSON()
}

func (this *UserBindGatewayController) GetAll() {
	queryVal, queryOp, err := this.ParseQueryParm()
	if err != nil {
		beego.Debug("ParseQuery:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("QueryVal:", queryVal)
	beego.Debug("QueryOp:", queryOp)

	order, err := this.ParseOrderParm()
	if err != nil {
		beego.Debug("ParseOrder:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("Order:", order)

	limit, err := this.ParseLimitParm()
	/*
		if err != nil {
			beego.Debug("ParseLimit:", err)
			this.RetError(errInputData)
			return
		}
	*/
	beego.Debug("Limit:", limit)

	offset, err := this.ParseOffsetParm()
	/*
		if err != nil {
			beego.Debug("ParseOffset:", err)
			this.RetError(errInputData)
			return
		}
	*/
	beego.Debug("Offset:", offset)

	userBindGateways, err := models.GetAllUserBindGateways(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllUserBindGateway:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllUserBindGateway:", &userBindGateways)

	for i, _ := range userBindGateways {
		userBindGateways[i].ClearPass()
	}

	this.Data["json"] = &models.UserBindGatewayGetAllInfo{Status:0, Code:0, UserBindGatewaysInfo: userBindGateways}
	this.ServeJSON()
}

func (this *UserBindGatewayController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserBindGatewayId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.UserBindGatewayPutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseUserBindGatewayPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseUserBindGatewayPut:", &form)

	userBindGateway := models.UserBindGateway{}
	if code, err := userBindGateway.UpdateById(id, &form); err != nil {
		beego.Error("UpdateUserBindGatewayById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := userBindGateway.FindById(id); err != nil {
		beego.Error("FindUserBindGatewayById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewUserBindGatewayInfo:", &userBindGateway)

	userBindGateway.ClearPass()

	this.Data["json"] = &models.UserBindGatewayPutInfo{Status:0, Code:0, UserBindGatewayInfo: &userBindGateway}
	this.ServeJSON()
}

func (this *UserBindGatewayController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserBindGatewayId:", err)
		this.RetError(errInputData)
		return
	}

	userBindGateway := models.UserBindGateway{}
	if code, err := userBindGateway.DeleteById(id); err != nil {
		beego.Error("DeleteUserBindGatewayById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	
	this.Data["json"] = &models.UserBindGatewayDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
