package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type DeviceGroupBindDeviceController struct {
	BaseController
}

func (this *DeviceGroupBindDeviceController) Post() {
	form := models.DeviceGroupBindDevicePostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceGroupBindDevicePost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceGroupBindDevicePost:", &form)

	regDate := time.Now()
	deviceGroupBindDevice := models.NewDeviceGroupBindDevice(&form, regDate)
	beego.Debug("NewDeviceGroupBindDevice:", deviceGroupBindDevice)

	if code, err := deviceGroupBindDevice.Insert(); err != nil {
		beego.Error("InsertDeviceGroupBindDevice:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}


	this.Data["json"] = &models.DeviceGroupBindDevicePostInfo{DeviceGroupBindDeviceInfo: deviceGroupBindDevice}
	this.ServeJSON()
}

func (this *DeviceGroupBindDeviceController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceGroupBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	deviceGroupBindDevice := models.DeviceGroupBindDevice{}
	if code, err := deviceGroupBindDevice.FindById(id); err != nil {
		beego.Error("FindDeviceGroupBindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceGroupBindDeviceInfo:", &deviceGroupBindDevice)


	this.Data["json"] = &models.DeviceGroupBindDeviceGetOneInfo{DeviceGroupBindDeviceInfo: &deviceGroupBindDevice}
	this.ServeJSON()
}

func (this *DeviceGroupBindDeviceController) GetAll() {
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

	deviceGroupBindDevices, err := models.GetAllDeviceGroupBindDevices(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceGroupBindDevice:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceGroupBindDevice:", &deviceGroupBindDevices)


	this.Data["json"] = &models.DeviceGroupBindDeviceGetAllInfo{DeviceGroupBindDevicesInfo: deviceGroupBindDevices}
	this.ServeJSON()
}

func (this *DeviceGroupBindDeviceController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceGroupBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.DeviceGroupBindDevicePutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceGroupBindDevicePut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceGroupBindDevicePut:", &form)

	deviceGroupBindDevice := models.DeviceGroupBindDevice{}
	if code, err := deviceGroupBindDevice.UpdateById(id, &form); err != nil {
		beego.Error("UpdateDeviceGroupBindDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := deviceGroupBindDevice.FindById(id); err != nil {
		beego.Error("FindDeviceGroupBindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceGroupBindDeviceInfo:", &deviceGroupBindDevice)

	this.Data["json"] = &models.DeviceGroupBindDevicePutInfo{DeviceGroupBindDeviceInfo: &deviceGroupBindDevice}
	this.ServeJSON()
}

func (this *DeviceGroupBindDeviceController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceGroupBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	deviceGroupBindDevice := models.DeviceGroupBindDevice{}
	if code, err := deviceGroupBindDevice.DeleteById(id); err != nil {
		beego.Error("DeleteDeviceGroupBindDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
}
