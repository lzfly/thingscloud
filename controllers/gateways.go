package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type GatewayController struct {
	BaseController
}

func (this *GatewayController) Post() {
	form := models.GatewayPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseGatewayPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseGatewayPost:", &form)

	regDate := time.Now()
	
	gateway1 := models.NewGateway(&form, regDate)
	if _, err1 := gateway1.FindByGatewaySN(form.Gateway_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	gateway := models.NewGateway(&form, regDate)
	beego.Debug("NewGateway:", gateway)

	if code, err := gateway.Insert(); err != nil {
		beego.Error("InsertGateway:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	gateway.ClearPass()

	this.Data["json"] = &models.GatewayPostInfo{Status:0, Code:0, GatewayInfo: gateway}
	this.ServeJSON()
}

func (this *GatewayController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseGatewayId:", err)
		this.RetError(errInputData)
		return
	}*/

	gateway := models.Gateway{}
	if code, err := gateway.FindByGatewaySN(idStr); err != nil {
		beego.Error("FindGatewayById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("GatewayInfo:", &gateway)

	gateway.ClearPass()

	this.Data["json"] = &models.GatewayGetOneInfo{Status:0, Code:0, GatewayInfo: &gateway}
	this.ServeJSON()
}

func (this *GatewayController) GetAll() {
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

	gateways, err := models.GetAllGateways(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllGateway:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllGateway:", &gateways)

	for i, _ := range gateways {
		gateways[i].ClearPass()
	}

	this.Data["json"] = &models.GatewayGetAllInfo{Status:0, Code:0, GatewaysInfo: gateways}
	this.ServeJSON()
}

func (this *GatewayController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseGatewayId:", err)
		this.RetError(errInputData)
		return
	}*/

	form := models.GatewayPutForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseGatewayPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseGatewayPut:", &form)

	gateway := models.Gateway{}
	if code, err := gateway.UpdateByGatewaySN(idStr, &form); err != nil {
		beego.Error("UpdateGatewayById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := gateway.FindByGatewaySN(idStr); err != nil {
		beego.Error("FindGatewayById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewGatewayInfo:", &gateway)

	gateway.ClearPass()

	this.Data["json"] = &models.GatewayPutInfo{Status:0, Code:0, GatewayInfo: &gateway}
	this.ServeJSON()
}

func (this *GatewayController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseGatewayId:", err)
		this.RetError(errInputData)
		return
	}*/

	gateway := models.Gateway{}
	if code, err := gateway.DeleteByGatewaySN(idStr); err != nil {
		beego.Error("DeleteGatewayById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	
	this.Data["json"] = &models.GatewayDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
