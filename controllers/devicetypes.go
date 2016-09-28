package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type DeviceTypeController struct {
	BaseController
}

func (this *DeviceTypeController) Post() {
	form := models.DeviceTypePostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceTypePost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceTypePost:", &form)

	regDate := time.Now()
	deviceType := models.NewDeviceType(&form, regDate)
	beego.Debug("NewDeviceType:", deviceType)

	if code, err := deviceType.Insert(); err != nil {
		beego.Error("InsertDeviceType:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	deviceType.ClearPass()

	this.Data["json"] = &models.DeviceTypePostInfo{Status:0, Code:0, DeviceTypeInfo: deviceType}
	this.ServeJSON()
}

func (this *DeviceTypeController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceTypeId:", err)
		this.RetError(errInputData)
		return
	}

	deviceType := models.DeviceType{}
	if code, err := deviceType.FindById(id); err != nil {
		beego.Error("FindDeviceTypeById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceTypeInfo:", &deviceType)

	deviceType.ClearPass()

	this.Data["json"] = &models.DeviceTypeGetOneInfo{Status:0, Code:0, DeviceTypeInfo: &deviceType}
	this.ServeJSON()
}

func (this *DeviceTypeController) GetAll() {
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

	deviceTypes, err := models.GetAllDeviceTypes(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceType:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceType:", &deviceTypes)

	for i, _ := range deviceTypes {
		deviceTypes[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceTypeGetAllInfo{Status:0, Code:0, DeviceTypesInfo: deviceTypes}
	this.ServeJSON()
}

func (this *DeviceTypeController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceTypeId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.DeviceTypePutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceTypePut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceTypePut:", &form)

	deviceType := models.DeviceType{}
	if code, err := deviceType.UpdateById(id, &form); err != nil {
		beego.Error("UpdateDeviceTypeById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := deviceType.FindById(id); err != nil {
		beego.Error("FindDeviceTypeById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceTypeInfo:", &deviceType)

	deviceType.ClearPass()

	this.Data["json"] = &models.DeviceTypePutInfo{Status:0, Code:0, DeviceTypeInfo: &deviceType}
	this.ServeJSON()
}

func (this *DeviceTypeController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceTypeId:", err)
		this.RetError(errInputData)
		return
	}

	deviceType := models.DeviceType{}
	if code, err := deviceType.DeleteById(id); err != nil {
		beego.Error("DeleteDeviceTypeById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	
	this.Data["json"] = &models.DeviceTypeDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
