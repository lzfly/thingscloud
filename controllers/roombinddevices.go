package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type RoomBindDeviceController struct {
	BaseController
}

func (this *RoomBindDeviceController) Post() {
	form := models.RoomBindDevicePostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseRoomBindDevicePost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseRoomBindDevicePost:", &form)

	regDate := time.Now()
	roomBindDevice := models.NewRoomBindDevice(&form, regDate)
	beego.Debug("NewRoomBindDevice:", roomBindDevice)

	if code, err := roomBindDevice.Insert(); err != nil {
		beego.Error("InsertRoomBindDevice:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}


	this.Data["json"] = &models.RoomBindDevicePostInfo{RoomBindDeviceInfo: roomBindDevice}
	this.ServeJSON()
}

func (this *RoomBindDeviceController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseRoomBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	roomBindDevice := models.RoomBindDevice{}
	if code, err := roomBindDevice.FindById(id); err != nil {
		beego.Error("FindRoomBindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("RoomBindDeviceInfo:", &roomBindDevice)


	this.Data["json"] = &models.RoomBindDeviceGetOneInfo{RoomBindDeviceInfo: &roomBindDevice}
	this.ServeJSON()
}

func (this *RoomBindDeviceController) GetAll() {
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

	roomBindDevices, err := models.GetAllRoomBindDevices(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllRoomBindDevice:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllRoomBindDevice:", &roomBindDevices)


	this.Data["json"] = &models.RoomBindDeviceGetAllInfo{RoomBindDevicesInfo: roomBindDevices}
	this.ServeJSON()
}

func (this *RoomBindDeviceController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseRoomBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.RoomBindDevicePutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseRoomBindDevicePut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseRoomBindDevicePut:", &form)

	roomBindDevice := models.RoomBindDevice{}
	if code, err := roomBindDevice.UpdateById(id, &form); err != nil {
		beego.Error("UpdateRoomBindDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := roomBindDevice.FindById(id); err != nil {
		beego.Error("FindRoomBindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewRoomBindDeviceInfo:", &roomBindDevice)

	this.Data["json"] = &models.RoomBindDevicePutInfo{RoomBindDeviceInfo: &roomBindDevice}
	this.ServeJSON()
}

func (this *RoomBindDeviceController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseRoomBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	roomBindDevice := models.RoomBindDevice{}
	if code, err := roomBindDevice.DeleteById(id); err != nil {
		beego.Error("DeleteRoomBindDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
}
