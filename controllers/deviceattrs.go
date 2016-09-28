package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type DeviceAttrController struct {
	BaseController
}

func (this *DeviceAttrController) Post() {
	form := models.DeviceAttrPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceAttrPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceAttrPost:", &form)

	regDate := time.Now()
	deviceAttr := models.NewDeviceAttr(&form, regDate)
	beego.Debug("NewDeviceAttr:", deviceAttr)

	if code, err := deviceAttr.Insert(); err != nil {
		beego.Error("InsertDeviceAttr:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	deviceAttr.ClearPass()

	this.Data["json"] = &models.DeviceAttrPostInfo{Status:0, Code:0, DeviceAttrInfo: deviceAttr}
	this.ServeJSON()
}

func (this *DeviceAttrController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceAttrId:", err)
		this.RetError(errInputData)
		return
	}

	deviceAttr := models.DeviceAttr{}
	if code, err := deviceAttr.FindById(id); err != nil {
		beego.Error("FindDeviceAttrById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceAttrInfo:", &deviceAttr)

	deviceAttr.ClearPass()

	this.Data["json"] = &models.DeviceAttrGetOneInfo{Status:0, Code:0, DeviceAttrInfo: &deviceAttr}
	this.ServeJSON()
}

func (this *DeviceAttrController) GetAll() {
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

	deviceAttrs, err := models.GetAllDeviceAttrs(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceAttr:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceAttr:", &deviceAttrs)

	for i, _ := range deviceAttrs {
		deviceAttrs[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceAttrGetAllInfo{Status:0, Code:0, DeviceAttrsInfo: deviceAttrs}
	this.ServeJSON()
}

func (this *DeviceAttrController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceAttrId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.DeviceAttrPutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceAttrPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceAttrPut:", &form)

	deviceAttr := models.DeviceAttr{}
	if code, err := deviceAttr.UpdateById(id, &form); err != nil {
		beego.Error("UpdateDeviceAttrById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := deviceAttr.FindById(id); err != nil {
		beego.Error("FindDeviceAttrById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceAttrInfo:", &deviceAttr)

	deviceAttr.ClearPass()

	this.Data["json"] = &models.DeviceAttrPutInfo{Status:0, Code:0, DeviceAttrInfo: &deviceAttr}
	this.ServeJSON()
}

func (this *DeviceAttrController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceAttrId:", err)
		this.RetError(errInputData)
		return
	}

	deviceAttr := models.DeviceAttr{}
	if code, err := deviceAttr.DeleteById(id); err != nil {
		beego.Error("DeleteDeviceAttrById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceAttrDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
