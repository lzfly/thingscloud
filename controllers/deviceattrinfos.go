package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	"strings"
)

type DeviceAttrInfoController struct {
	BaseController
}

func (this *DeviceAttrInfoController) Post() {
	form := models.DeviceAttrInfoPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceAttrInfoPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceAttrInfoPost:", &form)

	regDate := time.Now()
	
    deviceAttrInfo1 := models.NewDeviceAttrInfo(&form, regDate)
	if _, err1 := deviceAttrInfo1.FindByAttrSN(form.Device_sn, form.Attr_code); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	deviceAttr := models.DeviceAttr{}
	if _, err := deviceAttr.FindByAttrCode(form.Attr_code); err != nil {
		beego.Error("FindDeviceAttrByAttrCode:", err)
	}
	
	form.Attr_name = deviceAttr.Attr_name
	
	
	deviceAttrInfo := models.NewDeviceAttrInfo(&form, regDate)
	beego.Debug("NewDeviceAttrInfo:", deviceAttrInfo)

	if code, err := deviceAttrInfo.Insert(); err != nil {
		beego.Error("InsertDeviceAttrInfo:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	deviceAttrInfo.ClearPass()

	this.Data["json"] = &models.DeviceAttrInfoPostInfo{Status:0, Code:0, DeviceAttrInfoInfo: deviceAttrInfo}
	this.ServeJSON()
}

func (this *DeviceAttrInfoController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		beego.Debug("ParseDeviceAttrInfoId:", err)
		this.RetError(errInputData)
		return
	}*/

    var str []string
    str = strings.Split(idStr, ":")
    device_sn := str[0]
	id, _ :=  strconv.ParseInt(str[1], 0, 32)
	attr_code := int32(id)
	
	deviceAttrInfo := models.DeviceAttrInfo{}
	if code, err := deviceAttrInfo.FindByAttrSN(device_sn, attr_code); err != nil {
		beego.Error("FindDeviceAttrInfoById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceAttrInfoInfo:", &deviceAttrInfo)

	deviceAttrInfo.ClearPass()

	this.Data["json"] = &models.DeviceAttrInfoGetOneInfo{Status:0, Code:0, DeviceAttrInfoInfo: &deviceAttrInfo}
	this.ServeJSON()
}

func (this *DeviceAttrInfoController) GetAll() {
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

	deviceAttrInfos, err := models.GetAllDeviceAttrInfos(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceAttrInfo:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceAttrInfo:", &deviceAttrInfos)

	for i, _ := range deviceAttrInfos {
		deviceAttrInfos[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceAttrInfoGetAllInfo{Status:0, Code:0, DeviceAttrInfosInfo: deviceAttrInfos}
	this.ServeJSON()
}

func (this *DeviceAttrInfoController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceAttrInfoId:", err)
		this.RetError(errInputData)
		return
	}*/
	
    var str []string
    str = strings.Split(idStr, ":")
    device_sn := str[0]
	id, err :=  strconv.ParseInt(str[1], 0, 32)
	attr_code := int32(id)

	form := models.DeviceAttrInfoPutForm{}
	form.Attr_code = -1
	form.Is_control = -1

	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceAttrInfoPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceAttrInfoPut:", &form)

	deviceAttrInfo := models.DeviceAttrInfo{}
	if code, err := deviceAttrInfo.UpdateByAttrSN(device_sn, attr_code, &form); err != nil {
		beego.Error("UpdateDeviceAttrInfoById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := deviceAttrInfo.FindByAttrSN(device_sn, attr_code); err != nil {
		beego.Error("FindDeviceAttrInfoById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceAttrInfoInfo:", &deviceAttrInfo)

	deviceAttrInfo.ClearPass()

	this.Data["json"] = &models.DeviceAttrInfoPutInfo{Status:0, Code:0, DeviceAttrInfoInfo: &deviceAttrInfo}
	this.ServeJSON()
}

func (this *DeviceAttrInfoController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 32)
	if err != nil {
		beego.Debug("ParseDeviceAttrInfoId:", err)
		this.RetError(errInputData)
		return
	}*/

	var str []string
    str = strings.Split(idStr, ":")
    device_sn := str[0]
	id, _ :=  strconv.ParseInt(str[1], 0, 32)
	attr_code := int32(id)
	
	deviceAttrInfo := models.DeviceAttrInfo{}
	if code, err := deviceAttrInfo.DeleteByAttrSN(device_sn, attr_code); err != nil {
		beego.Error("DeleteDeviceAttrInfoById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceAttrInfoDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
