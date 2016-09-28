package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type DeviceController struct {
	BaseController
}

func (this *DeviceController) Post() {
	form := models.DevicePostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDevicePost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDevicePost:", &form)

	regDate := time.Now()
	
	device1 := models.NewDevice(&form, regDate)
	if _, err1 := device1.FindByDeviceSN(form.Device_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	deviceType := models.DeviceType{}
	if _, err := deviceType.FindByTypeCode(form.Type_code); err != nil {
		beego.Error("FindDeviceTypeByTypeCode:", err)
	}
	
	form.Type_name = deviceType.Type_name
	
	device := models.NewDevice(&form, regDate)
	beego.Debug("NewDevice:", device)

	if code, err := device.Insert(); err != nil {
		beego.Error("InsertDevice:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	device.ClearPass()

	this.Data["json"] = &models.DevicePostInfo{Status:0, Code:0, DeviceInfo: device}
	this.ServeJSON()
}

func (this *DeviceController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceId:", err)
		this.RetError(errInputData)
		return
	}*/

	device := models.Device{}
	if code, err := device.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceInfo:", &device)

	device.ClearPass()

	this.Data["json"] = &models.DeviceGetOneInfo{Status:0, Code:0, DeviceInfo: &device}
	this.ServeJSON()
}

func (this *DeviceController) GetAll() {
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

	devices, err := models.GetAllDevices(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDevice:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDevice:", &devices)

	for i, _ := range devices {
		devices[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceGetAllInfo{Status:0, Code:0, DevicesInfo: devices}
	this.ServeJSON()
}

func (this *DeviceController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceId:", err)
		this.RetError(errInputData)
		return
	}*/

	form := models.DevicePutForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDevicePut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDevicePut:", &form)

	device := models.Device{}
	if code, err := device.UpdateByDeviceSN(idStr, &form); err != nil {
		beego.Error("UpdateDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := device.FindByDeviceSN(idStr); err != nil {
		beego.Error("FindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceInfo:", &device)

	device.ClearPass()

	this.Data["json"] = &models.DevicePutInfo{Status:0, Code:0, DeviceInfo: &device}
	this.ServeJSON()
}

func (this *DeviceController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceId:", err)
		this.RetError(errInputData)
		return
	}*/

	device := models.Device{}

	deviceAttrInfo := models.DeviceAttrInfo{}
	if _, err := deviceAttrInfo.DeleteByDeviceSN(idStr); err != nil {
		beego.Error("DeleteDeviceAttrInfoById:", err)

	}
	
	if code, err := device.DeleteByDeviceSN(idStr); err != nil {
		beego.Error("DeleteDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
