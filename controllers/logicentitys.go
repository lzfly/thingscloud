package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type LogicEntityController struct {
	BaseController
}

func (this *LogicEntityController) Post() {
	form := models.LogicEntityPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseLogicEntityPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseLogicEntityPost:", &form)

	regDate := time.Now()
	logicEntity := models.NewLogicEntity(&form, regDate)
	beego.Debug("NewLogicEntity:", logicEntity)

	if code, err := logicEntity.Insert(); err != nil {
		beego.Error("InsertLogicEntity:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}


	this.Data["json"] = &models.LogicEntityPostInfo{LogicEntityInfo: logicEntity}
	this.ServeJSON()
}

func (this *LogicEntityController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseLogicEntityId:", err)
		this.RetError(errInputData)
		return
	}

	logicEntity := models.LogicEntity{}
	if code, err := logicEntity.FindById(id); err != nil {
		beego.Error("FindLogicEntityById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("LogicEntityInfo:", &logicEntity)

	this.Data["json"] = &models.LogicEntityGetOneInfo{LogicEntityInfo: &logicEntity}
	this.ServeJSON()
}

func (this *LogicEntityController) GetAll() {
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

	logicEntitys, err := models.GetAllLogicEntitys(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllLogicEntity:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllLogicEntity:", &logicEntitys)


	this.Data["json"] = &models.LogicEntityGetAllInfo{LogicEntitysInfo: logicEntitys}
	this.ServeJSON()
}

func (this *LogicEntityController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseLogicEntityId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.LogicEntityPutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseLogicEntityPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseLogicEntityPut:", &form)

	logicEntity := models.LogicEntity{}
	if code, err := logicEntity.UpdateById(id, &form); err != nil {
		beego.Error("UpdateLogicEntityById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := logicEntity.FindById(id); err != nil {
		beego.Error("FindLogicEntityById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewLogicEntityInfo:", &logicEntity)


	this.Data["json"] = &models.LogicEntityPutInfo{LogicEntityInfo: &logicEntity}
	this.ServeJSON()
}

func (this *LogicEntityController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseLogicEntityId:", err)
		this.RetError(errInputData)
		return
	}

	logicEntity := models.LogicEntity{}
	if code, err := logicEntity.DeleteById(id); err != nil {
		beego.Error("DeleteLogicEntityById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
}
