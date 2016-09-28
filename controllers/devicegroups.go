package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type DeviceGroupController struct {
	BaseController
}

func (this *DeviceGroupController) Post() {
	form := models.DeviceGroupPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceGroupPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceGroupPost:", &form)

	regDate := time.Now()
	deviceGroup := models.NewDeviceGroup(&form, regDate)
	beego.Debug("NewDeviceGroup:", deviceGroup)

	if code, err := deviceGroup.Insert(); err != nil {
		beego.Error("InsertDeviceGroup:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}


	this.Data["json"] = &models.DeviceGroupPostInfo{DeviceGroupInfo: deviceGroup}
	this.ServeJSON()
}

func (this *DeviceGroupController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceGroupId:", err)
		this.RetError(errInputData)
		return
	}

	deviceGroup := models.DeviceGroup{}
	if code, err := deviceGroup.FindById(id); err != nil {
		beego.Error("FindDeviceGroupById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceGroupInfo:", &deviceGroup)


	this.Data["json"] = &models.DeviceGroupGetOneInfo{DeviceGroupInfo: &deviceGroup}
	this.ServeJSON()
}

func (this *DeviceGroupController) GetAll() {
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

	deviceGroups, err := models.GetAllDeviceGroups(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceGroup:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceGroup:", &deviceGroups)


	this.Data["json"] = &models.DeviceGroupGetAllInfo{DeviceGroupsInfo: deviceGroups}
	this.ServeJSON()
}

func (this *DeviceGroupController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceGroupId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.DeviceGroupPutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceGroupPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceGroupPut:", &form)

	deviceGroup := models.DeviceGroup{}
	if code, err := deviceGroup.UpdateById(id, &form); err != nil {
		beego.Error("UpdateDeviceGroupById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := deviceGroup.FindById(id); err != nil {
		beego.Error("FindDeviceGroupById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceGroupInfo:", &deviceGroup)


	this.Data["json"] = &models.DeviceGroupPutInfo{DeviceGroupInfo: &deviceGroup}
	this.ServeJSON()
}

func (this *DeviceGroupController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceGroupId:", err)
		this.RetError(errInputData)
		return
	}

	deviceGroup := models.DeviceGroup{}
	
	if _, err := deviceGroup.FindById(id); err != nil {
		beego.Error("FindDeviceGroupById:", err)

	}
	
	deviceGroupBindDevice := models.DeviceGroupBindDevice{}
	if _, err := deviceGroupBindDevice.DeleteByGroupName(deviceGroup.Username, deviceGroup.Group_name); err != nil {
		beego.Error("DeleteDeviceGroupById:", err)

	}
	
	
	if code, err := deviceGroup.DeleteById(id); err != nil {
		beego.Error("DeleteDeviceGroupById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
}
