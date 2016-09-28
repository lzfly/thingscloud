package controllers

import (
	"thingscloud/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type UserController struct {
	BaseController
}

func (this *UserController) Post() {
	form := models.UserPostForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseUserPost:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseUserPost:", &form)

	regDate := time.Now()
	
	user1 := models.NewUser(&form, regDate)
	if _, err1 := user1.FindByUserName(form.Username); err1 == nil{
        this.RetError(errDupUser)
		return
	}
	
	
	user := models.NewUser(&form, regDate)
	beego.Debug("NewUser:", user)

	if code, err := user.Insert(); err != nil {
		beego.Error("InsertUser:", err)
		if code == models.ErrDupRows {
			this.RetError(errDupUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}

	user.ClearPass()

	this.Data["json"] = &models.UserPostInfo{Status:0, Code:0, UserInfo: user}
	this.ServeJSON()
}

func (this *UserController) GetOne() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserId:", err)
		this.RetError(errInputData)
		return
	}*/

	user := models.User{}
	//if code, err := user.FindById(id); err != nil {
	if code, err := user.FindByUserName(idStr); err != nil {
		beego.Error("FindUserById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("UserInfo:", &user)

	user.ClearPass()

	this.Data["json"] = &models.UserGetOneInfo{Status:0, Code:0, UserInfo: &user}
	this.ServeJSON()
}

func (this *UserController) GetAll() {
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

	users, err := models.GetAllUsers(queryVal, queryOp, order,
		limit, offset)
	if err != nil {
		beego.Error("GetAllUser:", err)
		this.RetError(errDatabase)
		return
	}
	beego.Debug("GetAllUser:", &users)

	for i, _ := range users {
		users[i].ClearPass()
	}

	this.Data["json"] = &models.UserGetAllInfo{Status:0, Code:0, UsersInfo: users}
	this.ServeJSON()
}

func (this *UserController) Put() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserId:", err)
		this.RetError(errInputData)
		return
	}*/

	form := models.UserPutForm{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &form)
	if err != nil {
		beego.Debug("ParseUserPut:", err)
		this.RetError(errInputData)
		return
	}
	beego.Debug("ParseUserPut:", &form)

	user := models.User{}
	//if code, err := user.UpdateById(id, &form); err != nil {
	if code, err := user.UpdateByUserName(idStr, &form); err != nil {
		beego.Error("UpdateUserById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUserChange)
		return
	}

	//if code, err := user.FindById(id); err != nil {
	if code, err := user.FindByUserName(idStr); err != nil {
		beego.Error("FindUserById:", err)
		if code == models.ErrNotFound {
			this.RetError(errNoUser)
		} else {
			this.RetError(errDatabase)
		}
		return
	}
	beego.Debug("NewUserInfo:", &user)

	user.ClearPass()

	this.Data["json"] = &models.UserPutInfo{Status:0, Code:0, UserInfo: &user}
	this.ServeJSON()
}

func (this *UserController) Delete() {
	idStr := this.Ctx.Input.Param(":id")
	/*id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		beego.Debug("ParseUserId:", err)
		this.RetError(errInputData)
		return
	}*/

	user := models.User{}
	//if code, err := user.DeleteById(id); err != nil {
	if code, err := user.DeleteByUserName(idStr); err != nil {
		beego.Error("DeleteUserById:", err)
		this.RetError(errDatabase)
		return
	} else if code == models.ErrNotFound {
		this.RetError(errNoUser)
		return
	}
	
	this.Data["json"] = &models.UserDeleteInfo{Status:0, Code:0}
	this.ServeJSON()
}
