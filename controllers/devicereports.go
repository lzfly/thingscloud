package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type DeviceReportController struct {
	BaseController
}

func (this *DeviceReportController) Post() {
	form := models.DeviceReportPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceReportPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceReportPost:", &form)

	regDate := time.Now()
	deviceReport := models.NewDeviceReport(&form, regDate)
	beego.Debug("NewDeviceReport:", deviceReport)

	if code, err := deviceReport.Insert(); err != nil {
		beego.Error("InsertDeviceReport:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	deviceReport.ClearPass()

	this.Data["json"] = &models.DeviceReportPostInfo{Status:0, Code:0, DeviceReportInfo: deviceReport}
	this.ServeJSON()
}

func (this *DeviceReportController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceReportId:", err)
		this.RetError(errInputData)
		return
	}

	deviceReport := models.DeviceReport{}
	if code, err := deviceReport.FindById(id); err != nil {
		beego.Error("FindDeviceReportById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("DeviceReportInfo:", &deviceReport)

	deviceReport.ClearPass()

	this.Data["json"] = &models.DeviceReportGetOneInfo{Status:0, Code:0, DeviceReportInfo: &deviceReport}
	this.ServeJSON()
}

func (this *DeviceReportController) GetAll() {
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

	deviceReports, err := models.GetAllDeviceReports(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllDeviceReport:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllDeviceReport:", &deviceReports)

	for i, _ := range deviceReports {
		deviceReports[i].ClearPass()
	}

	this.Data["json"] = &models.DeviceReportGetAllInfo{Status:0, Code:0, DeviceReportsInfo: deviceReports}
	this.ServeJSON()
}

func (this *DeviceReportController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceReportId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.DeviceReportPutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseDeviceReportPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseDeviceReportPut:", &form)

	deviceReport := models.DeviceReport{}
	if code, err := deviceReport.UpdateById(id, &form); err != nil {
		beego.Error("UpdateDeviceReportById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := deviceReport.FindById(id); err != nil {
		beego.Error("FindDeviceReportById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewDeviceReportInfo:", &deviceReport)

	deviceReport.ClearPass()

	this.Data["json"] = &models.DeviceReportPutInfo{Status:0, Code:0, DeviceReportInfo: &deviceReport}
	this.ServeJSON()
}

func (this *DeviceReportController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseDeviceReportId:", err)
		this.RetError(errInputData)
		return
	}

	deviceReport := models.DeviceReport{}
	if code, err := deviceReport.DeleteById(id); err != nil {
		beego.Error("DeleteDeviceReportById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	this.Data["json"] = &models.DeviceReportDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
