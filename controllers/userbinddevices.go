package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type UserBindDeviceController struct {
	BaseController
}

func (this *UserBindDeviceController) Post() {
	form := models.UserBindDevicePostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseUserBindDevicePost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseUserBindDevicePost:", &form)

	regDate := time.Now()
	
	form_d := models.DevicePostForm{}
	device1 := models.NewDevice(&form_d, regDate)
	if _, err1 := device1.FindByDeviceSN(form.Device_sn); err1 != nil{
        this.RetError(errDupUser)
		return
	}
	
	userBindDevice1 := models.NewUserBindDevice(&form, regDate)
	if _, err1 := userBindDevice1.FindByUsernameDeviceSN(form.Username, form.Device_sn); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	userBindDevice := models.NewUserBindDevice(&form, regDate)
	beego.Debug("NewUserBindDevice:", userBindDevice)

	if code, err := userBindDevice.Insert(); err != nil {
		beego.Error("InsertUserBindDevice:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	userBindDevice.ClearPass()

	this.Data["json"] = &models.UserBindDevicePostInfo{Status:0, Code:0, UserBindDeviceInfo: userBindDevice}
	this.ServeJSON()
}

func (this *UserBindDeviceController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	userBindDevice := models.UserBindDevice{}
	if code, err := userBindDevice.FindById(id); err != nil {
		beego.Error("FindUserBindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("UserBindDeviceInfo:", &userBindDevice)

	userBindDevice.ClearPass()

	this.Data["json"] = &models.UserBindDeviceGetOneInfo{Status:0, Code:0, UserBindDeviceInfo: &userBindDevice}
	this.ServeJSON()
}

func (this *UserBindDeviceController) GetAll() {
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

	userBindDevices, err := models.GetAllUserBindDevices(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllUserBindDevice:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllUserBindDevice:", &userBindDevices)

	for i, _ := range userBindDevices {
		userBindDevices[i].ClearPass()
	}

	this.Data["json"] = &models.UserBindDeviceGetAllInfo{Status:0, Code:0, UserBindDevicesInfo: userBindDevices}
	this.ServeJSON()
}

func (this *UserBindDeviceController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.UserBindDevicePutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseUserBindDevicePut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseUserBindDevicePut:", &form)

	userBindDevice := models.UserBindDevice{}
	if code, err := userBindDevice.UpdateById(id, &form); err != nil {
		beego.Error("UpdateUserBindDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := userBindDevice.FindById(id); err != nil {
		beego.Error("FindUserBindDeviceById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewUserBindDeviceInfo:", &userBindDevice)

	userBindDevice.ClearPass()

	this.Data["json"] = &models.UserBindDevicePutInfo{Status:0, Code:0, UserBindDeviceInfo: &userBindDevice}
	this.ServeJSON()
}

func (this *UserBindDeviceController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserBindDeviceId:", err)
		this.RetError(errInputData)
		return
	}

	userBindDevice := models.UserBindDevice{}
	if code, err := userBindDevice.DeleteById(id); err != nil {
		beego.Error("DeleteUserBindDeviceById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}

    this.Data["json"] = &models.UserBindDeviceDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
