package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type RoomController struct {
	BaseController
}

func (this *RoomController) Post() {
	form := models.RoomPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseRoomPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseRoomPost:", &form)

	regDate := time.Now()
	room := models.NewRoom(&form, regDate)
	beego.Debug("NewRoom:", room)

	if code, err := room.Insert(); err != nil {
		beego.Error("InsertRoom:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}


	this.Data["json"] = &models.RoomPostInfo{RoomInfo: room}
	this.ServeJSON()
}

func (this *RoomController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseRoomId:", err)
		this.RetError(errInputData)
		return
	}

	room := models.Room{}
	if code, err := room.FindById(id); err != nil {
		beego.Error("FindRoomById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("RoomInfo:", &room)


	this.Data["json"] = &models.RoomGetOneInfo{RoomInfo: &room}
	this.ServeJSON()
}

func (this *RoomController) GetAll() {
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

	rooms, err := models.GetAllRooms(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllRoom:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllRoom:", &rooms)


	this.Data["json"] = &models.RoomGetAllInfo{RoomsInfo: rooms}
	this.ServeJSON()
}

func (this *RoomController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseRoomId:", err)
		this.RetError(errInputData)
		return
	}

	form := models.RoomPutForm{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseRoomPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseRoomPut:", &form)

	room := models.Room{}
	if code, err := room.UpdateById(id, &form); err != nil {
		beego.Error("UpdateRoomById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	if code, err := room.FindById(id); err != nil {
		beego.Error("FindRoomById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewRoomInfo:", &room)


	this.Data["json"] = &models.RoomPutInfo{RoomInfo: &room}
	this.ServeJSON()
}

func (this *RoomController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseRoomId:", err)
		this.RetError(errInputData)
		return
	}

	room := models.Room{}
	if code, err := room.DeleteById(id); err != nil {
		beego.Error("DeleteRoomById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
}
