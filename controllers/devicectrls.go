package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type DeviceCtrlController struct {
	BaseController
}

func (this *DeviceCtrlController) Post() {
	form := models.DeviceCtrlPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceCtrlPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceCtrlPost:", &form)

	regDate := time.Now()
	deviceCtrl := models.NewDeviceCtrl(&form, regDate)
	beego.Debug("NewDeviceCtrl:", deviceCtrl)

	if code, err := deviceCtrl.Insert(); err != nil {
		beego.Error("InsertDeviceCtrl:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	deviceCtrl.ClearPass()

	this.Data["json"] = &models.DeviceCtrlPostInfo{Status:0, Code:0, DeviceCtrlInfo: deviceCtrl}
	this.ServeJSON()
}

func (this *DeviceCtrlController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceCtrlId:", err)
		this.RetError(errInputData)
		return
	}

	deviceCtrl := models.DeviceCtrl{}
	if code, err := deviceCtrl.FindById(id); err != nil {
		beego.Error("FindDeviceCtrlById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceCtrlInfo:", &deviceCtrl)

	deviceCtrl.ClearPass()

	this.Data["json"] = &models.DeviceCtrlGetOneInfo{Status:0, Code:0, DeviceCtrlInfo: &deviceCtrl}
	this.ServeJSON()
}

func (this *DeviceCtrlController) GetAll() {
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

	deviceCtrls, err := models.GetAllDeviceCtrls(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceCtrl:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceCtrl:", &deviceCtrls)

	for i, _ := range deviceCtrls {
		deviceCtrls[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceCtrlGetAllInfo{Status:0, Code:0, DeviceCtrlsInfo: deviceCtrls}
	this.ServeJSON()
}

func (this *DeviceCtrlController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceCtrlId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.DeviceCtrlPutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceCtrlPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceCtrlPut:", &form)

	deviceCtrl := models.DeviceCtrl{}
	if code, err := deviceCtrl.UpdateById(id, &form); err != nil {
		beego.Error("UpdateDeviceCtrlById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := deviceCtrl.FindById(id); err != nil {
		beego.Error("FindDeviceCtrlById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceCtrlInfo:", &deviceCtrl)

	deviceCtrl.ClearPass()

	this.Data["json"] = &models.DeviceCtrlPutInfo{Status:0, Code:0, DeviceCtrlInfo: &deviceCtrl}
	this.ServeJSON()
}

func (this *DeviceCtrlController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceCtrlId:", err)
		this.RetError(errInputData)
		return
	}

	deviceCtrl := models.DeviceCtrl{}
	if code, err := deviceCtrl.DeleteById(id); err != nil {
		beego.Error("DeleteDeviceCtrlById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceCtrlDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
